package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user User) (UserDto, error)
	GetByEmail(email string) (User, error)
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
	tx := r.db.Create(&user)

	if err := tx.Error; err != nil {

		if tx.Error.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			return UserDto{}, errors.New("user already exists")
		}

		return UserDto{}, err
	}

	userResponse := UserDto{Id: user.Id, Username: user.Username, Email: user.Email}
	return userResponse, nil
}

func (r *userRepository) GetByEmail(email string) (User, error) {
	var user User
	tx := r.db.Where("email = ?", email).First(&user)
	return user, tx.Error
}
