package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Doctor struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:20;not null"`
	CreatedAt time.Time
}

func (d *Doctor) BeforeSave(tx *gorm.DB) (err error) {
	if d.Role == "" {
		d.Role = DoctorRole
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	d.Password = string(hashedPassword)

	if !ValidRoles[d.Role] {
		return errors.New("invalid role")
	}
	return
}
