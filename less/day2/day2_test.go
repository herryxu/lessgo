package day2

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func Test_main(t *testing.T) {
	r := NewEngine()
	r.GET("/", func(c *Context) {
		c.HTML(http.StatusOK, "root API Success!")
	})

	r.GET("/hello", func(c *Context) {
		c.String(http.StatusOK, "hello boy this is name:%s,and path:%s", c.Query("name"), c.Path)
	})
	r.POST("/say", func(c *Context) {
		c.JSON(http.StatusOK, Gf{
			"user": c.PostForm("name"),
		})
	})

	if err := r.Run(":9090"); err != nil {
		fmt.Printf("server run faild.err:%s", err)
	}
	fmt.Printf("listening 9090")

}
func firstGet(c *context.Context) {

}

func firstPost(c *context.Context) {

}
