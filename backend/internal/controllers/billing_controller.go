package controllers

import (
	"sterling-hms/internal/services"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
)

type BillingController struct {
	service *services.BillingService
}

func NewBillingController(service *services.BillingService) *BillingController {
	return &BillingController{service: service}
}

func (c *BillingController) GetAll(ctx *gin.Context) {
	billings, err := c.service.GetAll(ctx.Request.Context())
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, billings)
}

func (c *BillingController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	billing, err := c.service.GetByID(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}
	utils.Success(ctx, billing)
}

func (c *BillingController) Create(ctx *gin.Context) {
	var req services.CreateBillingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	billing, err := c.service.Create(ctx.Request.Context(), req)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Created(ctx, billing)
}

func (c *BillingController) MarkAsPaid(ctx *gin.Context) {
	id := ctx.Param("id")
	var req services.MarkPaidRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	billing, err := c.service.MarkAsPaid(ctx.Request.Context(), id, req)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, billing)
}
