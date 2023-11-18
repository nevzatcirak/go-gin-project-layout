package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nevzatcirak/go-gin-poc/bootstrap"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	//Serving Static Contents
	gin.Static("/css", "./template/css")
	gin.LoadHTMLGlob("./template/*.html")

	publicRouter := gin.Group("/api")
	// All Public APIs
	NewVideoRouter(db, publicRouter)
	NewUserRouter(db, publicRouter)

	uiRouter := gin.Group("/view")
	//User Interface
	NewUIRouter(db, uiRouter)
}
