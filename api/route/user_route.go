package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nevzatcirak/go-gin-poc/api/controller"
	"github.com/nevzatcirak/go-gin-poc/repository"
	"github.com/nevzatcirak/go-gin-poc/service"
	"gorm.io/gorm"
)

func NewUserRouter(db *gorm.DB, group *gin.RouterGroup) {
	tr := repository.NewUserRepository(db)
	userController := controller.NewUserController(service.NewUserService(tr))

	group.GET("/users", userController.FindAll)
	group.POST("/users", userController.Save)
	group.PUT("/users/:id", userController.Update)
	group.DELETE("/users/:id", userController.Delete)
}
