package service

import (
	"ai-feed/internal/entity"
	"ai-feed/internal/storage"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"time"
)

type AuthConfig struct {
	SecretKey    string `yaml:"secret_key"`
	PasswordSalt string `yaml:"password_salt"`
}

func (service *AiFeed) AuthUser(credentials *entity.User) (string, error) {
	password := credentials.Password + service.authCfg.PasswordSalt

	h := sha1.New()
	h.Write([]byte(password))
	sha := h.Sum(nil)

	credentials.Password = hex.EncodeToString(sha)

	err := service.users.Check(credentials)
	if err == nil {
		return service.returnAuthToken(credentials.Login)
	}

	err = service.users.Create(credentials)

	if errors.Is(err, storage.ErrAlreadyExists) {
		return "", ErrLoginAlreadyExists
	}

	if err != nil {
		log.Err(err).Timestamp().Msg("failed auth user")
		return "", err
	}

	// add init personalities for new users.
	for _, personality := range entity.InitPersonalities {
		if err := service.CreatePersonality(
			context.WithValue(context.Background(), storage.UserLogin, credentials.Login), &personality); err != nil {
			return "", err
		}
	}

	return service.returnAuthToken(credentials.Login)
}

// VerifyToken returns user login from JWT token
func (service *AiFeed) VerifyToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.authCfg.SecretKey), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", storage.ErrUnauthorized
	}

	userLoginRaw, ok := claims["login"]
	if !ok || userLoginRaw == nil {
		return "", storage.ErrUnauthorized
	}

	userLogin, ok := userLoginRaw.(string)
	if !ok || userLogin == "" {
		return "", storage.ErrUnauthorized
	}

	return userLogin, nil
}

func (service *AiFeed) returnAuthToken(login string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days token expiration
		"iat":   time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(service.authCfg.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
