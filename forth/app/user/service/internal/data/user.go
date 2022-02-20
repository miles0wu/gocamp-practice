package data

import (
	"context"
	"github.com/miles0wu/gocamp-prctice/app/user/service/internal/biz"
)

type userRepo struct {
}

func NewUserRepo() biz.UserRepo {
	return &userRepo{}
}

func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	return &biz.User{Id: id, Username: "miles"}, nil
}
