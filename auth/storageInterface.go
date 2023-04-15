package auth

import (
	"context"
	"itmo/models"
)

type AuthStorage interface {
	CreateUser(ctx context.Context, user models.User) (uint64, error)
}
