package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/middlewares"
	"bri-rece/api/models/dto"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"net/http"

	"github.com/gorilla/mux"
)

type WalletController struct {
	router         *mux.Router
	parseJson      *httpParse.JsonParse
	responder      httpResponse.IResponder
	service        usecase.IWalletUsecase
	serviceHistory usecase.IWalletHistoryUsecase
}

func (t *WalletController) InitRoute(mdw ...mux.MiddlewareFunc) {
	wallet := t.router.PathPrefix("/wallet").Subrouter()
	wallet.Use(mdw...)
	wallet.HandleFunc("/topup/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.TopUp))).Methods("PUT")
	wallet.HandleFunc("/withdraw/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.WithDraw))).Methods("PUT")
	wallet.HandleFunc("/transfer/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.Transfer))).Methods("PUT")
	wallet.HandleFunc("/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(t.GetWalletById))).Methods("GET")
}

func NewWalletController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IWalletUsecase, serviceHistory usecase.IWalletHistoryUsecase) controllers.IDelivery {
	return &WalletController{
		router,
		parse,
		responder,
		service,
		serviceHistory,
	}
}

func (t *WalletController) TopUp(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	var wallet dto.TopUpRequest
	if err := t.parseJson.Parse(r, &wallet); err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	topUpData, err := t.service.TopUp(&wallet, param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	t.responder.Data(w, status.Success, status.StatusText(status.Success), topUpData)
}

func (t *WalletController) WithDraw(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var wallet dto.WithDrawRequest
	if err := t.parseJson.Parse(r, &wallet); err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	withDraw, err := t.service.WithDraw(&wallet, param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	t.responder.Data(w, status.Success, status.StatusText(status.Success), withDraw)
}

func (t *WalletController) Transfer(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	var wallet dto.TransferRequest
	if err := t.parseJson.Parse(r, &wallet); err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	transferData, err := t.service.Transfer(&wallet, param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	t.responder.Data(w, status.Success, status.StatusText(status.Success), transferData)
}

func (t *WalletController) GetWalletById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	walletId, err := t.service.GetWalletById(param["id"])
	if err != nil {
		t.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	t.responder.Data(w, status.Success, status.StatusText(status.Success), walletId)
}
