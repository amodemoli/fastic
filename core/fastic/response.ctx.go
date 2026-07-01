package fastic

import (
	"fmt"
	"os"
	"path/filepath"
)

// this method maded for show your message as text/plain model,
// if your website content type is application/json this method changes to text/plain and show's raw message on website.
func (c *Ctx) String(s string) {
	c.SetContentType("text/plain") // change content type to text/plain
	c.WriteString(s)               // write user string.
}

// Status method maded for change status code of ctx and return new ctx to user
func (c *Ctx) Status(status int) *Ctx {
	c.SetStatusCode(status) // update page(ctx) status code
	return c                // return new ctx use optinal
}

// Attachment method, you can add your file path for download. if user enter to this addres your uploaded
// file path starting to download.
func (c *Ctx) Attachment(path string) error {
	if _, err := os.Stat(path); err != nil { // if cannot find the file
		return fmt.Errorf("error on read file: %w", err)
	}
	c.Response.Header.Set("Content-Disposition", "attachment; filename="+filepath.Base(path)) // change response header
	c.SendFile(path)                                                                          // send file.

	return nil
}
