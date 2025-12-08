package service

import (
	"context"
	"fmt"

	"github.com/rianlucas/url-shortener/internal/clients/geo_localization/ipapi"
	"github.com/rianlucas/url-shortener/internal/database/repositories"
	"github.com/rianlucas/url-shortener/internal/dto"
	"github.com/rianlucas/url-shortener/internal/models"
)

type ClickService struct {
	ctx        context.Context
	repository repositories.ClickAnalyticsRepository
}

func NewClickService(ctx context.Context, clickRepository *repositories.ClickAnalyticsRepository) *ClickService {
	return &ClickService{
		ctx:        ctx,
		repository: *clickRepository,
	}
}

func (c *ClickService) Create(clickDto *dto.CreateClickDto) (models.ClickAnalytics, error) {
	ipapi := ipapi.NewIpapi("https://ipapi.co")

	localDto, err := ipapi.FindByLocalIp(clickDto.Ip)
	if err != nil {
		fmt.Println("click_service.go l:30")
		fmt.Println("error: ", err)
		return models.ClickAnalytics{}, err
	}

	clickDto.City = localDto.City
	clickDto.Country = localDto.CountryName
	clickDto.Timezone = localDto.Timezone

	fmt.Printf("localDto %+v", localDto)
	result, err := c.repository.Create(clickDto)
	if err != nil {
		return models.ClickAnalytics{}, err
	}

	return result, nil
}
