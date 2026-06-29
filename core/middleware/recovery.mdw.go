// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"fmt"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/amodemoli/fastic/core/color"
	"github.com/valyala/fasthttp"
)

// just a midlleware maded for recover panics and deffence server from panic's and crash.
func Recovery(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() { // use deffer for recover panic's

			// built-in function named "recover" show's to you panic message
			// check the panic message if have e panic (panic message is't nil) they change status code to 500
			// and print the panic with style
			if r := recover(); r != nil {
				fullStack := string(debug.Stack())
				lines := strings.Split(fullStack, "\n")

				// create targetline, type is string
				var targetLine string
				for _, line := range lines { // get range of fullStack
					if strings.Contains(line, ".go:") && !strings.Contains(line, "runtime/") {
						targetLine = strings.TrimSpace(line)
						break
					}
				}

				if targetLine != "" {
					parts := strings.Split(targetLine, ":")
					if len(parts) >= 2 {
						filePath := parts[0]
						lineNumber := parts[1]
						// get line's of part's =D
						lineParts := strings.Split(lineNumber, " ")
						fileName := filepath.Base(filePath)
						targetLine = fmt.Sprintf("%s:%s", fileName, lineParts[0]) // append linename and filename
					}
				}

				// show the message
				fmt.Printf("    %s∎%s Panic: %v (%s)\n", color.Red, color.Nc, r, targetLine)

				// change content type of website and write JSON
				ctx.SetStatusCode(fasthttp.StatusInternalServerError)
				ctx.SetContentType("application/json")
				ctx.Response.SetBodyString(fmt.Sprintf(`{"error": "%v"}`, r))
			}
		}()
		next(ctx) // call request
	}
}
