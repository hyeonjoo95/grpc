package repository

import (
	"context"
	"errors"
	"main/api/proto/server"
	"main/common"
	"main/database"
	"main/ent"
	"main/ent/user"
)

func CreateUser(ctx context.Context, req *server.CreateUserRequest) (userInfo *ent.User, err error) {
	// email 중복 체크
	isExist, err := database.Client.User.Query().
		Where(
			user.EmailEQ(req.Email),
			user.IsUsedEQ(true),
		).Exist(ctx)

	if err != nil {
		return
	}

	if isExist {
		err = errors.New("email already exists")
		return
	}

	hashPassword, err := common.HashPassword(req.Password)
	if err != nil {
		return
	}

	userInfo, err = database.Client.User.Create().
		SetEmail(req.Email).
		SetPassword(hashPassword).
		SetName(req.Name).
		SetAge(req.Age).
		SetPhoneNumber(req.PhoneNumber).
		Save(ctx)

	if err != nil {
		return
	}

	return
}

func UpdateUser(ctx context.Context, req *server.UpdateUserRequest) (userInfo *ent.User, err error) {

	query := database.Client.User.Update().
		Where(
			user.IDEQ(int(req.Id)),
			user.IsUsedEQ(true),
		)

	if req.Age != nil {
		query.SetAge(*req.Age)
	}

	if req.Name != nil {
		query.SetName(*req.Name)
	}

	if req.PhoneNumber != nil {
		query.SetPhoneNumber(*req.PhoneNumber)
	}

	err = query.Exec(ctx)
	if err != nil {
		return
	}

	userInfo, err = database.Client.User.Query().
		Where(
			user.IDEQ(int(req.Id)),
		).Only(ctx)

	if err != nil {
		return
	}

	return
}

func DeleteUser(ctx context.Context, req *server.DeleteUserRequest) (err error) {
	phoneNumber, err := database.Client.User.Query().
		Where(
			user.IDEQ(int(req.Id)),
			user.IsUsedEQ(true),
		).Select(user.FieldPhoneNumber).
		String(ctx)

	if err != nil {
		return
	}

	maskPhoneNumber := common.MaskPhoneNumber(phoneNumber)

	err = database.Client.User.Update().
		Where(
			user.IDEQ(int(req.Id)),
			user.IsUsedEQ(true),
		).
		SetIsUsed(false).
		SetPhoneNumber(maskPhoneNumber).
		Exec(ctx)

	if err != nil {
		return
	}

	return
}

func GetUser(ctx context.Context, req *server.GetUserRequest) (userInfo *ent.User, err error) {
	userInfo, err = database.Client.User.Query().
		Where(
			user.IDEQ(int(req.Id)),
			user.IsUsedEQ(true),
		).Only(ctx)

	if err != nil {
		return
	}

	return
}

func GetAllUser(ctx context.Context, isActived bool) (userInfos []*ent.User, err error) {
	userInfos, err = database.Client.User.Query().
		Where(
			user.IsUsedEQ(isActived),
		).All(ctx)

	if err != nil {
		return
	}

	return
}

func LoginCheck(ctx context.Context, req *server.LoginRequest) (err error) {
	userInfo, err := database.Client.User.Query().
		Where(
			user.EmailEQ(req.Email),
			user.IsUsedEQ(true),
		).Only(ctx)

	if err != nil {
		return
	}

	err = common.VerifyPassword(userInfo.Password, req.Password)
	if err != nil {
		return
	}

	return
}
