package psql

import (
	"context"
	"database/sql"
	"learn-api/internal/db"
	"learn-api/internal/entity"
	"learn-api/internal/repository"
	"strconv"
)

type userRepo struct {
	master *sql.DB
	slave  *sql.DB
}

func NewUserRepoPG(master, slave *sql.DB) repository.UserRepository {
	return &userRepo{
		master: master,
		slave:  slave,
	}
}

// func get all user
func (r *userRepo) All() ([]*entity.User, error) {
	q := db.New(r.slave)
	users, err := q.GetUsers(context.Background())
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	for _, u := range users {
		result = append(result, &entity.User{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return result, nil
}

// GetByID
func (r *userRepo) GetByID(id string) (*entity.User, error) {
	q := db.New(r.slave)
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	user, err := q.GetUserByID(context.Background(), int32(intID))
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// GetByEmail
func (r *userRepo) GetByEmail(email string) (*entity.User, error) {
	q := db.New(r.slave)
	user, err := q.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Create
func (r *userRepo) Create(user *entity.User) error {
	q := db.New(r.master)
	params := db.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	return q.CreateUser(context.Background(), params)
}
