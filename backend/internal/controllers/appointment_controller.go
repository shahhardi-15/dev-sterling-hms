package controllers

import (
	"sterling-hms/internal/services"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	service *services.AppointmentService
}

func NewAppointmentController(service *services.AppointmentService) *AppointmentController {
	return &AppointmentController{service: service}
}

func (c *AppointmentController) GetAll(ctx *gin.Context) {
	appointments, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointments)
}

func (c *AppointmentController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointment)
}

func (c *AppointmentController) GetByPatientID(ctx *gin.Context) {
	patientID := ctx.Param("id")
	appointments, err := c.service.GetByPatientID(ctx.Request.Context(), patientID)
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointments)
}

func (c *AppointmentController) GetByDoctorID(ctx *gin.Context) {
	doctorID := ctx.Param("id")
	appointments, err := c.service.GetByDoctorID(ctx.Request.Context(), doctorID)
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointments)
}

func (c *AppointmentController) Create(ctx *gin.Context) {
	var req services.CreateAppointmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	claims := ctx.MustGet("claims").(*utils.JWTClaims)

	appointment, err := c.service.Create(ctx.Request.Context(), req, claims.UserID)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Created(ctx, appointment)
}

func (c *AppointmentController) Approve(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := c.service.Approve(ctx.Request.Context(), id)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointment)
}

func (c *AppointmentController) Reject(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := c.service.Reject(ctx.Request.Context(), id)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointment)
}

func (c *AppointmentController) Cancel(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := c.service.Cancel(ctx.Request.Context(), id)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointment)
}

func (c *AppointmentController) Complete(ctx *gin.Context) {
	id := ctx.Param("id")
	appointment, err := c.service.Complete(ctx.Request.Context(), id)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, appointment)
}
