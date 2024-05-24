package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	securefilechanger "github.com/KodokuOdius/SecureFileChanger"
	"github.com/KodokuOdius/SecureFileChanger/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

// основные константы
const (
	salt       = "qp234895yw450otuhwsreolgs"
	tokenTTL   = 12 * time.Hour
	signingKey = "q78o423rytq4378rtywo487tghwoi43uythgw3i4oty"
)

// Структура для токена
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// Сервис авторизации
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Создание пользователя
func (s *AuthService) CreateUser(user securefilechanger.User) (int, error) {
	user.Password = hashPassword(user.Password)

	adminExist, err := s.repo.CheckAdminIsExists()
	if err != nil {
		return 0, nil
	}

	if !adminExist {
		user.IsAdmin = true
		user.IsApproved = true
	}

	userId, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	path := filepath.Join(".", fmt.Sprintf("files/user%d/bin", userId))
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return 0, err
	}

	return userId, err
}

func (s *AuthService) CheckAdminIsExists() (bool, error) {
	adminExist, err := s.repo.CheckAdminIsExists()
	if err != nil {
		return adminExist, err
	}

	return adminExist, nil
}

// Генерация токена
func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, hashPassword(password))
	if err != nil {
		return "", err
	}

	if !user.IsApproved {
		return "", errors.New("user not approved")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// Обработка токена
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invelid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token invelid")
	}

	return claims.UserId, nil
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
