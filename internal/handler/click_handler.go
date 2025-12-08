package handler

import (
	"net"
	"net/http"
	"strings"

	"github.com/rianlucas/url-shortener/internal/service"
)

type ClickHandler struct {
	clickService *service.ClickService
	urlService   *service.UrlService
}

func NewClickHandler(clickService *service.ClickService, urlService *service.UrlService) *ClickHandler {
	return &ClickHandler{
		clickService: clickService,
		urlService:   urlService,
	}
}

func getIP(r *http.Request) string {
	// Primeira tentativa: X-Forwarded-For
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Pode ter vários IPs, pegamos o primeiro
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Segunda tentativa: X-Real-IP
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fallback: RemoteAddr
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}
