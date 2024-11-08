// models/doctor.go
package models

import (
    "errors"
    "gorm.io/gorm"
    "time"
)

type Doctor struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"`
    Name      string    `gorm:"size:255;not null"`
    Password  string    `gorm:"size:255;not null"`
    Role      string    `gorm:"size:20;not null"`
    CreatedAt time.Time
}

// Set default role and validate role
func (d *Doctor) BeforeSave(tx *gorm.DB) (err error) {
    if d.Role == "" {
        d.Role = DoctorRole
    }
    if !ValidRoles[d.Role] {
        return errors.New("invalid role")
    }
    return
}

