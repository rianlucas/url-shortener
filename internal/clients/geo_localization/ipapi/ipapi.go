package ipapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	geolocalization "github.com/rianlucas/url-shortener/internal/clients/geo_localization"
)

type Ipapi struct {
	baseUrl string
	client  *http.Client
}

func NewIpapi(baseUrl string) *Ipapi {
	return &Ipapi{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (i *Ipapi) FindByLocalIp(ip string) (geolocalization.LocalDto, error) {
	var localDto geolocalization.LocalDto

	url := fmt.Sprintf("%s/%s/json", i.baseUrl, ip)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return geolocalization.LocalDto{}, fmt.Errorf("create request: %w", err)
	}

	res, err := i.client.Do(req)
	if err != nil {
		return geolocalization.LocalDto{}, fmt.Errorf("do request: %w", err)
	}
	defer res.Body.Close()

	// Lê o body inteiro pra poder logar independentemente do status
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return geolocalization.LocalDto{}, fmt.Errorf("read body: %w", err)
	}

	if res.StatusCode != 200 {
		return geolocalization.LocalDto{}, fmt.Errorf("unexpected error from ipapi API (status %d, body %s)", res.StatusCode, string(bodyBytes))
	}

	fmt.Println("=== IPAPI DEBUG ===")
	fmt.Println("URL:   ", url)
	fmt.Println("Status:", res.StatusCode)
	fmt.Println("Body:  ", string(bodyBytes))
	fmt.Println("===================")

	if len(bodyBytes) == 0 {
		return geolocalization.LocalDto{}, fmt.Errorf("empty response body from ipapi (status %d)", res.StatusCode)
	}

	// 🔄 Agora sim, tenta decodificar o JSON
	if err := json.Unmarshal(bodyBytes, &localDto); err != nil {
		return geolocalization.LocalDto{}, fmt.Errorf("decode ipapi json: %w", err)
	}

	return localDto, nil
}
