package day7

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
func Auth() HandlerFunc {
	return func(c *Context) {
		log.Printf("[%d]", c.StatusCode)
		c.Fail(501, "Internal Auth Failed")
	}
}
