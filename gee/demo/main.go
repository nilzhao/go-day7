package main

import (
	"fmt"
	"gee/demo/middleware"
	"gee/gee"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}

func main() {
	r := gee.New()
	r.Use(middleware.Logger) // global middleware
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Qi", Age: 30}
	stu2 := &student{Name: "Jun", Age: 20}

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "date",
			"now":   time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
