package middleware

import (
	"gee/gee"
	"log"
	"time"
)

func LoggerOnlyForV2(c *gee.Context) {
	// Start timer
	t := time.Now()
	// if a server error occurred
	c.Fail("Internal Server Error")
	// Calculate resolution time
	log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
}
