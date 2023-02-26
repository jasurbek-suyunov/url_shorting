package handler

import (
	"net/http"

	"github.com/SuyunovJasurbek/url_shorting/helper"
	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/gin-gonic/gin"
)

// SignUp User
// @Summary  SignUp User
// @Description  SignUp User
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        SignUp body models.UserSignUpRequest  true  "SignUp"
// @Success      201  {object}  models.UserResponse "SignUp successful"
// @Response     400 {object}  models.Error "Bad request"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /auth/singup	[post]
func (h *Handler) SignUp(c *gin.Context) {

	// variable
	var user models.UserSignUpRequest

	// bind
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//checking password and confirm password
	if user.Password != user.ConfirmPassword {
		c.JSON(http.StatusBadRequest, models.Error{
			Error: "Password and Confirm Password does not match",
		})
		return
	}

	// create user
	err := h.services.CreateUser(c, &user)

	// check error
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	// return result if no error
	c.JSON(201, models.Message{
		Message: "User created successfully",
	})
}

// SignIn User
// @Summary  SignIn User
// @Description  SignIn UserUrlRequest
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        SignIn body models.UserLoginRequest  true  "SingIn"
// @Success      201  {object}  models.UserResponse "SignIn successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Response     404 {object}  models.Error "Not found"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /auth/signin	[post]
func (h *Handler) SignIn(c *gin.Context) {

	// variable
	var login models.UserLoginRequest

	// bind
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	// get user by username
	user, err := h.services.GetUserByUsername(c, login.Username)

	// check error
	if err != nil {
		c.JSON(404, models.Error{
			Error: "Invalid username: " + err.Error(),
		})
		return
	}

	// check password
	if !helper.CheckPassword(user.PasswordHash, login.Password) {
		c.JSON(404, models.Error{
			Error: "Invalid password",
		})
		return
	}

	// generate token
	param := &models.Token{
		UserId:    user.ID,
		UserAgent: c.Request.UserAgent(),
	}
	token := helper.GenerateJWT(param)
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("token", token, 36000, "/", "localhost", false, true)
	c.JSON(http.StatusOK, &models.LoginResponse{
		Data: &models.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		},
		Error: "",
		Code:  0,
	})
}

// Logout User
// @Security ApiKeyAuth
// @Summary  Logout User
// @Description  Logout User
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.DefaultResponse "Logout successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Response     404 {object}  models.Error "Not found"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /auth/signout	[post]
func (h *Handler) SignOut(c *gin.Context) {
	c.SetCookie("token", "", 0, "", "", false, false)

	c.JSON(http.StatusOK, models.DefaultResponse{
		Data:  "succes logout",
		Error: "",
		Code:  200,
	})
}
