package models

type PositionA struct {
	Id          string `json:"id"`
	Value       string `json:"value"`
	AttributeId string `json:"attribute_id"`
	PositionId  string `json:"position_id"`
}

type CreatePA struct {
	Value       string `json:"value"`
	AttributeId string `json:"attribute_id"`
	PositionId  string `json:"position_id"`
}

type GetAllPositionAttributeResponse struct {
	PositionAttributes []PositionA `json:"position_attributes"`
	Count              int32       `json:"count"`
}

type GetPositionAttributeResponse struct {
	PositionAtt   PositionA `json:"position_attribute"`
	AttributeName string    `json:"attribute"`
	PositionName  string    `json:"position"`
}
