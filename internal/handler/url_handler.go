package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/service"
)

type UrlHandler struct {
	Service *service.UrlService
}

func CreateUrlHandler(urlService *service.UrlService) *UrlHandler {
	return &UrlHandler{Service: urlService}
}

func (u *UrlHandler) Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"message": fmt.Sprintf("O método %s não é permitido para essa rota. Utilize POST", r.Method),
		})
		return
	}

	var urlDto dto.CreateUrlDto
	err := json.NewDecoder(r.Body).Decode(&urlDto)
	if err != nil {
		fmt.Println(err)
	}

	result := u.Service.Create(urlDto)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}
