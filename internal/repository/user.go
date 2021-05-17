package repository

import (
	"fmt"
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

	users := make([]models.User, 0)
	if err := u.db.Select(&users, "SELECT * FROM users"); err != nil {
		return nil, err
	}

	for i, user := range users {

		comments := make([]models.Comment, 0)
		likes := make([]models.Like, 0)

		if err := u.db.Select(&comments, "SELECT (id, video_id, text) FROM comments WHERE user_id = ($1)", user.ID); err != nil {
			return nil, err
		}

		if err := u.db.Select(&likes, "SELECT (id, video_id) FROM likes"); err != nil {
			return nil, err
		}

		users[i].Comments = &comments
		users[i].Likes = &likes
	}

	return &users, nil
}

func (u *UsersRepositoryPostgres) AddUser(user *models.User) error {
	if _, err := u.db.Exec("INSERT INTO users (name) VALUES ($1)", user.Name); err != nil {
		return err
	}
	return nil
}

func (u *UsersRepositoryPostgres) AddComments(user *models.User) error {
	if user == nil {
		return fmt.Errorf("bad request")
	}
	if user.Comments == nil {
		return fmt.Errorf("bad request")
	}

	for _, comment := range *user.Comments {
		if _, err := u.db.Exec("INSERT INTO comments (video_id, text, user_id) VALUES ($1, $2, $3)", comment.VideoID, comment.Text, user.ID); err != nil {
			return err
		}
	}
	return nil
}

func (u *UsersRepositoryPostgres) AddLikes(user *models.User) error {
	if user == nil {
		return fmt.Errorf("bad request")
	}
	if user.Likes == nil {
		return fmt.Errorf("bad request")
	}

	for _, like := range *user.Likes {
		if _, err := u.db.Exec("INSERT INTO likes (video_id, user_id) VALUES ($1, $2)", like.VideoID, user.ID); err != nil {
			return err
		}
	}
	return nil
}
