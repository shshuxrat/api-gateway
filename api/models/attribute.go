package models

type Attribute struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateAttribute struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetAllAttributeResponse struct {
	Attributes []Attribute `json:"attributes"`
	Count      int32       `json:"count"`
}
