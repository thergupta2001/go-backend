package models

import (
    "errors"
    "gorm.io/gorm"
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
    if !ValidRoles[r.Role] {
        return errors.New("invalid role")
    }
    return
}
