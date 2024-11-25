package user

import (
	"errors"

	"github.com/GabNatali/trucode3-challenge-final/internal/crypto"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm"
)

type User struct {
	Id         uint       `gorm:"primaryKey" json:"id"`
	Username   string     `gorm:"type:char(50);not null" json:"username"`
	Email      string     `gorm:"type:char(50);not null;unique" json:"email"`
	Password   string     `gorm:"type:char(150);not null" json:"-"`
	ConfigUser ConfigUser `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;" json:"config"`
}

type ConfigUser struct {
	Id            uint   `gorm:"primaryKey" json:"id"`
	UserId        uint   `gorm:"not null;unique" json:"-"`
	Education     string `json:"education" gorm:"type:text"`
	Income        string `json:"income" gorm:"type:text"`
	MaritalStatus string `json:"marital_status" gorm:"type:text"`
	MaxAge        int    `json:"max_age" gorm:"type:int"`
	MinAge        int    `json:"min_age" gorm:"type:int"`
	Occupation    string `json:"occupation" gorm:"type:text"`
	OrderBy       string `json:"order_by" gorm:"type:varchar(50)"`
	OrderDir      string `json:"order_dir" gorm:"type:varchar(4)"`
	gorm.Model    `json:"-"`
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
