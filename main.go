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

	r.GET("/hello", func(c *zee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *zee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *zee.Context) {
		c.JSON(http.StatusOK, zee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
