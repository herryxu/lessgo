package day1

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func Test_main(t *testing.T) {
	r := New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	r.POST("/say", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	if err := r.Run(":9090"); err != nil {
		fmt.Printf("server run faild.err:%s", err)
	}
	fmt.Printf("listening 9090")

}
func firstGet(c *context.Context) {

}
