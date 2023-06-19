package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/service"
	"net/http"
)

const RouterPath = "/r/*param"

func RouterHandler(c *gin.Context) {
	urlString := c.Request.URL.Path
	lastIndex := len(urlString) - 1
	var shortURL string
	for i := lastIndex; i >= 0; i-- {
		if urlString[i] == '/' {
			shortURL = urlString[i+1:]
			break
		}
	}
	if shortURL == "" {
		sendBadRequest(c)
		return
	}
	urlService := service.URLService{}
	url, err := urlService.GetURLByShortURL(shortURL)
	if err != nil {
		sendErrorMessage(c, "Error fetching URL")
		return
	}
	c.Redirect(http.StatusMovedPermanently, url.OriginalURL)
	return
}
