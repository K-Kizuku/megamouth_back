package port

import (
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type UserInputPort interface {
	GetUserByID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
}

type UserOutputPort interface {
	Render(*schema.UserOutput)
	RenderError(error)
}
