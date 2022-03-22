package dto

type ApiResponse struct {
    Id       int     `json:"ID"`
	Name     string  `json:"NAME"`
	Contents  string `json:"CONTENTS"`
	Created  string  `json:"CREATED"`  
}

type ApiRequest  struct {
	Name     string   `json:"NAME"`
	Contents string   `json:"CONTENTS"`      
}

type ApisResponse  struct {
	Api   []ApiResponse   `json:"Users"`
}