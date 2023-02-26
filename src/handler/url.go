package handler

import (
	"net/http"

	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/gin-gonic/gin"
)
// Create Url
// @Security ApiKeyAuth
// @Summary  Create Url
// @Description  Create Url
// @Tags         URL
// @Accept       json
// @Produce      json
// @Param        url body models.UrlRequest true "Url"
// @Success      201  {object}  models.Url "Create successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /api/v1/url	[post]
func (h *Handler) CreateUrl(c *gin.Context) {

	var (
		url models.UrlRequest
	)
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	// create url
	url_result, err := h.services.CreateUrl(c, &url)
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	// return result if no error
	c.JSON(201, url_result)
}
// Create Url
// @Security ApiKeyAuth
// @Summary  Get Urls
// @Description  Get Urls
// @Tags         URL
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.GetAllUrl "GetUrls successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /api/v1/url	[get]
func (h *Handler) GetUrls(c *gin.Context) {
	urls, err := h.services.GetUrls(c)
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, urls)
}
// Get Url
// @Security ApiKeyAuth
// @Summary  Get Url
// @Description  Get Url
// @Tags         URL
// @Accept       json
// @Produce      json
// @Param        id query string  true "UrlId"
// @Success      200  {object}  models.Url "GetUrl successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /api/v1/url/{id}	[get]
func (h *Handler) GetUrlByID(c *gin.Context) {
	var (
		url models.GetUrlByIdRequest
	)
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	url_result, err := h.services.GetUrlByID(c, url.ID)
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, url_result)
}
// Delete Url
// @Security ApiKeyAuth
// @Summary  Delete Url
// @Description  Delete Url
// @Tags         URL
// @Accept       json
// @Produce      json
// @Param        id query string  true "UrlId"
// @Success      200  {object}   models.Message "Delete successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /api/v1/url/{id}	[delete]
func (h *Handler) DeleteUrl(c *gin.Context) {
	var (
		url models.DeleteUrlRequest
	)
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	err := h.services.DeleteUrl(c, url.ID)
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Message{
		Message: "Delete successful",
	})
}

func (h *Handler) GetUrl(c *gin.Context) {

	url := c.Param("id")

	// get url
	path, err := h.services.GetUrl(c, url)
	if err != nil {
		c.JSON(400, models.Error{
			Error: err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, path)
}
