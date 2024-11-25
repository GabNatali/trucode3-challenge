package user

type UserDto struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserInfo struct {
	Id            uint   `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Education     string `json:"education"`
	Income        string `json:"income"`
	MaritalStatus string `json:"marital_status"`
	MaxAge        int    `json:"max_age"`
	MinAge        int    `json:"min_age"`
	Occupation    string `json:"occupation"`
	OrderBy       string `json:"order_by"`
	OrderDir      string `json:"order_dir"`
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
