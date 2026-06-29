// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package middleware

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/amodemoli/fastic/core/color"
	"github.com/valyala/fasthttp"
)

// this pool maded for log buffer
var logBufferPool = sync.Pool{
	New: func() interface{} { // create a new function, this function return one interface
		return &logMessage{} // return
	},
}

// log message structure, on this struct saves messages buf/status code/path/ip/elapsed time.
type logMessage struct {
	buf     byte   // buffer
	status  int    // status code (example: 500)
	path    string // path of request (example: /login)
	ip      string // user ip ->in localhost ip is 1.127.0.0
	elapsed time.Duration
}

// this channel maked for save logMessages. and get log messages from goroutins
// this cannel can save 10000 messages.
var logChannel = make(chan *logMessage, 10000)

// init function maded for write messages on terminal, help by =
// goroutins , buffers , pools
func init() {
	go func() { // create new goroutine
		for msg := range logChannel { // get range of messages
			var clr string // new value for save response code stats color
			switch msg.status {
			case 200:
				clr = color.Green
			case 500:
				clr = color.Red
			default:
				clr = color.Yellow
			}

			var buf [256]byte // create new buffer
			b := buf[:0]      // set b value to buffer :0

			// append messages to buffer.
			b = append(b, "    | «"...)
			b = append(b, clr...)
			b = strconv.AppendInt(b, int64(msg.status), 10)
			b = append(b, color.Nc...)
			b = append(b, "» "...)
			b = append(b, msg.path...)
			b = append(b, " ["...)
			b = append(b, msg.ip...)
			b = append(b, "] ["...)
			b = append(b, msg.elapsed.String()...)
			b = append(b, "]\n"...)
			// write appended messages with use os.Stdout.Write *speed
			os.Stdout.Write(b)

			// put the msg value to pool =D
			logBufferPool.Put(msg)
		}
	}()
}

// request logger, is a middleware maded for log request method/path and response speed. =D
func RequestLogger(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		start := time.Now()          // start a new timer for get response time
		next(ctx)                    // call next middlewares or target function
		elapsed := time.Since(start) // get response time of function

		msg := logBufferPool.Get().(*logMessage) // get logMessage with pool
		msg.status = ctx.Response.StatusCode()   // save response code to msg.status
		msg.path = string(ctx.Path())            // save url path to msg.path value
		msg.ip = ctx.RemoteIP().String()         // save user ip and change ip to string
		msg.elapsed = elapsed                    // save response time of function to msg.elapsed

		select { // select log channel and send data to channel
		case logChannel <- msg:
			// saved to message
		default:
			// if channel if full of message, i drop message to bufferpool, dont need print (for performance)
			logBufferPool.Put(msg)
		}
	}
}
