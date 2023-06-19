package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/katerji/UserAuthKit/db"
	"github.com/katerji/UserAuthKit/db/query"
	"github.com/katerji/UserAuthKit/input"
	"github.com/katerji/UserAuthKit/model"
)

type URLService struct{}

func (service URLService) CreateURL(input input.URLInput) (int, error) {
	return db.GetDbInstance().Insert(query.InsertURL, input.UserID, input.ShortURL, input.OriginalURL)
}

func (service URLService) DeleteURL(urlID int) bool {
	return db.GetDbInstance().Exec(query.DeleteURL, urlID)
}

func (service URLService) GetURLByShortURL(shortURL string) (model.URL, error) {
	var url model.URL
	row := db.GetDbInstance().QueryRow(query.FetchURL, shortURL)
	err := row.Scan(&url.ID, &url.UserID, &url.ShortURL, &url.OriginalURL)
	if err != nil {
		return model.URL{}, err
	}

	return url, nil
}

func (service URLService) GetUserURLs(userID int) ([]model.URL, error) {
	row, err := db.GetDbInstance().Query(query.FetchUserURLs, userID)
	if err != nil {
		return []model.URL{}, err
	}
	var urls []model.URL
	for row.Next() {
		url := model.URL{}
		err := row.Scan(&url.ID, &url.UserID, &url.ShortURL, &url.OriginalURL)
		if err != nil {
			return []model.URL{}, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (service URLService) GetFullShortURL(c *gin.Context, shortURL string) string {
	serverURL := getServerURL(c)
	return fmt.Sprintf("%s/api/r/%s", serverURL, shortURL)
}
func getServerURL(c *gin.Context) string {
	host := c.Request.Host
	protocol := "http"
	if c.Request.TLS != nil {
		protocol = "https"
	}
	serverURL := protocol + "://" + host
	return serverURL
}
