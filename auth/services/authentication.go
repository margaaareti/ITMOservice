package services

import (
	"context"
	"itmo/auth"
	"itmo/models"
)

type AuthService struct {
	userStorage auth.AuthStorage
}

func CreateNewAuthService(userStorage auth.AuthStorage) *AuthService {
	return &AuthService{userStorage: userStorage}
}

func (a *AuthService) SignUp(ctx context.Context, user models.User) (uint64, error) {

	user.Password = auth.GeneratePasswordHash(user.Password)

	return a.userStorage.CreateUser(ctx, user)
}
