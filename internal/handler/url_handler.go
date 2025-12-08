package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/medama-io/go-useragent"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/service"
)

type UrlHandler struct {
	Service      *service.UrlService
	clickService *service.ClickService
}

func NewUrlHandler(urlService *service.UrlService, clickService *service.ClickService) *UrlHandler {
	return &UrlHandler{
		Service:      urlService,
		clickService: clickService,
	}
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
	_, err = u.Service.Update(url)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	var createClickDto dto.CreateClickDto

	ip := getIP(r)
	uaString := r.Header.Get("User-Agent")
	uaParser := useragent.NewParser()
	ua := uaParser.Parse(uaString)

	createClickDto.UrlId = url.ID
	createClickDto.Ip = ip
	createClickDto.Browser = ua.Browser().String()
	createClickDto.Os = ua.OS().String()

	_, err = u.clickService.Create(&createClickDto)
	if err != nil {
		fmt.Println("erro: ", err)
	}

	http.Redirect(w, r, url.LongUrl, 302)
}

func (u *UrlHandler) ShowQrCode(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/qr-code/")

	if shortCode == "" {
		log.Println("[QR-CODE] ERROR: ShortCode is empty!")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "shortCode is required in URL path. Use /qr-code/{shortCode}",
		})
		return
	}

	log.Printf("[QR-CODE] Searching for URL with shortCode: %s", shortCode)
	url, err := u.Service.FindByShortCode(shortCode)
	if err != nil {
		log.Printf("[QR-CODE] ERROR: URL not found for shortCode '%s': %v", shortCode, err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": fmt.Sprintf("URL not found: %v", err),
		})
		return
	}

	// Construir URL base dinamicamente a partir do request
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	baseURL := fmt.Sprintf("%s://%s", scheme, r.Host)
	shortURL := fmt.Sprintf("%s/%s", baseURL, url.ShortCode)

	log.Printf("[QR-CODE] URL found: %s - Generating QR Code for: %s", url.LongUrl, shortURL)
	qrcode, err := u.Service.GenerateQrCode(shortURL)
	if err != nil {
		log.Printf("[QR-CODE] ERROR: Failed to generate QR Code: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": fmt.Sprintf("failed to generate QR code: %v", err),
		})
		return
	}

	log.Printf("[QR-CODE] QR Code generated successfully for shortCode: %s", shortCode)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Write(qrcode)
}
