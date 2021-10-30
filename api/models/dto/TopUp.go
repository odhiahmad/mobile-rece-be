package dto

type TopUpRequest struct {
	Amount 			uint64
	UserMerchantId 	string
}

type TopUpResponse struct {
	TransactionCode string
	Amount 			uint64
	AdminFee		uint64
	Total			uint64
}
