package common

import (
	"main/constant"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// HashPassword: 비밀번호를 bcrypt를 사용하여 해시화
func HashPassword(password string) (result string, err error) {
	// 비밀번호를 bcrypt 해시로 변환합니다.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	result = string(hashedPassword)
	return
}

// VerifyPassword: 해시된 비밀번호와 평문 비밀번호를 비교하여 일치 여부를 확인
func VerifyPassword(hashedPassword, password string) (err error) {
	// 해시된 비밀번호와 평문 비밀번호를 비교합니다.
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return
	}

	return
}

// MaskPhoneNumber: 핸드폰 번호를 마스킹
func MaskPhoneNumber(phoneNumber string) (masked string) {
	masked = phoneNumber[:4] + strings.Repeat("*", 2) + phoneNumber[6:8] + strings.Repeat("*", 2) + phoneNumber[10:]
	return
}

// CreateUserToken: JWT를 사용하여 유저 토큰을 생성
func CreateUserToken(email string) (signedToken string, err error) {
	// JWT에 포함될 클레임 설정
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// JWT 생성
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// JWT 서명 키
	signingKey := []byte(constant.JwtSecret)

	// JWT 서명
	signedToken, err = token.SignedString(signingKey)
	if err != nil {
		return
	}

	return
}
