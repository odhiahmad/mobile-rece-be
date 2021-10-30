package repository

import (
	"bri-rece/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IWalletHistoryRepo interface {
	CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error)
	GetHistoryById(id string, page, size string) ([]*models.WalletHistory, error)
	GetAllHistory(page, pageSize, order string)([]*models.WalletHistory,error)
}

type walletHistoryRepo struct {
	db *gorm.DB
}

func (w *walletHistoryRepo) GetAllHistory(page, pageSize, order string) ([]*models.WalletHistory, error) {
	histories := make([]*models.WalletHistory, 0)
	if order != "" {
		if err := w.db.Debug().Scopes(Paginate(page, pageSize)).First(&histories).Limit(3).Error; err != nil {
			return nil, err
		}
		return histories, nil
	} else {
		if err := w.db.Debug().Scopes(Paginate(page, pageSize)).Order("created_at desc").Find(&histories).Limit(3).Error; err != nil {
			return nil, err
		}
		return histories, nil
	}

}

func (w *walletHistoryRepo) CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error) {
	var err error
	err = w.db.Debug().Create(&history).Error
	if err != nil {
		return &models.WalletHistory{}, err
	}
	return history, nil
}

func (w *walletHistoryRepo) GetHistoryById(id string, page, size string) ([]*models.WalletHistory, error) {
	uid, _ := uuid.FromString(id)
	var err error
	history := make([]*models.WalletHistory,0)
	if err = w.db.Debug().Table("wallet_histories").Where("wallet_id = ?", uid).Scopes(Paginate(page,size)).Order("created_at desc").Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}



func NewWalletHistoryRepo(db *gorm.DB) IWalletHistoryRepo  {
	return &walletHistoryRepo{
		db,
	}
}