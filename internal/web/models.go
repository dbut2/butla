package web

type ShortenRequest struct {
	Url  string `json:"url"`
	Code string `json:"code"`
}

type ShortenResponse struct {
	Link string `json:"link"`
}
