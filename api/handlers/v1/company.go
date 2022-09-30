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

// CreateCompany godoc
// @ID create-company
// @Router /v1/company [POST]
// @Summary create company
// @Description Create Company
// @Tags company
// @Accept json
// @Produce json
// @Param company body models.CreateCompany true "company"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreateCompany(c *gin.Context) {
	var createCompany models.CreateCompany
	if err := c.BindJSON(&createCompany); err != nil {
		h.handleErrorResponse(c, 400, "error wrong input", err)
		return
	}

	resp, err := h.services.CompanyService().Create(
		context.Background(),
		&position_service.CreateCompany{
			Name: createCompany.Name,
		},
	)
	if !handleError(h.log, c, err, "error creating  company ") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}

// GetAllCompany godoc
// @ID get-all-company
// @Router /v1/company [GET]
// @Summary get all company
// @Description Get All Company
// @Tags company
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllCompanyResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllCompany(c *gin.Context) {
	var getAllcompany models.GetAllCompanyResponse

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.CompanyService().GetAll(
		context.Background(),
		&position_service.GetAllCompanyRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Name:   c.Query("name"),
		},
	)

	if !handleError(h.log, c, err, "error while getting all company") {
		return
	}

	err = ParseToStruct(&getAllcompany, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while parsing", err)
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", getAllcompany)

}

// GetCompany godoc
// @ID get-company
// @Router /v1/company/{company_id} [GET]
// @Summary get company
// @Description Get Company
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string false "company_id"
// @Success 200 {object} models.ResponseModel{data=models.Company} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetCompany(c *gin.Context) {
	var company models.Company
	id := c.Param("company_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong company_id input ", errors.New("wrong company_id input"))
	}

	resp, err := h.services.CompanyService().GetById(
		context.Background(),
		&position_service.CompanyId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error getting company id") {
		return
	}

	err = ParseToStruct(&company, resp)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error parsing", err)
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", company)
}

// UpdateCompany godoc
// @ID update-company
// @Router /v1/company [PUT]
// @Summary update company
// @Description Update Company
// @Tags company
// @Accept json
// @Produce json
// @Param company body models.Company true "company"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdateCompany(c *gin.Context) {
	var updateCompany models.Company
	if err := c.BindJSON(&updateCompany); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong update company input", err)
	}

	_, err := h.services.CompanyService().Update(
		context.Background(),
		&position_service.Company{
			Id:   updateCompany.Id,
			Name: updateCompany.Name,
		},
	)

	if !handleError(h.log, c, err, "error update company ") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", models.MsgModel{Msg: "updated"})
}

// DeleteCompany godoc
// @ID delete-company
// @Router /v1/company/{company_id} [DELETE]
// @Summary delete company
// @Description Delete Company
// @Tags company
// @Accept json
// @Produce json
// @Param company_id path string true "company_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteCompany(c *gin.Context) {
	id := c.Param("company_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong company id input", errors.New("wrong company id input"))
	}
	_, err := h.services.CompanyService().Delete(
		context.Background(),
		&position_service.CompanyId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error deleting company") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", models.MsgModel{Msg: "Deleted"})

}
