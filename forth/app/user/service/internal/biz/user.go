package biz

import (
	"context"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Get(ctx context.Context, id int64) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}
