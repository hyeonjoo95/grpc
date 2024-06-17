package service

import (
	"context"
	"main/api/proto/server"
	"main/common"
	"main/database/repository"
)

type UserService struct {
	server.UserServiceServer
}

// CreateUser: 유저 생성
func (s *UserService) CreateUser(
	ctx context.Context,
	req *server.CreateUserRequest,
) (res *server.CreateUserResponse, err error) {

	err = ValidateCreateUserRequest(req)
	if err != nil {
		return
	}

	userInfo, err := repository.CreateUser(ctx, req)
	if err != nil {
		return
	}

	res = &server.CreateUserResponse{
		User: &server.User{
			Id:          uint32(userInfo.ID),
			Email:       userInfo.Email,
			Name:        userInfo.Name,
			Age:         userInfo.Age,
			PhoneNumber: userInfo.PhoneNumber,
		},
	}

	return
}

// GetUser: 유저 정보 조회
func (s *UserService) GetUser(
	ctx context.Context,
	req *server.GetUserRequest,
) (res *server.GetUserResponse, err error) {
	err = ValidateGetUserRequest(req)
	if err != nil {
		return
	}

	userInfo, err := repository.GetUser(ctx, req)
	if err != nil {
		return
	}

	res = &server.GetUserResponse{
		User: &server.User{
			Id:          uint32(userInfo.ID),
			Email:       userInfo.Email,
			Name:        userInfo.Name,
			Age:         userInfo.Age,
			PhoneNumber: userInfo.PhoneNumber,
		},
	}
	return
}

// GetAllActiveUser: 가입한 유저 목록 조회
func (s *UserService) GetAllActiveUser(
	ctx context.Context,
	req *server.DefaultRequest,
) (res *server.GetAllUserResponse, err error) {

	userInfos, err := repository.GetAllUser(ctx, true)
	if err != nil {
		return
	}

	var users []*server.User
	for _, u := range userInfos {
		pbUser := &server.User{
			Id:          uint32(u.ID),
			Email:       u.Email,
			Name:        u.Name,
			Age:         u.Age,
			PhoneNumber: u.PhoneNumber,
		}
		users = append(users, pbUser)
	}

	res = &server.GetAllUserResponse{
		User: users,
	}

	return
}

// GetAllInActiveUser: 탈퇴한 유저 목로 조회
func (s *UserService) GetAllInActiveUser(
	ctx context.Context,
	req *server.DefaultRequest,
) (res *server.GetAllUserResponse, err error) {

	userInfos, err := repository.GetAllUser(ctx, false)
	if err != nil {
		return
	}

	var users []*server.User
	for _, u := range userInfos {
		pbUser := &server.User{
			Id:          uint32(u.ID),
			Email:       u.Email,
			Name:        u.Name,
			Age:         u.Age,
			PhoneNumber: u.PhoneNumber,
		}
		users = append(users, pbUser)
	}

	res = &server.GetAllUserResponse{
		User: users,
	}

	return
}

// UpdateUser: 유저 정보 변경
func (s *UserService) UpdateUser(
	ctx context.Context,
	req *server.UpdateUserRequest,
) (res *server.UpdateUserResponse, err error) {
	err = ValidateUpdateUserRequest(req)
	if err != nil {
		return
	}

	userInfo, err := repository.UpdateUser(ctx, req)
	if err != nil {
		return
	}

	res = &server.UpdateUserResponse{
		User: &server.User{
			Id:          uint32(userInfo.ID),
			Email:       userInfo.Email,
			Name:        userInfo.Name,
			Age:         userInfo.Age,
			PhoneNumber: userInfo.PhoneNumber,
		},
	}

	return
}

// DeleteUser: 유저 탈퇴
func (s *UserService) DeleteUser(
	ctx context.Context,
	req *server.DeleteUserRequest,
) (res *server.DeleteUserResponse, err error) {
	err = ValidateDeleteUserRequest(req)
	if err != nil {
		return
	}

	err = repository.DeleteUser(ctx, req)
	if err != nil {
		return
	}

	res = &server.DeleteUserResponse{
		Success: true,
	}

	return
}

// Login: 로그인
func (s *UserService) Login(ctx context.Context, req *server.LoginRequest) (res *server.LoginResponse, err error) {
	err = ValidateLoginRequest(req)
	if err != nil {
		return
	}

	err = repository.LoginCheck(ctx, req)
	if err != nil {
		return
	}

	signedToken, err := common.CreateUserToken(req.Email)
	if err != nil {
		return
	}

	res = &server.LoginResponse{
		Token: signedToken,
	}

	return
}
