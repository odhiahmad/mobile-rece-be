package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserMerchant struct {
	ID         uuid.UUID  `gorm:"type:uuid;unique;index;"`
	Name       string     `gorm:"not null; size:255; column:name"`
	VaNumber   string     `gorm:"not null; size:255; column:va_number"`
	Balance    uint64     `gorm:"not null;" `
	AdminFee   uint64	  `gorm:"not null"`
	MerchantID uuid.UUID  `gorm:"not null"`
	Merchant   Merchant   `gorm:"foreignKey:MerchantID;references:ID" `
	CreatedAt  time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt  time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" `
	DeletedAt  *time.Time `gorm:"column:deleted_at" sql:"index"`
}

func (u *UserMerchant) Prepare() {
	u.ID = uuid.NewV4()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
