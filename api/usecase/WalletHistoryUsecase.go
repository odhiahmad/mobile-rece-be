package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/repository"
	"time"
)

type IWalletHistoryUsecase interface {
	CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error)
	GetAllHistory(page, pageSize, order string) ([]*models.WalletHistory, error)
	GetHistoryById(id string, page, size string) ([]*models.WalletHistory, error)
}

type walletHistoryUsecase struct {
	walletHistory repository.IWalletHistoryRepo
}

func (w *walletHistoryUsecase) GetHistoryById(id string, page, size string) ([]*models.WalletHistory, error) {
	return w.walletHistory.GetHistoryById(id, page, size)
}

func (w *walletHistoryUsecase) GetAllHistory(page, pageSize, order string) ([]*models.WalletHistory, error) {
	return w.walletHistory.GetAllHistory(page, pageSize,order)
}

func (w *walletHistoryUsecase) CreateHistory(history *models.WalletHistory) (*models.WalletHistory, error) {
	trxHistory := &models.WalletHistory{
		Amount: history.Amount,
		TypeTransaction: history.TypeTransaction,
		WalletID: history.WalletID,
		TransactionDate: time.Now(),
		UserMerchantID: history.UserMerchantID,
		TransactionCode: history.TransactionCode,
	}
	trxHistory.Prepare()
	return w.walletHistory.CreateHistory(trxHistory)
}

func NewWalletHistoryUsecase(walletHistory repository.IWalletHistoryRepo) IWalletHistoryUsecase {
	return &walletHistoryUsecase{
		walletHistory,
	}
}