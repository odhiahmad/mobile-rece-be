package dto

type BillRequest struct {
	Amount 			uint64
	UserMerchantId 	string
	SourceOfFunds 	string
}

type BillResponse struct {
	TransactionCode string
	Amount 			uint64
	AdminFee 		uint64
	Total 			uint64
}