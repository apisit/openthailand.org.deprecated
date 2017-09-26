package api

import (
	"log"
	"net/http"
	"openthailand/errors"
	"openthailand/middleware"
	"openthailand/models"

	"openthailand/route"
	"openthailand/services"

	"github.com/gorilla/mux"
)

func Routes() []route.Config {
	return []route.Config{
		route.Config{
			Path:        "/provinces",
			Func:        Provinces,
			Middlewares: route.Use(middleware.JsonResult, middleware.GzipResult),
			Method:      route.GET,
		},
		route.Config{
			Path:        "/provinces/{id}",
			Func:        ProvinceDetail,
			Middlewares: route.Use(middleware.JsonResult, middleware.GzipResult),
			Method:      route.GET,
		},
		route.Config{
			Path:        "/provinces/{provinceID}/districts",
			Func:        DistrictsInProvince,
			Middlewares: route.Use(middleware.JsonResult, middleware.GzipResult),
			Method:      route.GET,
		},
		route.Config{
			Path:        "/provinces/{provinceID}/districts/{districtID}/subdistricts",
			Func:        SubdistrictsInDistrict,
			Middlewares: route.Use(middleware.JsonResult, middleware.GzipResult),
			Method:      route.GET,
		},
	}
}

func Hook(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("before")
		h(w, r)
		log.Printf("after")
	}
}

func ProvinceDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceID := vars["id"]
	province, err := services.NewCountryService().GetProvinceDetail(provinceID)
	if err != nil {
		errors.ToAppError(err).Write(w)
		return
	}
	response := models.NewResponse(province)
	response.Write(w)
}

//GET /provinces
func Provinces(w http.ResponseWriter, r *http.Request) {
	provinces, err := services.NewCountryService().GetProvinces()
	if err != nil {
		errors.ToAppError(err).Write(w)
		return
	}
	provinceResponse := models.NewResponse(provinces)
	provinceResponse.Write(w)
}

//GET /provinces/{provinceID}/districts
func DistrictsInProvince(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	provinceID := vars["provinceID"]
	list, err := services.NewCountryService().GetDistrictsByProvinceID(provinceID)
	if err != nil {
		errors.ToAppError(err).Write(w)
		return
	}
	response := models.NewResponse(list)
	response.Write(w)
}

//GET /provinces/{provinceID}/districts/{districtID}/subdistricts
func SubdistrictsInDistrict(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//provinceID := vars["provinceID"]
	districtID := vars["districtID"]
	list, err := services.NewCountryService().GetSubdistrictsByDistrictID(districtID)
	if err != nil {
		errors.ToAppError(err).Write(w)
		return
	}
	response := models.NewResponse(list)
	response.Write(w)
}
