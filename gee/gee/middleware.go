// 内置中间件

package gee

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func Logger(c *Context) {
	// Start timer
	t := time.Now()
	// Process request
	c.Next()
	// Calculate resolution time
	log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
}

func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s%d", file, line))
	}
	return str.String()
}

func Recovery(ctx *Context) {
	defer func() {
		if err := recover(); err != nil {
			message := fmt.Sprintf("%s", err)
			log.Printf("%s\n\n", trace(message))
			ctx.Fail("Internal Server Error")
		}
	}()
	ctx.Next()
}
