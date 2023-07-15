package driver

/*
driver パッケージは，技術的な実装を持ちます．．
*/

import (
	"megamouth/api/adapter/controller"
	"megamouth/api/adapter/gateway"
	"megamouth/api/adapter/presenter"
	"megamouth/api/usecase/interactor"

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
	pr.GET("/:id", post.GetPostByID)
	pr.GET("/all", post.GetPosts)
	pr.POST("/create", post.CreatePost)
	pr.PUT("/:id", post.UpdatePostByID)
	pr.DELETE("/:id", post.DeletepostByID)
}
