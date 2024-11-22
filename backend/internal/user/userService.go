package user

import "github.com/GabNatali/trucode3-challenge-final/internal/crypto"

type UserService interface {
	Add(dto AddUserDto) (UserDto, error)
}

func NewUserService(repo UserRepository, crypto crypto.Crypto) UserService {
	return &userService{
		repo:   repo,
		Crypto: crypto,
	}
}

type userService struct {
	repo   UserRepository
	Crypto crypto.Crypto
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
