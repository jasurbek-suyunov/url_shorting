package handler

import (
	"log"

	"github.com/SuyunovJasurbek/url_shorting/config"
	_ "github.com/SuyunovJasurbek/url_shorting/docs"
	"github.com/SuyunovJasurbek/url_shorting/middlewares"
	"github.com/SuyunovJasurbek/url_shorting/src/service"
	"github.com/SuyunovJasurbek/url_shorting/src/storage/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// get configs
	cnf := config.NewConfig()

	// get db
	db, err := postgres.NewPostgres(cnf)

	// check error
	if err != nil {
		log.Println(err)
	}

	// get service
	services := service.NewService(db)
	handler := NewHandler(services)

	// swagger
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	public := r.Group("/")
	protected := r.Group("/")

	public.GET("/in/:url", GetOriginalUrl)

	public.POST("/singup", handler.SignUp)
	public.POST("/login", handler.Login)

	// Protected routes

	protected.Use(middlewares.Auth("secret"))

	protected.POST("/url", CreateUrl)
	protected.GET("/url", GetUrls)
	protected.GET("/url/:id", GetSingleUrl)
	protected.DELETE("/url/:id", DeleteUrl)
	protected.PUT("/url/:id", UpdateUrl)

	return r
}
