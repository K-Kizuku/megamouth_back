package repository

import (
	"megamouth/api/entity/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	GetUserByID(ctx *gin.Context) (*models.User, error)
	CreateUser(ctx *gin.Context) (*models.User, error)
	LoginUser(ctx *gin.Context) (*models.User, error)
	IsUsedName(ctx *gin.Context) (*models.User, error)
}
