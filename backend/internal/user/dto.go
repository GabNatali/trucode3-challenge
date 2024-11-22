package user

type UserDto struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (dto UserDto) MapFromModel(user User) UserDto {
	dto.Id = user.Id
	dto.Username = user.Username
	dto.Email = user.Email

	return dto
}

type AddUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto AddUserDto) MapToModel() (User, error) {
	return NewUser(
		dto.Username,
		dto.Email,
		dto.Password,
	)
}
