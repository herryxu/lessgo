package main

import (
	"context"
	"fmt"
	"net/http"

	less "lessgo/less/day4"
)

func main() {
	r := less.NewEngine()
	r.GET("/index", func(c *less.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *less.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *less.Context) {
			// expect /hello?name=lessgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *less.Context) {
			// expect /hello/lessgo
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *less.Context) {
			c.JSON(http.StatusOK, less.Gf{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	if err := r.Run(":9090"); err != nil {
		fmt.Printf("server run faild.err:%s", err)
	}
	fmt.Printf("listening 9090")

}
func firstGet(c *context.Context) {

}

func firstPost(c *context.Context) {

}
