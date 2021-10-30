package dto

type TransferRequest struct {
	Amount         uint64
	UserMerchantId string
}

type TransferResponse struct {
	TransactionCode string
	Amount          uint64
	AdminFee        uint64
	Total           uint64
}
