// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"github.com/valyala/fasthttp"
)

// this function maded for run multiplate middleware's. (use on main.go) # only for main middlewares
func ChainMdw(handler fasthttp.RequestHandler, middlewares ...func(fasthttp.RequestHandler) fasthttp.RequestHandler) fasthttp.RequestHandler {
	// load middlewares.
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	// return loded middleware's.
	return handler
}

// this method maded for run other middlewares.
func ChainCtx(handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) RequestHandler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
