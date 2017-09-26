package models

type Balance struct {
	Amount  Money `json:"amount"`
	Updated Time  `json:"updated"`
}
