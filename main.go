package main

import (
	"log"
	"net/http"
	"time"
	"zee"
)

func onlyForV2() zee.HandlerFunc {
	return func(c *zee.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := zee.New()
	r.Use(zee.Logger()) // global middleware
	r.GET("/", func(c *zee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello zee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *zee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
