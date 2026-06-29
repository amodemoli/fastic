package main

import (
	"github.com/amodemoli/fastic/core/fastic"
	"github.com/amodemoli/fastic/core/middleware"
)

func hightPerformance() {
	app := fastic.New()

	app.GET("/ping", func(c *fastic.Ctx) {
		c.RawJSON(`{"message": "pong"}`) // i use RawJSON because RawJSON method is speedly from JSON method.
	})

	// if you dont need panic recovery you can remove it this section. no problem. =D
	handler := app.Mdw.SChain(
		app.Handler,
		middleware.Recovery, // i only use Recovery. i can use CORS and SecurityHeaders for security on Real Projects. (CORS and SecurityHeaders can make my site 15% slowwes)
	)

	app.Run(handler)
}

// YOU CAN USE THIS CODE ON BENCHMARKS AND START PERFORMANCE FIGHT TO OTHER FRAMEWORKS FOR EXAMPLE "Fiber" AND "Gin" And "Echo".
