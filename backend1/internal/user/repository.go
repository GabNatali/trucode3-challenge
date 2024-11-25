package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user User) (UserDto, error)
	GetByEmail(email string) (User, error)
	GetMe(id uint) (User, error)
	UpdateConfig(newconfig UpdateBody) (ConfigUser, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Add(user User) (UserDto, error) {

	newUser := User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		ConfigUser: ConfigUser{
			UserId:        user.Id,
			Education:     "",
			Income:        "",
			MaritalStatus: "",
			MaxAge:        0,
			MinAge:        0,
			Occupation:    "",
			OrderBy:       "age",
			OrderDir:      "asc",
		},
	}

	tx := r.db.Create(&newUser)

	if err := tx.Error; err != nil {

		if tx.Error.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			return UserDto{}, errors.New("user already exists")
		}

		return UserDto{}, err
	}

	userResponse := UserDto{Id: newUser.Id, Username: newUser.Username, Email: newUser.Email}
	return userResponse, nil
}

func (r *userRepository) GetByEmail(email string) (User, error) {
	var user User
	tx := r.db.Where("email = ?", email).First(&user)
	return user, tx.Error
}

func (r *userRepository) GetMe(id uint) (User, error) {
	var user User
	tx := r.db.Preload("ConfigUser").First(&user, id)
	if tx.Error != nil {
		return User{}, tx.Error
	}

	return user, nil
}
func (r *userRepository) UpdateConfig(newconfig UpdateBody) (ConfigUser, error) {
	var config ConfigUser
	tx := r.db.Where("user_id = ?", newconfig.UserId).First(&config)
	if tx.Error != nil {
		return ConfigUser{}, tx.Error
	}

	fmt.Println("neew config")
	fmt.Println(newconfig)
	tx = r.db.Model(&config).Select("Education", "Income", "MaritalStatus", "Occupation").Updates(newconfig)
	if tx.Error != nil {
		return ConfigUser{}, tx.Error
	}

	return config, tx.Error
}
