package services

import (
	"encoding/json"
	"log"
	"openthailand/cache"
	"openthailand/constant"
	"openthailand/database"
	"openthailand/models"
	"openthailand/repositories/country"
)

type CountryService struct {
	CountryServiceManager
}

type CountryServiceManager interface {
	GetProvinces() ([]models.Province, error)
	GetProvinceDetail(provinceID string) (models.Province, error)
	GetDistrictsByProvinceID(provinceID string) ([]models.District, error)
	GetSubdistrictsByDistrictID(districtID string) ([]models.Subdistrict, error)
}

func (c *CountryService) GetProvinces() ([]models.Province, error) {
	cacheManager := cache.NewCacheManager()
	v, found := cacheManager.Get(constant.CACHE_KEY_PROVINCES)
	if found {
		log.Printf("cache hit")
		result := []models.Province{}
		json.Unmarshal(v, &result)
		return result, nil
	}
	log.Printf("cache missed")
	provinces, err := c.CountryServiceManager.GetProvinces()
	v, _ = json.Marshal(provinces)
	cacheManager.Set(constant.CACHE_KEY_PROVINCES, v)
	return provinces, err
}

func NewCountryService() *CountryService {
	db := database.Connect()
	r := &country.DatabaseRepository{DB: db}
	return &CountryService{CountryServiceManager: r}
}
