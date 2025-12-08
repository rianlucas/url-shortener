package geolocalization

type GeoLocalizationInterface interface {
	FindLocalByIp(ip string)
}

type LocalDto struct {
	City        string `json:"city"`
	CountryName string `json:"country_name"`
	Timezone    string `json:"timezone"`
	Currency    string `json:"currency"`
	Languages   string `json:"languagues"`
}
