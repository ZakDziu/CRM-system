package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID      uuid.UUID `json:"-"`
	UserID  uuid.UUID `json:"-"`
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.NewV4().String()
	tx.Statement.SetColumn("ID", uuid)

	return nil
}
