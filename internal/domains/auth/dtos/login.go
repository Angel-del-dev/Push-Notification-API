package dtos

type LoginType struct {
	Application string `json:"application"`
	Key         string `json:"key"`
	Password    string `json:"password"`
}
