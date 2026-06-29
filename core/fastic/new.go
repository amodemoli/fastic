// ▲ THIS FILE CODES THEY ARE ONLY EXAMPLE, THEY ARE NOT GOOD FOR REAL USE! PLEASE DONT USE THIS MIDDLEWARES FOR REAL PROJECT
package fastic

import (
	"fmt"
	"strings"

	"github.com/amodemoli/fastic/core/color"
	"github.com/fasthttp/router"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

// app structure example, you can add router and templates section!
type App struct {
	// router of fasthttp (if user need to other methods)
	Router *router.Router
	// handler for fasthttp, need this for run server. and load middlewares. =D
	Handler fasthttp.RequestHandler
	// router for env values (cache for speed responses)
	Env *Env
	// mdw for load chain middlewares and more =D
	Mdw *Mdw
}

// this function maded for create new application model =D
func New() *App {
	godotenv.Load()
	r := router.New()
	return &App{
		Router:  r,         // save created router
		Handler: r.Handler, // save fasthttp request handler
		Env:     LoadEnv(), // create new struct model for env laod
		Mdw:     MdwNew(),  // create new struct model for middlewares
	}
}

// = - = - = - = - = - = - = - = - Request Methods = - = - = - = - = - = - = - = -
// [TIPS] if you need speed web application you can use application.Router.GET, if you need easy coding use application.GET =D
// GET method example, you can use application.GET or application.Router.GET on your code.
// [UPDATE] updated methods model with ctx because added cutsom ctx for requests
func (a *App) GET(path string, handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) {

	finalHandler := handler // create new value for final hanlder
	if len(middlewares) > 0 {
		finalHandler = ChainCtx(handler, middlewares...) // run the handlers on chainctx
	}

	a.Router.GET(path, func(ctx *fasthttp.RequestCtx) {
		c := acquireCtx(ctx) // get ctx from pool
		defer releaseCtx(c)  // use defer if application get's panic, value put to pool
		finalHandler(c)      // run the handler with ctx
	})
}

// POST method example, you can use application.POST or application.Router.POST on your code.
func (a *App) POST(path string, handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) {

	finalHandler := handler // create new value for final hanlder
	if len(middlewares) > 0 {
		finalHandler = ChainCtx(handler, middlewares...) // run the handlers on chainctx
	}

	a.Router.POST(path, func(ctx *fasthttp.RequestCtx) {
		c := acquireCtx(ctx) // get ctx from pool
		defer releaseCtx(c)  // use defer if application get's panic, value put to pool
		finalHandler(c)      // run the handler with ctx
	})
}

// PUT method example, you can use application.PUT or application.Router.PUT on your code.
func (a *App) PUT(path string, handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) {

	finalHandler := handler // create new value for final hanlder
	if len(middlewares) > 0 {
		finalHandler = ChainCtx(handler, middlewares...) // run the handlers on chainctx
	}

	a.Router.PUT(path, func(ctx *fasthttp.RequestCtx) {
		c := acquireCtx(ctx) // get ctx from pool
		defer releaseCtx(c)  // use defer if application get's panic, value put to pool
		finalHandler(c)      // run the handler with ctx
	})
}

// DELETE method example, you can use application.DELETE or application.Router.DELETE on your code.
func (a *App) DELETE(path string, handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) {

	finalHandler := handler // create new value for final hanlder
	if len(middlewares) > 0 {
		finalHandler = ChainCtx(handler, middlewares...) // run the handlers on chainctx
	}

	a.Router.DELETE(path, func(ctx *fasthttp.RequestCtx) {
		c := acquireCtx(ctx) // get ctx from pool
		defer releaseCtx(c)  // use defer if application get's panic, value put to pool
		finalHandler(c)      // run the handler with ctx
	})
}

// OPTIONS method example, you can use application.OPTIONS or application.Router.OPTIONS on your code.
func (a *App) OPTIONS(path string, handler RequestHandler, middlewares ...func(RequestHandler) RequestHandler) {

	finalHandler := handler // create new value for final hanlder
	if len(middlewares) > 0 {
		finalHandler = ChainCtx(handler, middlewares...) // run the handlers on chainctx
	}

	a.Router.OPTIONS(path, func(ctx *fasthttp.RequestCtx) {
		c := acquireCtx(ctx) // get ctx from pool
		defer releaseCtx(c)  // use defer if application get's panic, value put to pool
		finalHandler(c)      // run the handler with ctx
	})
}

// = - = - = - = - = - = - = - = - Other Methods = - = - = - = - = - = - = - = -
// router group for creating groups and use custom middlewares for groups.
// u can use application.Router.Group() or application.Group()
func (a *App) Group(path string) *router.Group {
	return a.Router.Group(path)
}

// static router method maded for serve html,css ... files.
// u can use application.Static or application.Router.ServeFiles
func (a *App) Static(prefix, root string) {
	// add "/" prefix to file path. for debugging
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}

	// add /{filepath:*} suffix to file =D
	if !strings.HasSuffix(prefix, "/{filepath:*}") {
		prefix = prefix + "/{filepath:*}"
	}

	a.Router.ServeFiles(prefix, root) // serveFiles
}

// the print function, maded for print your string on terminal =D
func (a *App) Print(clr, model, message string) {
	fmt.Printf("    %s∎%s %s: %s\n", clr, color.Nc, model, message) // send a message on terminal.
}
