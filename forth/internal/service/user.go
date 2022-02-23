package service

import (
	"context"
	v1 "user/api/user"
)

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserInfo) (*v1.UserInfoResponse, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &v1.UserInfoResponse{
		Id:       rv.Id,
		Username: rv.Username,
		Password: rv.Password,
	}, nil
}
