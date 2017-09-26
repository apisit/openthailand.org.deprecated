package models

import (
	"fmt"
	"strings"
)

type Locale struct {
	Code                  string  `json:"code,omitempty"`
	TimeZone              string  `json:"time_zone,omitempty"`
	CountryName           string  `json:"country_name,omitempty"`
	CountryCode           string  `json:"country_code,omitempty"`
	CurrencyCode          string  `json:"currency_code,omitempty"`
	CurrencySymbol        string  `json:"currency_symbol,omitempty"`
	CurrencyFactor        float64 `json:"currency_factor,omitempty,string"`
	DefaultFractionDigits int     `json:"default_fraction_digits,omitempty,string"`
}

func AvailableLocales() []Locale {
	listLocales := SupportedLocales()
	list := []Locale{}
	for _, v := range listLocales {
		list = append(list, v)
	}
	return list
}

func SupportedLocales() map[string]Locale {
	return map[string]Locale{
		strings.ToLower("US"): Locale{
			Code:                  "us",
			TimeZone:              "",
			CountryName:           "United State",
			CountryCode:           "US",
			CurrencyCode:          "usd",
			CurrencySymbol:        "$",
			CurrencyFactor:        100.0,
			DefaultFractionDigits: 2},
		strings.ToLower("JP"): Locale{
			Code:                  "JP",
			TimeZone:              "Japan",
			CountryName:           "Japan",
			CountryCode:           "JP",
			CurrencyCode:          "jpy",
			CurrencySymbol:        "¥",
			CurrencyFactor:        1.0,
			DefaultFractionDigits: 0},
		strings.ToLower("TH"): Locale{
			Code:                  "TH",
			TimeZone:              "Asia/Bangkok",
			CountryName:           "Thailand",
			CountryCode:           "TH",
			CurrencyCode:          "thb",
			CurrencySymbol:        "฿",
			CurrencyFactor:        100.0,
			DefaultFractionDigits: 2},
		strings.ToLower("CN"): Locale{
			Code:                  "CN",
			TimeZone:              "Asia/Shanghai",
			CountryName:           "China",
			CountryCode:           "CN",
			CurrencyCode:          "cny",
			CurrencySymbol:        "¥",
			CurrencyFactor:        100.0,
			DefaultFractionDigits: 2},
	}
}

func LocaleByCountryCode(countryCode string) Locale {
	listLocales := SupportedLocales()
	return listLocales[strings.ToLower(countryCode)]
}

func FormatCurrency(currencyCode string, amount float64) string {
	var locale Locale
	listLocales := SupportedLocales()
	for _, l := range listLocales {
		if strings.ToLower(l.CurrencyCode) == strings.ToLower(currencyCode) {
			locale = l
		}
	}
	return fmt.Sprintf("%s%.2f", locale.CurrencySymbol, amount/locale.CurrencyFactor)
}
