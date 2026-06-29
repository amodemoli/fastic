// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"sync"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

// limiter struct model, for clean code and easy use i used struct
type Limiter struct {
	visitors   sync.Map // visitors as map
	mu         sync.Mutex
	rate       rate.Limit // rate limit
	brust      int
	expiration time.Duration // expiration time for clear automatic
}

// this function creates new rate limiter for use
func NewLimiter(r rate.Limit, b int) *Limiter {
	l := &Limiter{
		rate:       r,
		brust:      b,
		expiration: 5 * time.Minute,
	}
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for range ticker.C {
			l.visitors.Range(func(key, value interface{}) bool {
				return true
			})
		}
	}()
	return l
}

// this function maded for get limiter info and stats
func (l *Limiter) getLimiter(ip string) *rate.Limiter {
	// check for the ip on map.
	if val, ok := l.visitors.Load(ip); ok {
		return val.(*rate.Limiter) // return the limiter
	}

	// create new limiter
	limiter := rate.NewLimiter(l.rate, l.brust)
	l.visitors.Store(ip, limiter)
	return limiter
}

// Limiter middleware, maded for limit requests per second, anti ddos and anti dos. <3
func (l *Limiter) Limiter(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ip := ctx.RemoteIP().String() // get user ip with context of request

		if !l.getLimiter(ip).Allow() { // if user is lmit
			ctx.SetStatusCode(fasthttp.StatusTooManyRequests)                                       // change status code to TooManyRequests (Code 429).
			ctx.SetContentType("application/json")                                                  // change response content type to application/json (json response)
			ctx.WriteString(`{"limited": "to many requests, you can send 5 requests from second"}`) // show limited alert to user.
			return                                                                                  // if user is limited return the application
		}

		next(ctx) // if user is't limit run the request function.
	}
}
