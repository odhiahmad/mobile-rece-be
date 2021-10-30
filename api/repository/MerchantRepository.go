package repository

import (
	"bri-rece/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type IMerchatRepository interface {
	SaveMerchant(merchant *models.Merchant) (*models.Merchant, error)
	FindAllMerchant() (*[]models.Merchant, error)
	FindMerchantByIdRece(id string) (*models.Merchant, error)
	FindMerchantByIdBill(id string) (*models.Merchant, error)

}

type merchantRepositoryImpl struct {
	db *gorm.DB
}

func (u *merchantRepositoryImpl) FindMerchantByIdBill(id string) (*models.Merchant, error) {
	uid , _ := uuid.FromString(id)
	var merchant models.Merchant
	if err := u.db.Debug().Table("merchants").Where("id = ?", uid).Where("status = 1").Find(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func (u *merchantRepositoryImpl) FindMerchantByIdRece(id string) (*models.Merchant, error) {

	uid , _ := uuid.FromString(id)
	var merchant models.Merchant
	if err := u.db.Debug().Table("merchants").Where("id = ?", uid).Where("status = 0").Find(&merchant).Error; err != nil {
		return nil, err
	}
	return &merchant, nil
}

func NewMerchantRepository(db *gorm.DB) IMerchatRepository {
	return &merchantRepositoryImpl{
		db,
	}
}

func (u *merchantRepositoryImpl) SaveMerchant(newMerchant *models.Merchant) (*models.Merchant, error) {
	var err error
	err = u.db.Debug().Create(&newMerchant).Error
	if err != nil {
		return &models.Merchant{}, err
	}
	return newMerchant, nil
}

func (u *merchantRepositoryImpl) FindAllMerchant() (*[]models.Merchant, error) {
	var err error
	var merchants []models.Merchant
	err = u.db.Debug().Model(&models.Merchant{}).Find(&merchants).Error
	if err != nil {
		return &merchants, err
	}
	return &merchants, nil
}
