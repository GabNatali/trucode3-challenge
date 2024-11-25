package auth

import "github.com/GabNatali/trucode3-challenge-final/internal/user"

type LoginUserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggedUserDto struct {
	user.UserDto
	Token string `json:"token"`
}

func (dto LoggedUserDto) MapFromModel(model user.User, token string) LoggedUserDto {
	dto.Id = model.Id
	dto.Username = model.Username
	dto.Email = model.Email
	dto.Token = token

	return dto
}
