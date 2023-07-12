package interactor

import (
	"context"

	"megamouth/api/entity/repository"
	"megamouth/api/usecase/port"
	"megamouth/api/usecase/schema"
)

type User struct {
	OutputPort port.UserOutputPort
	UserRepo   repository.UserRepository
}

func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
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
