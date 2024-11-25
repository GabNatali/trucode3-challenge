package user

import (
	"strings"

	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
)

type UserService interface {
	Add(dto AddUserDto) (UserDto, error)
	GetMe(token string) (User, error)
	UpdateConfig(newconfig UpdateBody) (ConfigUser, error)
}

func NewUserService(repo UserRepository, crypto crypto.Crypto, config string) UserService {
	return &userService{
		repo:   repo,
		Crypto: crypto,
		Config: config,
	}
}

type userService struct {
	repo   UserRepository
	Crypto crypto.Crypto
	Config string
}

func (u *userService) Add(dto AddUserDto) (UserDto, error) {
	user, err := dto.MapToModel()
	if err != nil {
		return UserDto{}, err
	}

	if err := user.HashPassword(u.Crypto); err != nil {
		return UserDto{}, err
	}
	return u.repo.Add(user)
}

func (u *userService) GetMe(token string) (User, error) {

	part := strings.Split(token, "Bearer ")

	decode, err := u.Crypto.ParseAndValidateJWT(part[1], u.Config)
	if err != nil {
		return User{}, err
	}

	return u.repo.GetMe(uint(decode["userId"].(float64)))
}

func (u *userService) UpdateConfig(newconfig UpdateBody) (ConfigUser, error) {
	return u.repo.UpdateConfig(newconfig)
}
