package dto

type CreateUrlDto struct {
	LongUrl    string `json:"longUrl"`
	ShortCode  string `json:"shortCode"`
	AccessCode int    `json:"accessCount"`
}
