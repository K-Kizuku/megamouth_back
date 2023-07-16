package port

import (
	"megamouth/api/entity/models"
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type PostInputPort interface {
	CreatePost(ctx *gin.Context)
	GetPosts(ctx *gin.Context)
	GetPostByID(ctx *gin.Context)
	UpdatePostByID(ctx *gin.Context)
	DeletepostByID(ctx *gin.Context)
}

type PostOutputPort interface {
	Render(*schema.PostOutput)
	RenderAll([]models.Post)

	RenderError(error)
}
