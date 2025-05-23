package crypto

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Crypto interface {
	HashPassword(password string) (string, error)
	CompareHashAndPassword(hash string, password string) bool

	GenerateJWT(payload map[string]interface{}, secret string, exp time.Time) (string, error)
	ParseAndValidateJWT(token string, secret string) (map[string]interface{}, error)
	ParseJWT(token string, secret string) (map[string]interface{}, error)

	GenerateUUID() (string, error)
}

func NewCrypto() Crypto {
	return &cryptoImpl{}
}

type cryptoImpl struct{}

func (*cryptoImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (*cryptoImpl) CompareHashAndPassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (*cryptoImpl) GenerateJWT(payload map[string]interface{}, secret string, exp time.Time) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = exp.Unix()

	for key, value := range payload {
		claims[key] = value
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func (*cryptoImpl) ParseAndValidateJWT(token string, secret string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return map[string]interface{}{}, err
	}
	if !parsedToken.Valid {
		return map[string]interface{}{}, errors.New("token validation error")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]interface{}{}, errors.New("token validation error")
	}

	payload := make(map[string]interface{})

	for key, value := range claims {
		payload[key] = value
	}

	return payload, nil
}

func (*cryptoImpl) ParseJWT(token string, secret string) (map[string]interface{}, error) {
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return map[string]interface{}{}, errors.New("token parsing error")
	}

	payload := make(map[string]interface{})

	for key, value := range claims {
		payload[key] = value
	}

	return payload, nil
}

func (*cryptoImpl) GenerateUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
