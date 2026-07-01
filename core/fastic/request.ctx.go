package fastic

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
