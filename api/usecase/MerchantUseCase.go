package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
)

type IMerchantUseCase interface {
	SaveMerchant(merchant *models.Merchant) (*models.Merchant, error)
	FindAllMerchant() (*[]models.Merchant, error)
	FindMerchantById(id string) (*models.Merchant, error)
}

type merchantUseCaseRepo struct {
	merchantRepo repository.IMerchatRepository
}

func (a *merchantUseCaseRepo) FindMerchantById(id string) (*models.Merchant, error) {
	return a.merchantRepo.FindMerchantByIdRece(id)
}

func (a *merchantUseCaseRepo) SaveMerchant(merchant *models.Merchant) (*models.Merchant, error) {
	merchant.Prepare()
	return a.merchantRepo.SaveMerchant(merchant)
}

func (a *merchantUseCaseRepo) FindAllMerchant() (*[]models.Merchant, error) {
	return a.merchantRepo.FindAllMerchant()
}

func NewMerchantUseCase(merchantRepo repository.IMerchatRepository) IMerchantUseCase {
	return &merchantUseCaseRepo{
		merchantRepo,
	}
}
