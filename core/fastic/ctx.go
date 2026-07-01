// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"sync"

	"github.com/valyala/fasthttp"
)

// handler function for fastic.
type RequestHandler func(*Ctx)

// create new structure named Ctx for create custom request ctx
// only for requests for fastic framework
type Ctx struct {
	*fasthttp.RequestCtx
}

var ctxPool = sync.Pool{
	New: func() interface{} {
		return &Ctx{}
	},
}

// new pool for Ctx named ctxPool, for speed and performance i use pool.

// this function maded for get ctx from pool and drop for
// my target function (dont need create new struct) i can create struct for first time and use struct for +1000 times =D
// only need to get struct from pool and drop it to pool after use.
func acquireCtx(ctx *fasthttp.RequestCtx) *Ctx {
	c := ctxPool.Get().(*Ctx) // get Ctx struct from pool
	c.RequestCtx = ctx        // save correct ctx to c.RequestCtx value =D
	return c
}

// this function maded for drop ctx after using =D
// only for use other times.
func releaseCtx(c *Ctx) {
	c.RequestCtx = nil // remove c.RequestCtx pointer for performance of GC (Garbge Colector)
	ctxPool.Put(c)     // back the value to pool
}