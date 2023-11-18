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

type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	service     service.VideoService
	userService service.UserService
}

var videoValidate *validator.Validate

func NewVideoController(service service.VideoService, userService service.UserService) VideoController {
	videoValidate = validator.New()
	videoValidate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service:     service,
		userService: userService,
	}
}

func (c *videoController) FindAll(ctx *gin.Context) {
	all := c.service.FindAll()
	ctx.JSON(http.StatusOK, all)
}

func (c *videoController) Save(ctx *gin.Context) {
	var video domain.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if video.Author.ID != 0 {
		foundUser := c.userService.Find(video.Author.ID)
		video.UserID = foundUser.ID
	}
	err = videoValidate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = c.service.Save(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Success!"})
}

func (c *videoController) Update(ctx *gin.Context) {
	var video domain.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	video.ID = id
	if video.Author.ID != 0 {
		foundUser := c.userService.Find(video.Author.ID)
		video.UserID = foundUser.ID
	}
	err = videoValidate.Struct(video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = c.service.Update(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated!"})
}

func (c *videoController) Delete(ctx *gin.Context) {
	var video domain.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	video.ID = id
	err = c.service.Delete(video)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted!"})
}

func (c *videoController) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
