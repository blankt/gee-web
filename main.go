package main

import (
	"gee-web/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.GET("/hello", func(c *gee.Context) {
		c.StatusCode = 200
		c.JSON(http.StatusOK, gee.H{
			"name": "tlf",
		})
	})

	r.Run(":8080")
}
