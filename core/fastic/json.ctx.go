package fastic

import (
	"encoding/json"
	"fmt"
)

// this method maded for show your message as json message,
// chnages content website type to application/json
func (c *Ctx) JSON(data interface{}) error {
	c.SetContentType("application/json")   // change content type to application/json
	return json.NewEncoder(c).Encode(data) // encode your json response.
}

// RawJSON method helps you on write json response example: `{"message": "pong"}` this is easy and don't need to maps.
func (c *Ctx) RawJSON(rawJSON string) {
	c.SetContentType("application/json") // change content type to application/json
	c.Response.SetBodyString(rawJSON)    // write raw json
}

// Bind method maded for get json from request body (for rest api)
func (c *Ctx) Bind(v interface{}) error {
	if len(c.RequestCtx.Response.Body()) == 0 {
		return fmt.Errorf("request body is empty")
	}
	// read the request body and save body to interface.
	return json.Unmarshal(c.RequestCtx.Response.Body(), v) 
}
