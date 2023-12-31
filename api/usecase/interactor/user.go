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
	res := schema.UserOutput{ID: user.ID, Name: user.Name}
	u.OutputPort.Render(&res)
}

func (u *User) CreateUser(ctx *gin.Context) {
	user, err := u.UserRepo.CreateUser(ctx)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.RenderJWT(user)
}

func (u *User) LoginUser(ctx *gin.Context) {
	user, img, err := u.UserRepo.LoginUser(ctx)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.RenderJWTwithUser(user, img)
}

func (u *User) IsUsedName(ctx *gin.Context) {
	user, err := u.UserRepo.GetUserByID(ctx)
	if err == nil && user == nil {
		u.OutputPort.RenderIsUsed(false)
		return
	}
	u.OutputPort.RenderIsUsed(true)

}

func NewUserInputPort(outputPort port.UserOutputPort, userRepository repository.UserRepository) port.UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}
