package repository

import (
	"bri-rece/api/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

type IUserMerchantRepository interface {
	SaveUserMerchant(userMerchant *models.UserMerchant) (*models.UserMerchant, error)
	FindUserMerchantByName(name string) (*models.UserMerchant, error)
	FindUserMerchantByVaNumber(vanumber string) (*models.UserMerchant, error)
	FindUserMerchantById(id string) (*models.UserMerchant, error)
	FindAllUserMerchant()([]*models.UserMerchant, error)
}

type userMerchantRepositoryImpl struct {
	db *gorm.DB
}

func (u *userMerchantRepositoryImpl) FindAllUserMerchantByBill() ([]*models.UserMerchant, error) {
	userMerchant := make([]*models.UserMerchant,0)
	if err := u.db.Debug().Table("user_merchants").Joins("join merchants on user_merchants.merchant_id = merchants.merchant_id").Where("merchants.status = 1").Find(&userMerchant).Error; err != nil {
		return nil, err
	}
	return userMerchant, nil
}

func (u *userMerchantRepositoryImpl) FindAllUserMerchant() ([]*models.UserMerchant, error) {
	userMerchant := make([]*models.UserMerchant,0)
	if err := u.db.Debug().Table("user_merchants").Preload(clause.Associations).Find(&userMerchant).Error;err != nil {
		return nil, err
	}
	return userMerchant,nil
}

func NewUserMerchantRepository(db *gorm.DB) IUserMerchantRepository {
	return &userMerchantRepositoryImpl{
		db,
	}
}

func (u *userMerchantRepositoryImpl) SaveUserMerchant(newUserMerchant *models.UserMerchant) (*models.UserMerchant, error) {
	var err error
	err = u.db.Debug().Create(&newUserMerchant).Error
	if err != nil {
		return &models.UserMerchant{}, err
	}
	return newUserMerchant, nil
}

func (u *userMerchantRepositoryImpl) FindUserMerchantByName(name string) (*models.UserMerchant, error) {
	var userMerchants models.UserMerchant
	err := u.db.Debug().Model(&models.UserMerchant{}).First(&userMerchants, map[string]interface{}{"name": name}).Error
	if err != nil {
		return &userMerchants, err
	}
	return &userMerchants, nil
}

func (u *userMerchantRepositoryImpl) FindUserMerchantByVaNumber(vanumber string) (*models.UserMerchant, error) {
	var userMerchants models.UserMerchant
	err := u.db.Debug().Model(&models.UserMerchant{}).First(&userMerchants, map[string]interface{}{"va_number": vanumber}).Error
	if err != nil {
		return &userMerchants, err
	}
	return &userMerchants, nil
}

func (u *userMerchantRepositoryImpl) FindUserMerchantById(id string) (*models.UserMerchant, error) {
	uid, _ := uuid.FromString(id)
	var userMerchant models.UserMerchant
	if err := u.db.Debug().Table("user_merchants").Where("id = ?", uid).Preload(clause.Associations).Find(&userMerchant).Error; err != nil {
		return nil, err
	}
	return &userMerchant, nil
}

