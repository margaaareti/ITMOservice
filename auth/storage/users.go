package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"itmo/models"
	"itmo/server/repository"
)

type UserStorage struct {
	pgsqlDb *pgxpool.Pool
}

func CreateNewUserStorage(pgsqlDb *pgxpool.Pool) *UserStorage {
	return &UserStorage{pgsqlDb: pgsqlDb}
}

func (s *UserStorage) CreateUser(ctx context.Context, user models.User) (uint64, error) {

	var (
		id         uint64
		mailStatus int
		userStatus int
	)

	isExistCheck := fmt.Sprintf(`SELECT COUNT(id), (SELECT COUNT(id) FROM %[1]s WHERE username = $2) FROM %[1]s WHERE email = $1 `, repository.UserTable)
	isExistRow := s.pgsqlDb.QueryRow(ctx, isExistCheck, user.Email, user.Username)
	if err := isExistRow.Scan(&mailStatus, &userStatus); err != nil {
		return 0, err
	}

	if mailStatus != 0 {
		return 0, errors.New(userAlrExist)
	} else if mailStatus != 0 {
		return 0, errors.New(emailAlrExist)
	} else {
		query := fmt.Sprintf(`INSERT INTO %s (name,surname,patronymic,username,password,email) values ($1,$2,$3,$4,$5,$6) RETURNING id`, repository.UserTable)
		row := s.pgsqlDb.QueryRow(ctx, query, user.Name, user.Surname, user.Patronymic, user.Username, user.Password, user.Email)
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil

}
