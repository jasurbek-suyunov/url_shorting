package handler

import (
	"github.com/SuyunovJasurbek/url_shorting/models"
	"github.com/gin-gonic/gin"
)

// Create Url
// @Summary  Create Url
// @Description  Create Url
// @Tags         URL
// @Accept       json
// @Produce      json
// @Success      201  {object}  models.UrlRequest "Create successful"
// @Response     400 {object}  models.Error "Bad request"
// @Response     401 {object}  models.Error "Unauthorized"
// @Failure  	 500  {object}  models.Error "Internal server error"
// @Router       /url	[post]
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

func (h *Handler) GetUrls(c *gin.Context) {}

func (h *Handler) GetSingleUrl(c *gin.Context) {}

func (h *Handler) DeleteUrl(c *gin.Context) {}

func (h *Handler) UpdateUrl(c *gin.Context) {}

func (h *Handler) GetOriginalUrl(c *gin.Context) {}
