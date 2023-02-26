package main

import (
	"io"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()

  r.POST("/singup", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })
  r.POST("/login", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })
  r.POST("/url", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })
  r.GET("/url", func(c *gin.Context) {
    c.String(200, "Hello")
  })
  r.GET("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })
  r.DELETE("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })
  r.PUT("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })

  // Redirect url
  r.GET("/in/:url", func(c *gin.Context) {
    url := c.Param("url")
    c.String(200, "Hello %s", url)
  })

  return r
}
