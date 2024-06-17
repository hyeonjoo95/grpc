package service

import (
	"errors"
	"main/api/proto/server"
	"regexp"
	"strings"
)

// ValidateCreateUserRequest: 유저 생성 요청에 대한 유효성 확인
func ValidateCreateUserRequest(req *server.CreateUserRequest) (err error) {
	if req.Email == "" {
		err = errors.New("email is required")
		return
	}

	err = validateEmail(req.Email)
	if err != nil {
		return
	}

	if req.Password == "" {
		err = errors.New("password is required")
		return
	}

	if req.Name == "" {
		err = errors.New("name is required")
		return
	}

	if req.Age <= 0 {
		err = errors.New("age must be greater than zero")
		return
	}

	if req.PhoneNumber == "" {
		err = errors.New("phone_number is required")
		return
	}

	req.PhoneNumber = strings.Replace(req.PhoneNumber, "-", "", -1) // 핸드폰번호의 경우 -를 제외한 숫자만 입력되도록 처리
	err = validatePhoneNumber(req.PhoneNumber)
	if err != nil {
		return
	}

	return
}

// ValidateUpdateUserRequest: 유저 정보 변경 요청에 대한 유효성 확인
func ValidateUpdateUserRequest(req *server.UpdateUserRequest) (err error) {
	if req.Id == 0 {
		err = errors.New("id is required")
		return
	}

	if req.Age != nil && *req.Age <= 0 {
		err = errors.New("age must be greater than zero")
		return
	}

	if req.PhoneNumber != nil {
		if *req.PhoneNumber == "" {
			err = errors.New("phone_number is required")
			return
		} else {
			*req.PhoneNumber = strings.Replace(*req.PhoneNumber, "-", "", -1)
			err = validatePhoneNumber(*req.PhoneNumber)
			if err != nil {
				return
			}
		}
	}

	return
}

// ValidateDeleteUserRequest: 유저 탈퇴 요청에 대한 유효성 확인
func ValidateDeleteUserRequest(req *server.DeleteUserRequest) (err error) {
	if req.Id == 0 {
		err = errors.New("id is required")
		return
	}

	return
}

// ValidateGetUserRequest: 유저 정보 조회 요청에 대한 유효성 확인
func ValidateGetUserRequest(req *server.GetUserRequest) (err error) {
	if req.Id == 0 {
		err = errors.New("id is required")
		return
	}

	return
}

// ValidateLoginRequest: 로그인 요청에 대한 유효성 확인
func ValidateLoginRequest(req *server.LoginRequest) (err error) {
	if req.Email == "" {
		err = errors.New("email is required")
		return
	}

	if req.Password == "" {
		err = errors.New("password is required")
		return
	}

	return
}

// validateEmail: 이메일 유효성 체크
func validateEmail(email string) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if !regex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

// validatePhoneNumber: 핸드폰번호 유효성 체크
func validatePhoneNumber(phoneNumber string) error {
	regex := regexp.MustCompile("^[0-9]+$")

	if !regex.MatchString(phoneNumber) {
		return errors.New("invalid phone number format")
	}

	return nil
}
