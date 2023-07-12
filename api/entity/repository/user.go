package repository

import (
	"context"
	"megamouth/api/entity/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, userID string) (*models.User, error)
}
