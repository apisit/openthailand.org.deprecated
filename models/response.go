package models

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}
type Response struct {
	Code   int          `json:"code"`
	Result ResponseData `json:"result"`
}

func NewResponse(o interface{}) Response {
	return Response{Code: 200, Result: ResponseData{Data: o}}
}

type CustomResponse Response

func (r Response) MarshalJSON() ([]byte, error) {
	result, err := json.MarshalIndent(CustomResponse(r), "", "\t")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r Response) Write(w http.ResponseWriter) {
	result, _ := r.MarshalJSON()
	w.Write(result)
}

func (r *Response) ToBytes() ([]byte, error) {
	result, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Response) FromBytes(b []byte) {
	if b == nil {
		return
	}

	err := json.Unmarshal(b, &r)
	if err != nil {
		return
	}
}
