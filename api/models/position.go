package models

type Position struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ProfessionId string `json:"profession_id"`
	CompanyId    string `json:"company_id"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}

type CreatePosition struct {
	Name         string `json:"name"`
	ProfessionId string `json:"profession_id"`
	CompanyId    string `json:"company_id"`
}

type GetAllPositionResponse struct {
	Positions []Position `json:"positions"`
	Count     int32      `json:"count"`
}

type GetPosition struct {
	PositionFull   Position `json:"position"`
	CompanyName    string   `json:"company"`
	ProfessionName string   `json:"profession"`
}
