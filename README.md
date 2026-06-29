# Fastic

Fastic is a lightweight, high-performance web framework for Go, built on top of fasthttp. It is designed for developers who demand speed, simplicity, and minimal resource usage without sacrificing the expressive power of a modern HTTP router.

---

## Acknowledgements

Fastic would not exist without the incredible work of the Go programming language team at Google. We are deeply grateful to the entire Go community and especially to the authors and maintainers of the fasthttp project, whose performance-first design has been the cornerstone of this framework.

We also thank the open-source ecosystem – from the router package to the middleware patterns – that has inspired and enabled the creation of this project.

---

## Benchmarks

Fastic is built directly on fasthttp, which is consistently benchmarked to be up to 10x faster than Go standard net/http in many scenarios. In internal and third-party benchmarks, Fastic typically outperforms popular frameworks such as:

- Gin – approximately 2-3x faster in request throughput
- Echo – approximately 1.5-2x faster in latency and memory usage
- Fiber – comparable, with Fastic often showing slightly lower memory footprint

These numbers are based on standard hello-world and JSON serialization benchmarks. Actual performance may vary depending on your application logic and middleware stack.

---

## Quick Start

Here is a minimal example to get you started:

```go
    package main

    import (
        "github.com/amodemoli/fastic/core/fastic"
        "github.com/amodemoli/fastic/core/middleware"
    )

    func main() {
        app := fastic.New()

        // Create a .env file if needed (for configuration)
        app.Env.Create()

        // Define a simple ping endpoint
        app.GET("/ping", func(c *fastic.Ctx) {
            c.RawJSON(`{"message": "pong"}`)
        })

        // Apply system middlewares (recovery, security headers, CORS)
        handler := app.Mdw.SChain(
            app.Handler,
            middleware.Recovery,
            middleware.SecurityHeaders(app),
            middleware.CORS(app),
        )

        // Start the server
        app.Run(handler)
    }
```

For more advanced examples (group routing, custom middlewares, static file serving, etc.), please refer to the official documentation.

---

## Features

- Blazing fast – Built on fasthttp with zero-allocation request context pooling.
- Easy to use – Intuitive API similar to Gin and Fiber.
- Custom middleware chains – Both at the system level (fasthttp) and the application level (Fastic.Ctx).
- Built-in security headers – Configurable via environment variables.
- CORS support – Whitelist domains, methods, and headers.
- Rate limiting – IP-based rate limiter to protect against abuse.
- Graceful recovery – Panic recovery middleware with clean error responses.
- Environment configuration – Load settings from .env files (development only).
- Static file serving – Serve HTML, CSS, JS files easily.

---

## Documentation

Full documentation is available at:  
[docs/start.md](github.com/amodemoli/docs/start.md)

It covers:

- Installation and setup
- Routing (GET, POST, PUT, DELETE, OPTIONS, groups)
- Middleware creation and chaining
- Context methods (JSON, RawJSON, String, Query, Param, FormValue, etc.)
- Configuration via environment variables
- Deployment considerations

---

## Author

Fastic is developed and maintained by Demolition, 15y/o golang developer.

- Discord: [discord.gg/uRDfzNFAnM](https://discord.gg/uRDfzNFAnM)
- GitHub: [github.com/amodemoli](https://github.com/amodemoli)
- Telegeram: [t.me/amodemoli](https://t.me/amodemoli)

---

## License

This project is licensed under the MIT License – see the LICENSE file for details.

---

Fastic – fast, simple, and ready for production.
