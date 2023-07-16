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
func InitPostRouter(pr *gin.RouterGroup, conn *gorm.DB) {
	post := controller.Post{
		OutputFactory: presenter.NewPostOutputPort,
		InputFactory:  interactor.NewPostInputPort,
		RepoFactory:   gateway.NewPostRepository,
		Conn:          conn,
	}
	pr.GET("/:id", middleware.AuthMiddleware, post.GetPostByID)
	pr.GET("/all", middleware.AuthMiddleware, post.GetPosts)
	pr.POST("/create", middleware.AuthMiddleware, post.CreatePost)
	pr.PUT("/:id", middleware.AuthMiddleware, post.UpdatePostByID)
	pr.DELETE("/:id", middleware.AuthMiddleware, post.DeletepostByID)
}
