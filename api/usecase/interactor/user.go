package interactor

import (
	"megamouth/api/entity/repository"
	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"

	"github.com/gin-gonic/gin"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   repository.UserRepository
}

func (u *User) GetUserByID(ctx *gin.Context) {
	user, err := u.UserRepo.GetUserByID(ctx)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	res := schema.UserOutput{ID: user.ID, Name: user.Name, Email: user.Email}
	u.OutputPort.Render(&res)
}

func NewUserInputPort(outputPort port.UserOutputPort, userRepository repository.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}
