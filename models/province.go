package models

import "encoding/json"

type Province struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	NameThai string `json:"nameThai"`
}
type CustomProvince Province

func (r Province) MarshalJSON() ([]byte, error) {
	result, err := json.MarshalIndent(CustomProvince(r), "", "\t")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Province) ToBytes() ([]byte, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Province) FromBytes(b []byte) {
	if b == nil {
		return
	}

	err := json.Unmarshal(b, &r)
	if err != nil {
		return
	}
}
