// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"fmt"
	"time"

	"github.com/amodemoli/fastic/core/color"
	"github.com/amodemoli/fastic/core/tools"
	"github.com/valyala/fasthttp"
)

// a function for run server. and show default run server message to user.
func (app *App) Run(h fasthttp.RequestHandler) {
	// this function clear's terminal old messages.
	// [TIPS] you can remove it.
	tools.ClearScreen()

	// server port, len is 4 and value is number, read value on ".env" file
	// example website open with port: http://localhost:$_PORT/
	// get port and change port return type as string To int, if port is invalid or cannot change string to int set port to 8000
	port := app.Env.Port

	// just a print for show started server message to user,
	// this message show server port and server link.
	fmt.Printf("\n Fastic • %s\n  %s▲%s Powered by fasthttp, maded by demolition\n  %s᳃%s Server served on http://localhost:%d\n", Version, color.Purple, color.Nc, color.Purple, color.Nc, port)
	// if cannot load .env file!
	// read developer mode stats for developer mode message
	tools.DevMessage(app.Env.DevelopemtMode, color.Yellow, "∎", "INFO", `server is running on development mode.`)
	if !app.Env.LoadedFile {
		if app.Env.AllowedMethods == "" {
			tools.DevMessage(app.Env.DevelopemtMode, color.Yellow, "∎", "WARNING", "No .env file found, using system environment variables.") // if cannot load ".env" file. show this warning message. with templates.DevMessage() function help
		}
	} else {
		tools.DevMessage(app.Env.DevelopemtMode, color.Green, "∎", "SUCCSESS", `".env" file loaded successfully.`) // if ".env" file loaded, show a message if development mode is enable! =D
	}

	// create server val, for customize server settings
	server := &fasthttp.Server{
		Handler:      h,                                                 // handler of request
		ReadTimeout:  time.Duration(app.Env.ReadTimeout) * time.Second,  // read timeout with seconds
		WriteTimeout: time.Duration(app.Env.WriteTimeout) * time.Second, // write timeout with seconds

		MaxConnsPerIP:      app.Env.MaxConnsPerIP, // max connection per ip number (int)
		MaxRequestBodySize: app.Env.MaxRequestBodySize,
		IdleTimeout:        time.Duration(app.Env.IdleTimeout) * time.Second,
		Concurrency:        app.Env.Concurrency,
		DisableKeepalive:   app.Env.DisableKeepalive,

		// enable this value for security, Dont change this please.
		NoDefaultServerHeader: true,
	}

	// this function run's server on custom port
	// get's port from ".env" file, and log errors on lunch server.
	if err := server.ListenAndServe(fmt.Sprintf(":%d", port)); err != nil {
		fmt.Printf("    %s∎%s Error: %v\n\n", color.Red, color.Nc, err) // show lunch error message on terminal
	}
}
