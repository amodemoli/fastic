// ▲ THIS FILE CODES THEY ARE ONLY EXAMPLE, THEY ARE NOT GOOD FOR REAL USE! PLEASE DONT USE THIS MIDDLEWARES FOR REAL PROJECT
package main

import (
	"github.com/amodemoli/fastic/core/fastic"
)

// example one, auth middleware
func exampleAuthMiddleware(next fastic.RequestHandler) fastic.RequestHandler {
	return func(c *fastic.Ctx) {
		login := false // example login status. [THIS IS FAKE]
		if !login {
			c.RawJSON(`{"error": "please login to account first"}`)
			return // user need to login cannot join to this section
		}

		next(c) // call next function (middleware or real function)
	}
}

// example two, admin auth middleware
func exampleAdminAuthMiddleware(next fastic.RequestHandler) fastic.RequestHandler {
	return func(c *fastic.Ctx) {
		adminPerm := true // example admin perm. [THIS IS FAKE]
		if !adminPerm {
			c.RawJSON(`{"error": "please login to account first"}`)
			return // user is not admin cannot enter. return from request
		}

		next(c) // call next function (middleware or real function)
	}
}
