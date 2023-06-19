package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/service"
)

const URLDeletePath = "/url"

type URLDeleteRequest struct {
	ID int `json:"id"`
}

type URLDeleteResponse struct {
	Success bool `json:"success"`
}

func URLDeleteHandler(c *gin.Context) {
	var request URLDeleteRequest
	if err := c.BindJSON(&request); err != nil {
		sendBadRequest(c)
		return
	}
	urlService := service.URLService{}
	success := urlService.DeleteURL(request.ID)
	sendJSONResponse(c, URLDeleteResponse{success})
}
