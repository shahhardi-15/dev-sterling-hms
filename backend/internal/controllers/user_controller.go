package controllers

import (
	"log"
	"sterling-hms/internal/services"
	"sterling-hms/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.service.GetAllUsers(ctx.Request.Context())
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, users)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.service.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}
	utils.Success(ctx, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var req services.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	user, err := c.service.CreateUser(ctx.Request.Context(), req)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Created(ctx, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var req services.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}

	user, err := c.service.UpdateUser(ctx.Request.Context(), id, req)
	if err != nil {
		utils.BadRequest(ctx, err.Error())
		return
	}
	utils.Success(ctx, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.service.DeleteUser(ctx.Request.Context(), id)
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, gin.H{"message": "user deleted successfully"})
}

func (c *UserController) GetAllPatients(ctx *gin.Context) {
	patients, err := c.service.GetAllPatients(ctx.Request.Context())
	if err != nil {
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, patients)
}

func (c *UserController) GetAllDoctors(ctx *gin.Context) {
	doctors, err := c.service.GetAllDoctors(ctx.Request.Context())
	if err != nil {
		log.Printf("GetAllDoctors error: %v", err)
		utils.InternalError(ctx, err.Error())
		return
	}
	utils.Success(ctx, doctors)
}

func (c *UserController) GetDoctorByUserID(ctx *gin.Context) {
	userID := ctx.Param("id")
	doctor, err := c.service.GetDoctorByUserID(ctx.Request.Context(), userID)
	if err != nil {
		utils.NotFound(ctx, err.Error())
		return
	}
	utils.Success(ctx, doctor)
}
