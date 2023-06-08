package services

import (
	"casbin-go_gin/internal/models"
	"casbin-go_gin/internal/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SignKey = "qiweurypuyub3894khlo7ho78f0fhh0hhkh"

type AuthSRV struct {
	repo *repository.Auth
}

type TokenClaims struct {
	jwt.StandardClaims
	UserRole string `json:"user_role"`
	UserId   int
}

func NewAuthSRV(repo *repository.Auth) *AuthSRV {
	return &AuthSRV{repo: repo}
}

func (s *AuthSRV) CreateUser(user models.User) error {
	user.Password = HashPassword(user.Password)
	return s.repo.Crete(user)
}

func HashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte("Salt")))
}

func (s *AuthSRV) GenerateToken(name string, password string) (string, error) {
	user, err := s.repo.Get(name, HashPassword(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserRole: user.Role,
		UserId:   user.Id,
	})

	tokenString, err := token.SignedString([]byte(SignKey))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func (s *AuthSRV) Parse(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(SignKey), nil
		})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("error: unauthorized")
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}
	role := claims.UserRole
	return role, nil
}
