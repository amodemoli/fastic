// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"github.com/amodemoli/fastic/core/fastic"
	"github.com/valyala/fasthttp"
)

// some security response headers, you can edit the values on .env file =D
func SecurityHeaders(application *fastic.App) func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	// reading values from ".env" file, with help's default value!
	// if value is "" valus sets to default value =D
	// [UPDATE] used byte for values for speedly set
	xframeOptions := []byte(application.Env.XFrameOptions)
	referrerPolicy := []byte(application.Env.ReferrerPolicy)
	contentSecurityPolicy := []byte(application.Env.ContentSecurityPolicy)
	strictTransportSecurity := []byte(application.Env.StrictTransportSecurity)

	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			// set values to response header,
			// some response's dont need to edit or remove! for example X-XSS-Protection and X-Content-Type-Options
			ctx.Response.Header.SetBytesV("X-Content-Type-Options", []byte("nosniff"))
			ctx.Response.Header.SetBytesV("X-Frame-Options", xframeOptions)
			ctx.Response.Header.SetBytesV("X-XSS-Protection", []byte("1; mode=block"))
			ctx.Response.Header.SetBytesV("Referrer-Policy", referrerPolicy)
			ctx.Response.Header.SetBytesV("Content-Security-Policy", contentSecurityPolicy)
			ctx.Response.Header.SetBytesV("Strict-Transport-Security", strictTransportSecurity)

			next(ctx) // call the next function (response or next middleware) C=
		}
	}
}
