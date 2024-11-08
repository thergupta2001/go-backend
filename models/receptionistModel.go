package models

import (
    "errors"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
    "time"
)

type Receptionist struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    Name      string    `gorm:"size:255;not null"`
    Password  string    `gorm:"size:255;not null"`
    Role      string    `gorm:"size:20;not null"`
    CreatedAt time.Time
}

func (r *Receptionist) BeforeSave(tx *gorm.DB) (err error) {
    if r.Role == "" {
        r.Role = ReceptionistRole
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
    if err != nil {
		return err
	}
	r.Password = string(hashedPassword)

    if !ValidRoles[r.Role] {
        return errors.New("invalid role")
    }
    return
}
