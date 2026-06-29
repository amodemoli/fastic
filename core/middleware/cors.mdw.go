// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"strconv"

	"github.com/amodemoli/fastic/core/fastic"
	"github.com/valyala/fasthttp"
)

func CORS(application *fastic.App) func(fasthttp.RequestHandler) fasthttp.RequestHandler {
	// get allowed domains, with application.Env struct.
	vary := []byte("Origin")
	isWildcard, allowedOrigins := application.Env.IsWildcard, application.Env.AllowedDomains

	// get allowed header and allowed methods from ".env"
	// with default values (if value is "" returns default value)
	// [UPDATE] used byte for values for speedly set
	allowedMethods, allowedHeaders := []byte(application.Env.AllowedMethods), []byte(application.Env.AllowedHeaders)
	maxAge := []byte(strconv.Itoa(application.Env.MaxAge))

	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			// set headers Vary to orgin model. (:
			ctx.Response.Header.SetBytesV("Vary", vary)

			// set allowed methods to ".env" file value, with getListFromEnd function help. and default value
			ctx.Response.Header.SetBytesV("Access-Control-Allow-Methods", allowedMethods)
			// set allowed headers to ".env" file value, with getListFromEnd function help. and default value
			ctx.Response.Header.SetBytesV("Access-Control-Allow-Headers", allowedHeaders)
			// set allowed max age to ".env" file value, with getListFromEnd function help. and default value
			ctx.Response.Header.SetBytesV("Access-Control-Max-Age", maxAge) // change maxAge to string.

			// if allowed orgins is *, means im accept all domain sites
			if isWildcard {
				ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", []byte("*"))
			} else if len(allowedOrigins) > 0 { // else you need to write whitelisted domains.
				origin := ctx.Request.Header.Peek("Origin")
				if len(origin) > 0 && allowedOrigins[string(origin)] {
					ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", origin)
				} else if len(origin) > 0 {
					ctx.SetStatusCode(fasthttp.StatusForbidden) // set status code to 403.
					return
				}
			}

			// default OPTIONS Method
			if string(ctx.Method()) == "OPTIONS" {
				ctx.SetStatusCode(fasthttp.StatusOK)
				return
			}

			next(ctx) // call the request for lunch.
		}
	}
}
