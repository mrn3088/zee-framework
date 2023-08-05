package main

import (
	"net/http"
	"zee"
)

func main() {
	// engine
	r := zee.New()

	// add handlers
	r.GET("/", func(c *zee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Zee</h1>")
	})

	r.GET("/hello", func(c *zee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *zee.Context) {
		c.JSON(http.StatusOK, zee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

}
