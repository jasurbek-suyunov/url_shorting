package handler

import (
	"log"

	"github.com/SuyunovJasurbek/url_shorting/config"
	_ "github.com/SuyunovJasurbek/url_shorting/docs"
	"github.com/SuyunovJasurbek/url_shorting/middlewares"
	"github.com/SuyunovJasurbek/url_shorting/src/service"
	"github.com/SuyunovJasurbek/url_shorting/src/storage/postgres"
	"github.com/SuyunovJasurbek/url_shorting/src/storage/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// get configs
	cnf := config.NewConfig()
	r.Static("/uploads", "./uploads")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// get db
	db, err := postgres.NewPostgres(cnf)
	// check error
	if err != nil {
		log.Println(err)
	} else {
		log.Println("db connected")
	}

	// get redis
	redis, err := redis.NewRedisCache(cnf)
	// check error
	if err != nil {
		log.Println(err)
	} else {
		log.Println("redis connected")
	}

	// get service and handler
	services := service.NewService(db, redis)
	handler := NewHandler(services)

	// Routes
	r.GET("/ping", handler.Ping)

	// swagger
	swagger := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swagger))

	r.GET("", handler.API)
	r.GET("/:id", handler.GetUrl)

	app := r.Group("/api/v1")
	app.GET("", handler.API)

	auth := app.Group("auth")
	{
		auth.POST("singup", handler.SignUp)
		auth.POST("signin", handler.SignIn)
		auth.POST("signout", handler.SignOut)
	}

	app.Use(middlewares.Auth())
	url := app.Group("url")
	{
		url.POST("", handler.CreateUrl)
		url.GET("", handler.GetUrls)
		url.GET(":id", handler.GetUrlByID)
		url.DELETE(":id", handler.DeleteUrl)

	}

	return r
}
