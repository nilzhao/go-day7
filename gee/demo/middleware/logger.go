package middleware

import (
	"gee/gee"
	"log"
	"time"
)

func Logger(c *gee.Context) {
	// Start timer
	t := time.Now()
	// Process request
	c.Next()
	// Calculate resolution time
	log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
}
