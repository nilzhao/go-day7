package gee

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	Convey("解析模式串", t, func() {
		parts := parsePattern("/p/:name")
		So(parts, ShouldResemble, []string{"p", ":name"})

		parts = parsePattern("/p/*")
		So(parts, ShouldResemble, []string{"p", "*"})

		parts = parsePattern("/p/*name/*")
		So(parts, ShouldResemble, []string{"p", "*name"})
	})
}

func TestGetRoute(t *testing.T) {
	Convey("获取路由", t, func() {
		r := newTestRouter()
		n, params := r.getRoute("GET", "hello/qi")
		So(n, ShouldNotBeNil)
		So(n.pattern, ShouldEqual, "/hello/:name")
		So(params, ShouldNotBeNil)
		So(params["name"], ShouldEqual, "qi")
		fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, params["name"])
	})
}
