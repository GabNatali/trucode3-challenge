package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
	"github.com/GabNatali/trucode3-challenge-final/internal/user"
)

type AuthService interface {
	Login(dto LoginUserDto) (LoggedUserDto, error)
	VerifyAccessToken(accessToken string) (uint, error)
	ParseAccessToken(accessToken string) (uint, error)
}

type AuthServiceOpts struct {
	UserRepository user.UserRepository
	Config         string
	Crypto         crypto.Crypto
}

func NewAuthService(opts AuthServiceOpts) AuthService {
	return &authService{
		UserRepository: opts.UserRepository,
		Config:         opts.Config,
		Crypto:         opts.Crypto,
	}
}

type authService struct {
	user.UserRepository
	Config string
	crypto.Crypto
}

func (s *authService) Login(in LoginUserDto) (out LoggedUserDto, err error) {

	user, err := s.UserRepository.GetByEmail(in.Email)
	if err != nil {
		return out, errors.New("invalid Credentials")
	}

	if !user.ComparePassword(in.Password, s.Crypto) {
		return out, errors.New("invalid Credentials")
	}

	token, err := s.generateAccessToken(user)
	if err != nil {
		return out, err
	}

	return out.MapFromModel(user, token), nil

}

func (s *authService) VerifyAccessToken(accessToken string) (uint, error) {
	payload, err := s.ParseAndValidateJWT(accessToken, s.Config)

	if err != nil {
		return 0, errors.New("invalid token")
	}

	value, ok := payload["userId"].(float64)
	if !ok {
		return 0, errors.New("unauthorized")
	}
	userId := uint(value)
	return userId, nil
}

func (s *authService) ParseAccessToken(accessToken string) (uint, error) {
	payload, err := s.ParseJWT(accessToken, s.Config)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	userId, ok := payload["userId"].(uint)
	if !ok {
		return 0, errors.New("unauthorized")
	}

	return userId, nil
}

func (s *authService) generateAccessToken(user user.User) (string, error) {
	payload := map[string]interface{}{
		"userId": user.Id,
		"email":  strings.TrimSpace(user.Email),
		"name":   strings.TrimSpace(user.Username),
	}
	exp := time.Now().Add(time.Hour * 1)
	return s.GenerateJWT(
		payload,
		s.Config,
		exp,
	)
}
