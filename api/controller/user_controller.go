package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nevzatcirak/go-gin-poc/domain"
	"github.com/nevzatcirak/go-gin-poc/service"
	validators "github.com/nevzatcirak/go-gin-poc/validator"
	"net/http"
	"strconv"
)

type UserController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type userController struct {
	service service.UserService
}

var userValidate *validator.Validate

func NewUserController(service service.UserService) UserController {
	userValidate = validator.New()
	userValidate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &userController{
		service: service,
	}
}

func (c *userController) FindAll(ctx *gin.Context) {
	all := c.service.FindAll()
	ctx.JSON(http.StatusOK, all)
}

func (c *userController) Find(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	foundUser := c.service.Find(id)
	ctx.JSON(http.StatusOK, foundUser)
}

func (c *userController) Save(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = userValidate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = c.service.Save(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

func (c *userController) Update(ctx *gin.Context) {
	var user domain.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user.ID = id
	err = userValidate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated!"})
}

func (c *userController) Delete(ctx *gin.Context) {
	var user domain.User
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user.ID = id
	err = c.service.Delete(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}
