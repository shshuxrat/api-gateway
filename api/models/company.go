package models

type Company struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateCompany struct {
	Name string `json:"name"`
}

type GetAllCompanyResponse struct {
	Companys []Company `json:"companys"`
	Count    int32     `json:"count"`
}
