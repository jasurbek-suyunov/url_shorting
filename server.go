package main

import (
  "github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()
  r.GET("/in/:url", func(c *gin.Context) {
    url := c.Param("url")
    c.String(200, "Hello %s", url)
  })
  return r
}
