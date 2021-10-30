package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"github.com/gorilla/mux"
	"net/http"
)

type WalletHistories struct {
	router         *mux.Router
	parseJson      *httpParse.JsonParse
	responder      httpResponse.IResponder
	service		   usecase.IWalletHistoryUsecase
}

func (w *WalletHistories) InitRoute(mdw ...mux.MiddlewareFunc) {
	histories := w.router.PathPrefix("/histories").Subrouter()
	histories.Use(mdw...)
	histories.HandleFunc("/all", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(w.GetAllHistories))).Methods("GET")
	histories.HandleFunc("", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(w.GetHistoryByIdWallet))).Methods("GET")
}

func NewHistoriesController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IWalletHistoryUsecase) controllers.IDelivery {
	return &WalletHistories{
		router,
		parse,
		responder,
		service,
	}
}

func (wa *WalletHistories) GetAllHistories(w http.ResponseWriter, r *http.Request)  {
	page := r.Header.Get("page")
	size := r.Header.Get("size")
	order := r.Header.Get("order")
	wallets, err := wa.service.GetAllHistory(page, size, order)
	if err != nil {
		wa.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	wa.responder.Data(w, status.Success, status.StatusText(http.StatusOK), wallets)
}

func (wa *WalletHistories) GetHistoryByIdWallet(w http.ResponseWriter, r *http.Request)  {
	page := r.Header.Get("page")
	size := r.Header.Get("size")
	walletId := r.Header.Get("wallet")

	histories, err := wa.service.GetHistoryById(walletId,page,size)
	if err != nil {
		wa.responder.Error(w, http.StatusNotFound, err.Error())
	}
	wa.responder.Data(w, status.Success, status.StatusText(http.StatusOK), histories)
}