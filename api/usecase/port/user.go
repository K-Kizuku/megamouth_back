package port

import (
	"context"
	"megamouth/api/usecase/schema"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*schema.UserOutput)
	RenderError(error)
}
