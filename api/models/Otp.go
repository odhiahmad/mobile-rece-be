package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Otp struct {
	ID      uuid.UUID `gorm:"type:uuid;unique;index" json:"id"`
	OtpCode string    `gorm:"not null; size:255; column:otp_code"`
	UserID  uuid.UUID `gorm:"null"`
	User    *User     `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

type OtpRequest struct {
	UserId  string
	OtpCode string
}

func (o *Otp) Prepare()  {
	o.ID = uuid.NewV4()
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}