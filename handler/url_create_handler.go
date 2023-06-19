package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/model"
	"github.com/katerji/UserAuthKit/service"
	"time"
)

const URLCreatePath = "/url"

type URLCreateRequest struct {
	OriginalURL string `json:"original_url"`
}

func URLCreateHandler(c *gin.Context) {
	var request URLCreateRequest
	if err := c.BindJSON(&request); err != nil {
		sendBadRequest(c)
		return
	}
	user := c.MustGet("user").(model.User)
	urlService := service.URLService{}
	urlInput := input.URLInput{
		OriginalURL: request.OriginalURL,
		UserID:      user.ID,
		ShortURL:    generateUniqueURL(user.ID),
	}
	urlID, err := urlService.CreateURL(urlInput)
	if err != nil {
		sendErrorMessage(c, "URL already exists.")
		return
	}
	serverURL := urlService.GetFullShortURL(c, urlInput.ShortURL)
	output := model.URLOutput{
		ID:          urlID,
		ShortURL:    serverURL,
		OriginalUrl: urlInput.OriginalURL,
	}
	response := map[string]model.URLOutput{
		"url": output,
	}
	sendJSONResponse(c, response)
}

func generateUniqueURL(userID int) string {
	currentTime := time.Now().Unix()
	return fmt.Sprintf("%d%d", userID, currentTime)
}
