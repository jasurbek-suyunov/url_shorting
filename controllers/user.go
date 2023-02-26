package controllers

import (
  "encoding/json"
  "github.com/SuyunovJasurbek/url_shorting/models"
  "github.com/gin-gonic/gin"
)

func GetUrl(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "pong",
  })
}
func Login(c *gin.Context) {
  decoder := json.NewDecoder(c.Request.Body)

  form := struct{
    Email string `json:"email"`
    Password string `json:"password"`
  }{}
  err := decoder.Decode(&form)
  if err != nil {
    c.AbortWithError(400, err)
  }

  c.JSON(200, gin.H{
    "message": form,
  })
}
func SignUp(c *gin.Context) {
  decoder := json.NewDecoder(c.Request.Body)

  user := models.User{}
  err := decoder.Decode(&user)
  if err != nil {
    c.AbortWithError(400, err)
  }

  c.JSON(200, gin.H{
    "data": user,
  })
}


