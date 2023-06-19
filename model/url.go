package model

type URL struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

func (url URL) ToOutput() URLOutput {
	return URLOutput{
		ID:          url.ID,
		ShortURL:    url.ShortURL,
		OriginalUrl: url.OriginalURL,
	}
}
