package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/service"
)

type UrlHandler struct {
	Service *service.UrlService
}

func NewUrlHandler(urlService *service.UrlService) *UrlHandler {
	return &UrlHandler{Service: urlService}
}

func (u *UrlHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("the method %s is not supported by this route, use POST instead", r.Method),
		})
		return
	}

	var urlDto dto.CreateUrlDto
	err := json.NewDecoder(r.Body).Decode(&urlDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	if urlDto.LongUrl == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "longUrl is required",
		})
		return
	}

	result, err := u.Service.Create(urlDto)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("error while creating a Url: %v", err),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (u *UrlHandler) FindByShortCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	if shortCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "shortCode is required in URL path",
		})
		return
	}

	url, err := u.Service.FindByShortCode(shortCode)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	http.Redirect(w, r, url.LongUrl, 302)
}
