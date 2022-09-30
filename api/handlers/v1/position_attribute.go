package handlers

import (
	"api-gateway/api/models"
	"api-gateway/genproto/profession_service"
	"api-gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePositionAttribute godoc
// @ID create-position-attribute
// @Router /v1/position_attribute [POST]
// @Summary create postion attribute
// @Description Create Position Atttibute
// @Tags position_attribute
// @Accept json
// @Produce json
// @Param position_attribute body models.CreatePA true "position_attribute"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreatePositionAttribute(c *gin.Context) {
	var pa models.CreatePA
	if err := c.BindJSON(&pa); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input position attribute", errors.New("wrong input position attribute"))
		return
	}

	resp, err := h.services.PositionAttributeService().Create(
		context.Background(),
		&profession_service.CreatePositionAttribute{
			Value:       pa.Value,
			PositionId:  pa.PositionId,
			AttributeId: pa.AttributeId,
		},
	)

	if !handleError(h.log, c, err, "error while creating position attribute") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// GerAllPositionAttribute godoc
// @ID get-all-position-attribute
// @Router /v1/position_attribute [GET]
// @Summary get all position attribute
// @Description Get All Position Attribute
// @Tags position_attribute
// @Accept json
// @Produce json
// @Param value query string false "value"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllPositionAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllPositionAttribute(c *gin.Context) {
	//var arrPA models.GetAllPositionAttributeResponse
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionAttributeService().GetAll(
		context.Background(),
		&profession_service.GetAllPositionAttributeRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Value:  c.Query("value"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all position attributes") {
		return
	}

	/*err = ParseToStruct(&arrPA, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while parsing", err)
	}*/

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// GetPositionAttribute godoc
// @ID get-position-attribute
// @Router /v1/position_attribute/{position_attribute_id} [GET]
// @Summary get position attribute
// @Description Get Position Attribute
// @Tags position_attribute
// @Accept json
// @Produce json
// @Param position_attribute_id path string false "position_attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.GetPositionAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetPositionAttribute(c *gin.Context) {
	//var position models.GetPositionAttributeResponse
	id := c.Param("position_attribute_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong position attribut id", errors.New("wrong position attribut id"))
		return
	}

	resp, err := h.services.PositionAttributeService().GetById(
		context.Background(),
		&profession_service.PositionAttributeId{
			Id: id,
		},
	)

	if !handleError(h.log, c, err, "error while parsing") {
		return
	}

	//err = ParseToStruct(&position, resp)
	/*if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while parsing", err)
	}*/

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// UpdatePositionAttribute godoc
// @ID update-position-attribute
// @Router /v1/position_attribute [PUT]
// @Summary update position attribute
// @Description Update Position Attribute
// @Tags position_attribute
// @Accept json
// @Produce json
// @Param position_attribute body models.PositionA false "position_attribute"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdatePositionAttribute(c *gin.Context) {
	var positionAttribute models.PositionA
	if err := c.BindJSON(&positionAttribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input position attribute", errors.New("wrong input position attribute"))
		return
	}
	_, err := h.services.PositionAttributeService().Update(
		context.Background(),
		&profession_service.PositionAttribute{
			Id:          positionAttribute.Id,
			PositionId:  positionAttribute.PositionId,
			AttributeId: positionAttribute.AttributeId,
			Value:       positionAttribute.Value,
		},
	)
	if !handleError(h.log, c, err, "error while updating position attribute") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", models.MsgModel{Msg: "Updated"})
}

// DeletePositionAttribute godoc
// @ID delete-position-attribute
// @Router /v1/position_attribute/{position_attribute_id} [DELETE]
// @Summary delete position attribute
// @Description Delete Position Attribute
// @Tags position_attribute
// @Accept json
// @Produce json
// @Param position_attribute_id  path string true "position_attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeletePositionAttribute(c *gin.Context) {
	id := c.Param("position_attribute_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong position attribute id ", errors.New("wrong position attribute id "))
		return
	}

	_, err := h.services.PositionAttributeService().Delete(
		context.Background(),
		&profession_service.PositionAttributeId{
			Id: id,
		},
	)

	if !handleError(h.log, c, err, "error while deleting position attribute ") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", models.MsgModel{Msg: "Deleted"})

}
