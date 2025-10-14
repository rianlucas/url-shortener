package dto

type CreateUrlDto struct {
	LongUrl     string `json:"longUrl"`
	ShortCode   string `json:"shortCode,omitempty"`
	AccessCount int    `json:"accessCount,omitempty"`
}
