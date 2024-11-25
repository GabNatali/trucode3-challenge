package database

import (
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
		return fmt.Errorf("failed to connect to database: %w AND error: %v", err, s.dns)
	}

	s.DB = db
	return nil
}
