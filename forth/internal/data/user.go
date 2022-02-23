package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"user/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type user struct {
	Id       int64 `gorm:"primaryKey"`
	Username string
	Password string
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	var poUser user
	result := r.data.db.First(&poUser, id)
	if result.RowsAffected != 1 {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	// convert persistent object to domain object
	doUser := biz.User{
		Id:       poUser.Id,
		Username: poUser.Username,
		Password: poUser.Password,
	}

	return &doUser, nil
}
