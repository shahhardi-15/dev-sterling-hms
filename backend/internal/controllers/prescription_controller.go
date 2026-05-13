package controllers

import (
	"sterling-hms/internal/services"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
)

type PrescriptionController struct {
	service *services.PrescriptionService
}

func NewPrescriptionController(service *services.PrescriptionService) *PrescriptionController {
	return &PrescriptionController{service: service}
}

func (c *PrescriptionController) Create(ctx *gin.Context) {
	var req services.CreatePrescriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	claims := ctx.MustGet("claims").(*utils.JWTClaims)

	prescription, err := c.service.Create(ctx.Request.Context(), req, claims.UserID)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Created(ctx, prescription)
}

func (c *PrescriptionController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	prescription, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}
	utils.Success(ctx, prescription)
}

func (c *PrescriptionController) GetByPatientID(ctx *gin.Context) {
	patientID := ctx.Param("id")
	prescriptions, err := c.service.GetByPatientID(ctx.Request.Context(), patientID)
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, prescriptions)
}

func (c *PrescriptionController) Dispense(ctx *gin.Context) {
	id := ctx.Param("id")
	prescription, err := c.service.Dispense(ctx.Request.Context(), id)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, prescription)
}
