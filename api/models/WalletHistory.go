package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type statusTransaction string

const (
	WITHDRAW statusTransaction = "WITHDRAW"
	TRANSFER statusTransaction = "TRANSFER"
	TOPUP    statusTransaction = "TOPUP"
	PAYMENT  statusTransaction = "PAYMENT"
)

type WalletHistory struct {
	ID              uuid.UUID         `gorm:"primary_key;unique;index;type:uuid" `
	TransactionDate time.Time         `gorm:"not null;type:date"`
	Amount          uint64            `gorm:"not null" `
	TypeTransaction statusTransaction `gorm:"not null;" `
	WalletID        uuid.UUID         `gorm:"not null"`
	Wallet          *Wallet           `gorm:"foreignKey:WalletID;references:ID" `
	UserMerchantID  uuid.UUID         `gorm:"null"`
	UserMerchant    *UserMerchant     `gorm:"foreignKey:UserMerchantID;references:ID" `
	TransactionCode string            `gorm:"not null" `
	CreatedAt       time.Time         `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt       time.Time         `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" `
	DeletedAt       *time.Time        `gorm:"column:deleted_at" sql:"index"`
}

type WalletHistoryResponse struct {
	ID              uuid.UUID
	TransactionDate time.Time
	Amount          uint64
	TypeTransaction statusTransaction
}

func (wh *WalletHistory) Prepare() {
	wh.ID = uuid.NewV4()
	wh.CreatedAt = time.Now()
	wh.UpdatedAt = time.Now()
}
