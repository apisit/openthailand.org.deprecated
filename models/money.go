package models

import (
	"fmt"
	"strconv"
)

type Money struct {
	Amount          int64  `json:"amount"`
	Locale          Locale `json:"locale"`
	FormattedAmount string `json:"formatted_amount"`
}

func (m *Money) FactoredAmount() float64 {
	return float64(float64(m.Amount) / m.Locale.CurrencyFactor)
}

func (m Money) FormattedString() string {
	return fmt.Sprintf("%s%s", m.Locale.CurrencySymbol, m.FactoredAmountString())
}

func (m Money) FactoredAmountString() string {
	if m.Amount == 0 {
		return strconv.FormatFloat(float64(0.00), 'f', m.Locale.DefaultFractionDigits, 64)
	}
	return strconv.FormatFloat(float64(float64(m.Amount)/m.Locale.CurrencyFactor), 'f', m.Locale.DefaultFractionDigits, 64)
}

func (m *Money) FormattedAmountWithLocale(locale Locale) string {
	formattedAmount := strconv.FormatFloat(float64(float64(m.Amount)/locale.CurrencyFactor), 'f', locale.DefaultFractionDigits, 64)
	return fmt.Sprintf("%s%s", locale.CurrencySymbol, formattedAmount)
}

func NewMoneyWithAmountAndLocale(amount int64, locale Locale) Money {
	formattedAmount := strconv.FormatFloat(float64(float64(amount)/locale.CurrencyFactor), 'f', locale.DefaultFractionDigits, 64)
	return Money{
		Amount:          amount,
		Locale:          locale,
		FormattedAmount: fmt.Sprintf("%s%s", locale.CurrencySymbol, formattedAmount),
	}
}
