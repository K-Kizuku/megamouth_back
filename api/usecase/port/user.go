package port

import (
	"megamouth/api/entity/models"
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type UserInputPort interface {
	GetUserByID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	IsUsedName(ctx *gin.Context)
}

type UserOutputPort interface {
	Render(*schema.UserOutput)
	RenderJWT(*models.User)
	RenderJWTwithUser(*models.User)
	RenderError(error)
	RenderIsUsed(bool)
}
