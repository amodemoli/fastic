// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"fmt"
	"sync"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

// this struct maded for visitor info.
type visitor struct {
	limiter  *rate.Limiter // limiter
	lastSeen time.Time     // last seen of visitor.
}

// limiter struct model, for clean code and easy use i used struct
type Limiter struct {
	visitors   map[string]*visitor // visitors as map
	mu         sync.RWMutex
	rate       rate.Limit // rate limit
	brust      int
	expiration time.Duration // expiration time for clear automatic
}

// this function creates new rate limiter for use
func NewLimiter(r rate.Limit, b int) *Limiter {
	return &Limiter{
		visitors:   make(map[string]*visitor), // create the map
		rate:       r,
		brust:      b,
		expiration: 5 * time.Minute, // expire seen data on map.
	}
}

// this function maded for get limiter info and stats
func (l *Limiter) getLimiter(ip string) *rate.Limiter {
	l.mu.Lock()         // lock the mutex for write
	defer l.mu.Unlock() // use defer for unlock

	v, exists := l.visitors[ip]
	if exists {
		v.lastSeen = time.Now()
		return v.limiter
	}

	// create new limiter
	limiter := rate.NewLimiter(l.rate, l.brust)
	l.visitors[ip] = &visitor{ // save correct user ip and last seen data to map
		limiter:  limiter,
		lastSeen: time.Now(),
	}
	return limiter // return new limiter.
}

// cleanUp removes IPs that haven't been seen for longer than expiration
func (l *Limiter) cleanUp() {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	for ip, v := range l.visitors {
		// if IP hasn't been seen for more than expiration time, remove it
		if now.Sub(v.lastSeen) > l.expiration {
			delete(l.visitors, ip)
		}
	}
}

// startCleanupRoutine runs the cleanup periodically
func (l *Limiter) startCleanupRoutine() {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			l.cleanUp()
		}
	}()
}

// Limiter middleware, maded for limit requests per second, anti ddos and anti dos. <3
func (l *Limiter) Limiter(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	l.startCleanupRoutine() // first clean up
	return func(ctx *fasthttp.RequestCtx) {
		ip := ctx.RemoteIP().String() // get user ip

		limiter := l.getLimiter(ip) // get the limiter info for ip

		if !limiter.Allow() { // if limiter dont allow
			ctx.SetStatusCode(fasthttp.StatusTooManyRequests)                                                                                      // change status code
			ctx.SetContentType("application/json")                                                                                                 // change page content type for json
			ctx.Response.SetBodyString(fmt.Sprintf(`{"message": "too many requests, maximum %s requests per second"}`, string(rune(int(l.rate))))) // write json
			return
		}

		next(ctx) // call next
	}
}
// fix: fixed the memory leak.