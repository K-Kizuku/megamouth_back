package repository

import (
	"megamouth/api/entity/models"

	"github.com/gin-gonic/gin"
)

type PostRepository interface {
	CreatePost(ctx *gin.Context) (*models.Post, error)
	GetPosts(ctx *gin.Context) (*models.Post, error)
	GetPostByID(ctx *gin.Context) (*models.Post, error)
	UpdatePostByID(ctx *gin.Context) (*models.Post, error)
	DeletepostByID(ctx *gin.Context) (*models.Post, error)
}
