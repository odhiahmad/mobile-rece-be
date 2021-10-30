package dto

type WithDrawRequest struct {
	Amount         uint64
	UserMerchantId string
}

type WithDrawResponse struct {
	CodeTransaction string
	Amount 			uint64
	AdminFee		uint64
	Total 			uint64

}