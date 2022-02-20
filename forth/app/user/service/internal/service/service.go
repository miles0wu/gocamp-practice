package service

import (
	"github.com/google/wire"
	v1 "github.com/miles0wu/gocamp-prctice/api/user/service/v1"
	"github.com/miles0wu/gocamp-prctice/app/user/service/internal/biz"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	v1.UnimplementedUserServiceServer

	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}
