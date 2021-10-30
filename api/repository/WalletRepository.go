package repository

import (
	"bri-rece/api/models"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IWalletRepository interface {
	CreateWallet(wallet *models.Wallet) (*models.Wallet, error)
	UpdateBalance(wallet *models.Wallet, id string) (*models.Wallet, error)
	FindById(id string) (*models.Wallet, error)
	WithDraw(wallet *models.Wallet, id string) (*models.Wallet, error)
	Transfer(wallet *models.Wallet, id string) (*models.Wallet, error)
	Bill(wallet *models.Wallet, id string) (*models.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

func (w *walletRepository) Bill(wallet *models.Wallet, id string) (*models.Wallet, error) {
	temp, err := w.FindById(id)
	result := temp.Balance - wallet.Balance
	if result >= 0 {
		w.db = w.db.Debug().Model(&models.Wallet{}).Where("id = ?", id).Take(&models.Wallet{}).UpdateColumns(
			map[string]interface{}{
				"balance":    result,
				"updated_at": time.Now(),
			},
		)
		if w.db.Error != nil {
			return &models.Wallet{}, err
		}
		return &models.Wallet{}, nil
	} else {
		return nil, errors.New("Amount is not sufficient")
	}
}

func (w *walletRepository) WithDraw(wallet *models.Wallet, id string) (*models.Wallet, error) {
	temp, err := w.FindById(id)
	result := temp.Balance - wallet.Balance
	if result >= 0 {
		w.db = w.db.Debug().Model(&models.Wallet{}).Where("id = ?", id).Take(&models.Wallet{}).UpdateColumns(
			map[string]interface{}{
				"balance":    result,
				"updated_at": time.Now(),
			},
		)
		if w.db.Error != nil {
			return &models.Wallet{}, err
		}
		return &models.Wallet{}, nil
	} else {
		return nil, errors.New("Amount is not sufficient")
	}
}

func (w *walletRepository) Transfer(wallet *models.Wallet, id string) (*models.Wallet, error) {
	temp, err := w.FindById(id)
	result := temp.Balance - wallet.Balance
	if result >= 0 {
		w.db = w.db.Debug().Model(&models.Wallet{}).Where("id = ?", id).Take(&models.Wallet{}).UpdateColumns(
			map[string]interface{}{
				"balance":    result,
				"updated_at": time.Now(),
			},
		)
		if w.db.Error != nil {
			return &models.Wallet{}, err
		}
		return &models.Wallet{}, nil
	} else {
		return nil, errors.New("Amount is not sufficient")
	}
}

func NewWalletRepository(db *gorm.DB) IWalletRepository {
	return &walletRepository{
		db,
	}
}

func (w *walletRepository) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	var err error
	err = w.db.Debug().Create(&wallet).Error
	if err != nil {
		return &models.Wallet{}, err
	}
	return wallet, nil
}

func (w *walletRepository) UpdateBalance(wallet *models.Wallet, id string) (*models.Wallet, error) {
	temp, err := w.FindById(id)
	result := wallet.Balance + temp.Balance
	w.db = w.db.Debug().Model(&models.Wallet{}).Where("id = ?", id).Find(&models.Wallet{}).UpdateColumns(
		map[string]interface{}{
			"balance":    result,
			"updated_at": time.Now(),
		},
	)
	if w.db.Error != nil {
		return &models.Wallet{}, err
	}
	return &models.Wallet{}, nil
}

func (w *walletRepository) FindById(id string) (*models.Wallet, error) {
	a, _ := uuid.FromString(id)
	var walletDb models.Wallet
	if err := w.db.Debug().Table("wallets").Where("id = ?", a).Preload("WalletHistory").Find(&walletDb).Error; err != nil {
		return nil, err
	}
	return &walletDb, nil

}
