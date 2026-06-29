// ▲ PLEASE DONT CHANGE THIS CODE THIS IS A DEFAULT APP CODE, CHANGEING THIS CODE CAN BROKE YOUR SITE.
package fastic

import (
	"encoding/json"
	"sync"

	"github.com/valyala/fasthttp"
)

// handler function for fastic.
type RequestHandler func(*Ctx)

// create new structure named Ctx for create custom request ctx
// only for requests for fastic framework
type Ctx struct {
	*fasthttp.RequestCtx
}

var ctxPool = sync.Pool{
	New: func() interface{} {
		return &Ctx{}
	},
}

// new pool for Ctx named ctxPool, for speed and performance i use pool.

// this function maded for get ctx from pool and drop for
// my target function (dont need create new struct) i can create struct for first time and use struct for +1000 times =D
// only need to get struct from pool and drop it to pool after use.
func acquireCtx(ctx *fasthttp.RequestCtx) *Ctx {
	c := ctxPool.Get().(*Ctx) // get Ctx struct from pool
	c.RequestCtx = ctx        // save correct ctx to c.RequestCtx value =D
	return c
}

// this function maded for drop ctx after using =D
// only for use other times.
func releaseCtx(c *Ctx) {
	c.RequestCtx = nil // remove c.RequestCtx pointer for performance of GC (Garbge Colector)
	ctxPool.Put(c)     // back the value to pool
}

// = - = - = - = - = - = - = - = - HelperMethods On CTX = - = - = - = - = - = - = - = -

// this method maded for show your message as text/plain model,
// if your website content type is application/json this method changes to text/plain and show's raw message on website.
func (c *Ctx) String(s string) {
	c.SetContentType("text/plain") // change content type to text/plain
	c.WriteString(s)               // write user string.
}

// this method maded for show your message as json message,
// chnages content website type to application/json
func (c *Ctx) JSON(data interface{}) error {
	c.SetContentType("application/json")   // change content type to application/json
	return json.NewEncoder(c).Encode(data) // encode your json response.
}

func (c *Ctx) RawJSON(rawJSON string) {
	c.SetContentType("application/json") // change content type to application/json
	c.Response.SetBodyString(rawJSON)    // write raw json
}

// Status method maded for change status code of ctx and return new ctx to user
func (c *Ctx) Status(status int) *Ctx {
	c.SetStatusCode(status) // update page(ctx) status code
	return c                // return new ctx use optinal
}

// this method is for use ctx paramter easy, UserValue is too hard for daily use
// they are created easy method named Param
func (c *Ctx) Param(key string) string {
	val := c.UserValue(key)
	if str, ok := val.(string); ok {
		return str
	}
	return "" // if cannot find paramter return's nill string > ""
}

// this method maded for show paramter querys,
// example of querys: /example?name=value.
func (c *Ctx) Query(key string) string {
	return string(c.RequestCtx.QueryArgs().Peek(key))
}

// FormValue method maded for show POST method value's
func (c *Ctx) FormValue(key string) string {
	return string(c.RequestCtx.FormValue(key)) // get value and change response to string.
}

// this method maded for show you a request body =D
func (c *Ctx) Body() []byte {
	return c.RequestCtx.Request.Body()
}
