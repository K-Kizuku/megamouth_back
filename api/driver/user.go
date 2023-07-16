package driver

/*
driver パッケージは，技術的な実装を持ちます．．
*/

import (
	"megamouth/api/adapter/controller"
	"megamouth/api/adapter/gateway"
	"megamouth/api/adapter/presenter"
	"megamouth/api/usecase/interactor"
	"megamouth/api/utils/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Serve はserverを起動させます．
func InitUserRouter(ur *gin.RouterGroup, conn *gorm.DB) {
	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}
	ur.GET("/:id", middleware.AuthMiddleware, user.GetUserByID)
	ur.POST("/create", user.CreateUser)
	ur.POST("/login", user.LoginUser)
	ur.GET("/is_used/:id", user.IsUsedName)
}
