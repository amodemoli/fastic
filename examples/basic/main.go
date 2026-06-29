/*
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
|                              Fastic • v1.0
| Author: Demolition
| Discord: discord.gg/uRDfzNFAnM
| Source: github.com/amodemoli/fastic
| Info: fastic, just a fast web library for easy use fasthttp.
|
= - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - = - =
*/
// ▲ YOU CAN CUSTOMIZE YOUR WEBTIE SECURITY SETTINGS ON ".env" FILE, PLEASE DONT USE ".env" FILE ON REAL PROJECTS. USE PLATFORM ENV VALUE SAVEV,
// > IF YOU CAN'T SEE ".env" FILE YOU CAN USE app.Env.Create() METHOD FOR CREATE ".env" FILE EXAMPLE. TNX FOR USING <3
package main

import (
	"github.com/amodemoli/fastic/core/color"
	"github.com/amodemoli/fastic/core/fastic"
	"github.com/amodemoli/fastic/core/middleware"
)

func main() {
	app := fastic.New() // create new fastic application, and get application struct as app var.

	// create ".env" file with default template, only for see and edit server (please dont use ".env" file on real projects),
	// this method return's a boolian value, if finded ths file or writed on this file return's true.
	if !app.Env.Create() {
		// app print method maded for ptint a values on terminal. please dont use println, because app.Print is best for clearn ui =D
		app.Print(color.Red, "ERROR", `cannot create ".env" file with default value.`)
	}

	// create new path, named ping. with help's GET method
	app.GET("/ping", func(c *fastic.Ctx) {
		c.RawJSON(`{"message": "pong"}`) // u can use RawJSON, this method is fast. you can use c.JSON but you need create new map. (this problem fixes coming soon)
	})

	// create main path, only for show SendFile example.
	app.GET("/", func(c *fastic.Ctx) {
		c.SendFile("public/index.html") // load your html file with css and javascript sources.
		// if you need to connect frontend to backend please use gRPC or RestApi.
	})

	// you can create custom middleware for custom paths:
	app.GET("/dashboard", func(c *fastic.Ctx) {
		c.RawJSON(`{"message": "welcome to dashboard"}`)
	}, exampleAdminAuthMiddleware) // you can add your target middleware for run on this request.

	// example of +2 middlewares for one path =D
	app.GET("/admin", func(c *fastic.Ctx) {
		c.RawJSON(`{"message": "welcome to admin panel"}`)
	}, exampleAuthMiddleware, exampleAdminAuthMiddleware) // first running auth middleware and next running admin auth middleware.

	// create new middleware chain. (this code midllewares are running in all requests),
	// app.Mdw.SChain method maded for system middlewares, dont use this method for other middlewares.
	handler := app.Mdw.SChain(
		app.Handler, // you need to add this handler function.!
		// other middlewares are optinal you can remove.
		middleware.Recovery,      // this middleware maded for recovery and show server panics
		middleware.RequestLogger, // this middleware maded for log all requests and response times
		// this middleware need to application.
		middleware.SecurityHeaders(app), // this middleware maded for enable security headers on requests
		middleware.CORS(app),            // this middleware is created for request cors settings
		// You can customize "CORS" and "SecurityHeaders" setting on ".env" file.
		// [*READ] SOME MIDDLEWARES CAN SLOW YOUR SPEED RESPONSES, EXAMPLE: REQUEST LOGGER. I DONT USE THIS MIDDLEWARER BUT THIS MIDDLEWARE CAN SLOW YOUR SERVER TO 30% AND MORE...
		// * PLEASE DONT ADD SOME MIDDLEWARES ARE THERE, THIS SECTION MIDDLEWARES ARE RUNNING IN ALL REQUESTS. (CAN MAKE YOUR SERVER VERY SLOW)
	)

	// im used app.Run and dropped handler var to the method, if you dont need to use app.Mdw.Chain drop app.Handler to this method,
	// for example: app.Run(r.Handler)
	app.Run(handler)
}

// YOU CAN SEE A HIGTH PERFORMANCE CODE ON maxPerformance.example.go

// [*] DONT FORGET, FRAMEWORK'S DONT MAKE YOUR WEBSITE VERY SPEED, YOUR CODING PATTERN CAN MAKE YOUR WEBSITE SPEED.
// IF YOU THINK YOU CAN WRITE SPEED AND HIGTH PERFORMANCE DODE, AND YOU NEED A FAST WEBSITE AND FAST WEB FRAMEWORK, FASTIC IS A GOOD FRAMEWORK.
