package main

import (
	"github.com/SuyunovJasurbek/url_shorting/middlewares"
	"github.com/SuyunovJasurbek/url_shorting/src/handler"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()

  public := r.Group("/")
  protected := r.Group("/")

  public.GET("/in/:url", handler.GetOriginalUrl)

  public.POST("/singup", handler.SignUp)
  public.POST("/login", handler.Login)

  // Protected routes

  protected.Use(middlewares.Auth("secret"))

  protected.POST("/url",  handler.CreateUrl)
  protected.GET("/url", handler.GetUrls)
  protected.GET("/url/:id", handler.GetSingleUrl)
  protected.DELETE("/url/:id", handler.DeleteUrl)
  protected.PUT("/url/:id", handler.UpdateUrl)

  return r
}
