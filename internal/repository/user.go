package repository

import (
	"github.com/jmoiron/sqlx"
	"someproject/internal/models"
)

type UsersRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepositoryPostgres {
	return &UsersRepositoryPostgres{db: db}
}

func (u *UsersRepositoryPostgres) GetUsers() (*[]models.User, error) {

	var users *[]models.User
	err := u.db.Select(users, `SELECT users.name, likes.video_id, comments.video_id, comments.text FROM users JOIN likes ON (users.id = likes.user_id) JOIN comments ON (users.id = comments.user_id)`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
