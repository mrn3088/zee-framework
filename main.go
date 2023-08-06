package main

import (
	"net/http"
	"zee"
)

func main() {
	r := zee.New()

	r.GET("/", func(c *zee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello zee</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *zee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})
		v1.GET("/hello", func(c *zee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *zee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *zee.Context) {
			c.JSON(http.StatusOK, zee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
