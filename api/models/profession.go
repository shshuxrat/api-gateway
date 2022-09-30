package models

type Profession struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
type GetAllProfessionResponse struct {
	Professions []Profession `json:"professions"`
	Count       uint32       `json:"count"`
}

type CreateProfession struct {
	Name string `json:"name"`
}

type MsgModel struct {
	Msg string `json:"message"`
}
