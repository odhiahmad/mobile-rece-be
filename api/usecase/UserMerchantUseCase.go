package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
)

type IUserMerchantUseCase interface {
	SaveUserMerchant(userMerchant *models.UserMerchant) (*models.UserMerchant, error)
	FindUserMerchantByName(name string) (*models.UserMerchant, error)
	FindUserMerchantByVaNumber(vanumber string) (*models.UserMerchant, error)
	FindUserMerchantById(id string) (*models.UserMerchant, error)
	FindAllUserMerchant()([]*models.UserMerchant, error)
}

type userMerchantUseCaseRepo struct {
	userMerchantRepo repository.IUserMerchantRepository
}

func (a *userMerchantUseCaseRepo) FindAllUserMerchant() ([]*models.UserMerchant, error) {
	return a.userMerchantRepo.FindAllUserMerchant()
}

func (a *userMerchantUseCaseRepo) FindUserMerchantById(id string) (*models.UserMerchant, error) {
	return a.userMerchantRepo.FindUserMerchantById(id)
}

func (a *userMerchantUseCaseRepo) SaveUserMerchant(userMerchant *models.UserMerchant) (*models.UserMerchant, error) {
	userMerchant.Prepare()
	return a.userMerchantRepo.SaveUserMerchant(userMerchant)
}

func (a *userMerchantUseCaseRepo) FindUserMerchantByName(name string) (*models.UserMerchant, error) {
	return a.userMerchantRepo.FindUserMerchantByName(name)
}

func (a *userMerchantUseCaseRepo) FindUserMerchantByVaNumber(vanumber string) (*models.UserMerchant, error) {
	return a.userMerchantRepo.FindUserMerchantByVaNumber(vanumber)
}



func NewUserMerchantUseCase(userMerchantRepo repository.IUserMerchantRepository) IUserMerchantUseCase {
	return &userMerchantUseCaseRepo{
		userMerchantRepo,
	}
}
