package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB  *gorm.DB
	dns string
}

func NewService(dns string) *Service {
	return &Service{
		dns: dns,
	}
}

func (s *Service) Connect() error {
	fmt.Println("connecting to database: " + s.dns)
	db, err := gorm.Open(postgres.Open(s.dns), &gorm.Config{})
	if err != nil {
		return errors.New("cannot connect to database")
	}

	s.DB = db
	return nil
}
