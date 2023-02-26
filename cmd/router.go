package main

import (
  "io"

  "github.com/gin-gonic/gin"
  "github.com/SuyunovJasurbek/url_shorting/middlewares"
)

func setupRouter() *gin.Engine {
  r := gin.Default()

  public := r.Group("/")
  protected := r.Group("/")

  public.GET("/in/:url", func(c *gin.Context) {
    url := c.Param("url")
    c.String(200, "Hello %s", url)
  })

  public.POST("/singup", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })
  public.POST("/login", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })

  // Protected routes

  protected.Use(middlewares.Auth("secret"))

  protected.POST("/url", func(c *gin.Context) {
    body, _ := io.ReadAll(c.Request.Body)

    c.String(200, string(body))
  })
  protected.GET("/url", func(c *gin.Context) {
    c.String(200, "Hello")
  })
  protected.GET("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })
  protected.DELETE("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })
  protected.PUT("/url/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.String(200, "Hello %s", id)
  })

  return r
}
