package auth

import (
	"context"
	"itmo/models"
)

const CtxUserId = "user_id"
const CtxUserName = "user_name"
const CtxUserSurname = "user_surname"
const CtxUserPatronymic = "userID"

type AuthService interface {
	SignUp(ctx context.Context, user models.User) (uint64, error)
}
