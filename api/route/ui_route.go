package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nevzatcirak/go-gin-poc/api/controller"
	"github.com/nevzatcirak/go-gin-poc/repository"
	"github.com/nevzatcirak/go-gin-poc/service"
	"gorm.io/gorm"
)

func NewUIRouter(db *gorm.DB, group *gin.RouterGroup) {
	vr := repository.NewVideoRepository(db)
	ur := repository.NewUserRepository(db)
	videoController := controller.NewVideoController(service.NewVideoService(vr), service.NewUserService(ur))

	group.GET("/videos", videoController.ShowAll)
}
