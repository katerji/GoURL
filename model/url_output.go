package model

type URLOutput struct {
	ID          int    `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalUrl string `json:"original_url"`
}
