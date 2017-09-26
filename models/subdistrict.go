package models

type Subdistrict struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	NameThai string   `json:"nameThai"`
	Location Location `json:"location"`
}
