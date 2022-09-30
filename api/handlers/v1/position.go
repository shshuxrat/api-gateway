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

// CreatePosition godoc
// @ID create-position
// @Router /v1/position [POST]
// @Summary create postion
// @Description Create Position
// @Tags position
// @Accept json
// @Produce json
// @Param position body models.CreatePosition true "position"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreatePosition(c *gin.Context) {
	var createPosition models.CreatePosition
	if err := c.BindJSON(&createPosition); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input for position", errors.New("wrong input for position"))
		return
	}

	resp, err := h.services.PositionService().Create(
		context.Background(),
		&profession_service.CreatePosition{
			Name:         createPosition.Name,
			CompanyId:    createPosition.CompanyId,
			ProfessionId: createPosition.ProfessionId,
		},
	)

	if !handleError(h.log, c, err, "error while creating position ") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp.Id)
}

// GetAllPosition godoc
// @ID get-all-position
// @Router /v1/position [GET]
// @Summary get all position
// @Description Get All Position
// @Tags position
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllPositionResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllPosition(c *gin.Context) {
	//var positions models.GetAllPositionResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.PositionService().GetAll(
		context.Background(),
		&profession_service.GetAllPositionRequest{
			Name:   c.Query("name"),
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)

	if !handleError(h.log, c, err, "error while getting all positions") {
		return
	}

	//err = ParseToStruct(&positions, resp)

	/*if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while parsing ", err)
		return
	}*/

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetPosition godoc
// @ID get-position
// @Router /v1/position/{position_id} [GET]
// @Summary get position
// @Description Get Position
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string false "position_id"
// @Success 200 {object} models.ResponseModel{data=models.GetPosition} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetPosition(c *gin.Context) {
	//var position models.GetPosition
	id := c.Param("position_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong position_id", errors.New("wrong position_id"))
		return
	}

	resp, err := h.services.PositionService().GetById(
		context.Background(),
		&profession_service.PositionId{
			Id: id,
		},
	)

	if !handleError(h.log, c, err, "error while getting position") {
		return
	}

	/*err = ParseToStruct(&position, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing", err)
	}*/

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// UpdatePosition godoc
// @ID update-aposition
// @Router /v1/position [PUT]
// @Summary update position
// @Description Update Position
// @Tags position
// @Accept json
// @Produce json
// @Param position body models.Position true "position"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdatePosition(c *gin.Context) {

	var updatePosition models.Position
	if err := c.BindJSON(&updatePosition); err != nil {
		h.handleErrorResponse(c, 400, "wrong input for position update", err)
		return
	}

	_, err := h.services.PositionService().Update(
		context.Background(),
		&profession_service.Position{
			Id:           updatePosition.Id,
			Name:         updatePosition.Name,
			ProfessionId: updatePosition.ProfessionId,
			CompanyId:    updatePosition.CompanyId,
		},
	)
	if !handleError(h.log, c, err, "error Update attribute") {
		return
	}
	h.handleSuccessResponse(c, 200, "Updated", models.MsgModel{Msg: "Updated"})

}

// DeletePosition godoc
// @ID delete-position
// @Router /v1/position/{position_id} [DELETE]
// @Summary delete position
// @Description Delete Position
// @Tags position
// @Accept json
// @Produce json
// @Param position_id path string true "position_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeletePosition(c *gin.Context) {
	id := c.Param("position_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input position id", errors.New("wrong input position id"))
	}

	_, err := h.services.PositionService().Delete(
		context.Background(),
		&profession_service.PositionId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error while deleting position") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", models.MsgModel{Msg: "deleted"})

}
