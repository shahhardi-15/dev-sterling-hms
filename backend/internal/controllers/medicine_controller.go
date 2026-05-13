package controllers

import (
	"sterling-hms/internal/repositories"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type MedicineController struct {
	repo *repositories.MedicineRepository
}

func NewMedicineController(db *sqlx.DB) *MedicineController {
	return &MedicineController{repo: repositories.NewMedicineRepository(db)}
}

func (c *MedicineController) GetAll(ctx *gin.Context) {
	medicines, err := c.repo.GetAll(ctx.Request.Context())
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, medicines)
}

func (c *MedicineController) Create(ctx *gin.Context) {
	var req struct {
		Name         string  `json:"name" binding:"required"`
		GenericName  string  `json:"generic_name"`
		Category     string  `json:"category"`
		Unit         string  `json:"unit"`
		Price        float64 `json:"price" binding:"required"`
		ReorderLevel int     `json:"reorder_level"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	medicine, err := c.repo.Create(ctx.Request.Context(), req.Name, req.GenericName, req.Category, req.Unit, req.Price, req.ReorderLevel)
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Created(ctx, medicine)
}
