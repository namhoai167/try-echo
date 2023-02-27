package models

type LoginResponse struct {
	Token string `json:"token"`
}

type SimpleResponse struct {
	Text string `json:"text"`
}
