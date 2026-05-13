package controllers

import (
	"sterling-hms/internal/services"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req services.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	response, err := c.service.Login(ctx.Request.Context(), req)
	if err != nil {
		utils.Unauthorized(ctx, err.Error())
		return
	}

	utils.Success(ctx, response)
}

func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=8"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	claims := ctx.MustGet("claims").(*utils.JWTClaims)

	err := c.service.ChangePassword(ctx.Request.Context(), claims.UserID, req.OldPassword, req.NewPassword)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	utils.Success(ctx, gin.H{"message": "password changed successfully"})
}
