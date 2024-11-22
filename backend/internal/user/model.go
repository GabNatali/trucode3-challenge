package user

import (
	"errors"

	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:char(50);not null"`
	Email    string `gorm:"type:char(50);not null;unique"`
	Password string `gorm:"type:char(150);not null"`
	gorm.Model
}

func NewUser(username, email, password string) (User, error) {
	user := User{
		Username: username,
		Email:    email,
		Password: password,
	}

	if err := user.Validate(); err != nil {
		return User{}, err
	}

	return user, nil

}

func (user *User) HashPassword(crypto crypto.Crypto) error {
	passwordHashed, err := crypto.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = passwordHashed

	return nil
}

func (user *User) ComparePassword(password string, crypto crypto.Crypto) bool {
	return crypto.CompareHashAndPassword(user.Password, password)
}

func (user *User) Validate() error {
	err := validation.ValidateStruct(user,
		validation.Field(&user.Username, validation.Required, validation.Length(2, 100)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 8)),
	)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
