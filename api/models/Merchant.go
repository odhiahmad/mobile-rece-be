package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type isStatus int

const (
	RECE 	isStatus = iota
	BILL
)

type Merchant struct {
	ID           uuid.UUID  `gorm:"type:uuid;unique;index; column:merchant_id"`
	Name         string     `gorm:"not null; size:255; column:name"`
	TypeMerchant string     `gorm:"not null; size:255; column:type_merchant"`
	Status		 isStatus	`gorm:"null; column:status"`
	CreatedAt    time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"column:deleted_at" json:"deleted_at" sql:"index"`
}

func (u *Merchant) Prepare() {
	u.ID = uuid.NewV4()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
