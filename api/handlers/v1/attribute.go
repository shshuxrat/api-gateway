package handlers

import (
	"api-gateway/api/models"
	"api-gateway/genproto/position_service"
	"api-gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAttribute godoc
// @ID create-attribute
// @Router /v1/attribute [POST]
// @Summary create attribute
// @Description Create Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body models.CreateAttribute true "attribute"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreateAttribute(c *gin.Context) {
	var createAttribute models.CreateAttribute
	if err := c.BindJSON(&createAttribute); err != nil {
		h.handleErrorResponse(c, 400, "wrong input or type ", err)
		return
	}
	resp, err := h.services.AttributeService().Create(
		context.Background(),
		&position_service.CreateAttribute{
			Name: createAttribute.Name,
			Type: createAttribute.Type,
		},
	)

	if !handleError(h.log, c, err, "error while creating Attribute") {
		return
	}

	h.handleSuccessResponse(c, 200, "Created", resp)
}

// GetAllAttribute godoc
// @ID get-all-attribute
// @Router /v1/attribute [GET]
// @Summary get all attribute
// @Description Get All Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllAttributeResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllAttribute(c *gin.Context) {
	var attributes models.GetAllAttributeResponse
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}
	resp, err := h.services.AttributeService().GetAll(
		context.Background(),
		&position_service.GetAllAttributeRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error  getting all attributes") {
		return
	}

	err = ParseToStruct(&attributes, resp)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error parsing ", err)
		return
	}
	h.handleSuccessResponse(c, 200, "ok", attributes)
}

// GetAttribute godoc
// @ID get-attribute
// @Router /v1/attribute/{attribute_id} [GET]
// @Summary get attribute by id
// @Description Get Attribute By Id
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.Attribute} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAttribute(c *gin.Context) {
	var attribute models.Attribute
	id := c.Param("attribute_id")

	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute_id is not valid ", errors.New("attribute_id is not valid"))
		return
	}

	resp, err := h.services.AttributeService().GetById(
		context.Background(),
		&position_service.AttributeId{
			Id: id,
		},
	)

	if !handleError(h.log, c, err, "error getting attribute by id") {
		return
	}
	err = ParseToStruct(&attribute, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error parsing", err)
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", attribute)

}

// UpdateAttribute godoc
// @ID update-attribute
// @Router /v1/attribute [PUT]
// @Summary update attribute
// @Description Update Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute body models.Attribute true "attribute"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdateAttribute(c *gin.Context) {
	var attribute models.Attribute
	if err := c.BindJSON(&attribute); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input for attribute", err)
		return
	}

	_, err := h.services.AttributeService().Update(
		context.Background(),
		&position_service.Attribute{
			Id:   attribute.Id,
			Name: attribute.Name,
			Type: attribute.Type,
		},
	)
	if !handleError(h.log, c, err, "error Update attribute") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "Updated", models.MsgModel{Msg: "Updated"})
}

// DeleteAttribute godoc
// @ID delete-attribute
// @Router /v1/attribute/{attribute_id} [DELETE]
// @Summary delete attribute
// @Description Delete Attribute
// @Tags attribute
// @Accept json
// @Produce json
// @Param attribute_id path string true "attribute_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteAttribute(c *gin.Context) {
	attribute_id := c.Param("attribute_id")
	if !util.IsValidUUID(attribute_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "attribute_id is not valid ", errors.New("attribute_id is not valid"))
		return
	}

	_, err := h.services.AttributeService().Delete(
		context.Background(),
		&position_service.AttributeId{
			Id: attribute_id,
		},
	)
	if !handleError(h.log, c, err, "error while deleting attribute") {
		return
	}
	h.handleSuccessResponse(c, 200, "deteled", models.MsgModel{Msg: "Deleled"})

}
