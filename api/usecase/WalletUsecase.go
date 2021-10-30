package usecase

import (
	"bri-rece/api/models"
	"bri-rece/api/models/dto"
	"bri-rece/api/repository"
	"crypto/rand"
	"io"
	"time"

	uuid "github.com/satori/go.uuid"
)

type IWalletUsecase interface {
	CreateWallet(wallet *models.Wallet) (*models.Wallet, error)
	TopUp(wallet *dto.TopUpRequest, id string) (*dto.TopUpResponse, error)
	WithDraw(wallet *dto.WithDrawRequest, id string) (*dto.WithDrawResponse, error)
	Transfer(wallet *dto.TransferRequest, id string) (*dto.TransferResponse, error)
	GetWalletById(id string) (*models.Wallet, error)
	Bill(wallet *dto.BillRequest, id string) (*dto.BillResponse, error)
}

type walletUsecaseRepo struct {
	walletRepo   repository.IWalletRepository
	history      IWalletHistoryUsecase
	userMerchant IUserMerchantUseCase
}

func (w *walletUsecaseRepo) Bill(wallet *dto.BillRequest, id string) (*dto.BillResponse, error) {

	var err error
	idWallet, _ := uuid.FromString(id)
	userMerchantId , _ := uuid.FromString(wallet.UserMerchantId)
	codeWithDraw := encodeToString(6)

	_, err = w.userMerchant.FindUserMerchantById(wallet.UserMerchantId)
	if err != nil {
		return nil, err
	}

	bill := models.Wallet{
		Balance: wallet.Amount + 2500,
		UpdatedAt: time.Now(),
	}
	_, err = w.walletRepo.Bill(&bill, id)
	if err != nil {
		return nil, err
	}

	histories := models.WalletHistory{
		TransactionDate: time.Now(),
		Amount:          wallet.Amount,
		WalletID:        idWallet,
		UserMerchantID:  userMerchantId,
		TypeTransaction: models.PAYMENT,
		TransactionCode: codeWithDraw,
		CreatedAt:       time.Now(),
	}
	histories.Prepare()
	_,err = w.history.CreateHistory(&histories)
	if err != nil {
		return nil, err
	}

	return &dto.BillResponse{
		TransactionCode: codeWithDraw,
		Amount: wallet.Amount + 2500,
		AdminFee: 2500,
		Total: 2500 + bill.Balance,
	}, nil
}

func (w *walletUsecaseRepo) GetWalletById(id string) (*models.Wallet, error) {
	return w.walletRepo.FindById(id)
}

func (w *walletUsecaseRepo) WithDraw(wallet *dto.WithDrawRequest, id string) (*dto.WithDrawResponse, error) {
	var err error
	idWallet, _ := uuid.FromString(id)
	userMerchantId, _ := uuid.FromString(wallet.UserMerchantId)
	codeWithDraw := encodeToString(6)

	withDraw := models.Wallet{
		Balance: wallet.Amount+2500,
		UpdatedAt: time.Now(),
	}
	_, err = w.walletRepo.WithDraw(&withDraw, id)
	if err != nil {
		return nil, err
	}

	histories := models.WalletHistory{
		TransactionDate: time.Now(),
		Amount:          wallet.Amount,
		WalletID:        idWallet,
		UserMerchantID:  userMerchantId,
		TypeTransaction: models.WITHDRAW,
		TransactionCode: codeWithDraw,
		CreatedAt:       time.Now(),
	}
	histories.Prepare()
	_, err = w.history.CreateHistory(&histories)
	if err != nil {
		return nil, err
	}

	return &dto.WithDrawResponse{
		CodeTransaction: codeWithDraw,
		Amount:          wallet.Amount,
		AdminFee:        2500,
		Total:           2500 + wallet.Amount,
	}, nil
}

func (w *walletUsecaseRepo) CreateWallet(wallet *models.Wallet) (*models.Wallet, error) {
	return w.walletRepo.CreateWallet(wallet)
}

func (w *walletUsecaseRepo) TopUp(topUp *dto.TopUpRequest, id string) (*dto.TopUpResponse, error) {

	var err error
	uid, _ := uuid.FromString(id)
	userMerchantId, _ := uuid.FromString(topUp.UserMerchantId)
	transactionCode := encodeToString(6)

	findUserMerchant, errUserMerchant := w.userMerchant.FindUserMerchantById(topUp.UserMerchantId)
	if errUserMerchant != nil {
		return nil, errUserMerchant
	}

	wallet := models.Wallet{
		ID: uid,
		Balance: topUp.Amount - findUserMerchant.AdminFee,
		UpdatedAt: time.Now(),
	}
	_, err = w.walletRepo.UpdateBalance(&wallet, id)
	if err != nil {
		return nil, err
	}

	history := &models.WalletHistory{
		Amount:          topUp.Amount + findUserMerchant.AdminFee,
		TypeTransaction: models.TOPUP,
		TransactionCode: transactionCode,
		WalletID:        uid,
		UserMerchantID:  userMerchantId,
	}
	history.Prepare()
	_, err = w.history.CreateHistory(history)
	if err != nil {
		return nil, err
	}

	return &dto.TopUpResponse{
		TransactionCode: history.TransactionCode,
		Amount: topUp.Amount,
		AdminFee: findUserMerchant.AdminFee,
		Total: topUp.Amount - findUserMerchant.AdminFee,
	}, nil
}

func (w *walletUsecaseRepo) Transfer(transfer *dto.TransferRequest, id string) (*dto.TransferResponse, error) {

	var err error
	uid, _ := uuid.FromString(id)
	userMerchantId, _ := uuid.FromString(transfer.UserMerchantId)
	transactionCode := encodeToString(6)

	wallet := models.Wallet{
		ID:        uid,
		Balance:   transfer.Amount,
		UpdatedAt: time.Now(),
	}
	_, err = w.walletRepo.Transfer(&wallet, id)
	if err != nil {
		return nil, err
	}

	findUserMerchant, errUserMerchant := w.userMerchant.FindUserMerchantById(transfer.UserMerchantId)
	if errUserMerchant != nil {
		return nil, errUserMerchant
	}

	history := &models.WalletHistory{
		Amount:          transfer.Amount + findUserMerchant.AdminFee,
		TypeTransaction: models.TRANSFER,
		TransactionCode: transactionCode,
		WalletID:        uid,
		UserMerchantID:  userMerchantId,
	}
	history.Prepare()
	_, err = w.history.CreateHistory(history)
	if err != nil {
		return nil, err
	}

	return &dto.TransferResponse{
		TransactionCode: history.TransactionCode,
		Amount:          transfer.Amount,
		AdminFee:        findUserMerchant.AdminFee,
		Total:           transfer.Amount + findUserMerchant.AdminFee,
	}, nil
}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func NewWalletUsecase(walletRepo repository.IWalletRepository, history repository.IWalletHistoryRepo, userMerchant repository.IUserMerchantRepository) IWalletUsecase {
	return &walletUsecaseRepo{
		walletRepo,
		history,
		userMerchant,
	}
}
