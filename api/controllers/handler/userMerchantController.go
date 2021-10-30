package handler

import (
	"bri-rece/api/controllers"
	"bri-rece/api/models"
	"bri-rece/api/usecase"
	"bri-rece/api/utils/formaterror"
	"bri-rece/api/utils/httpParse"
	"bri-rece/api/utils/httpResponse"
	"bri-rece/api/utils/status"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type userMerchantController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IUserMerchantUseCase
}

func NewUserMerchantController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IUserMerchantUseCase) controllers.IDelivery {
	return &userMerchantController{
		router,
		parse,
		responder,
		service,
	}
}

func (a *userMerchantController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := a.router.PathPrefix("/usermerchants").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", a.SaveUserMerchant).Methods("POST")
	u.HandleFunc("/{vanumber}", a.FindUserMerchantByVaNumber).Methods("GET")
	u.HandleFunc("", a.FindAllUserMerchant).Methods("GET")
	// u.HandleFunc(":name", a.SaveUserMerchant).Methods("GET")
}

func (a *userMerchantController) SaveUserMerchant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var userMerchant models.UserMerchant

	err = json.Unmarshal(body, &userMerchant)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	userMerchantCreated, err := a.service.SaveUserMerchant(&userMerchant)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		a.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return
	}

	a.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), userMerchantCreated)
}

func (a *userMerchantController) FindUserMerchantByVaNumber(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	userMerchants, err := a.service.FindUserMerchantByVaNumber(param["vanumber"])
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, status.Success, status.StatusText(status.Success), userMerchants)
}

func (a *userMerchantController) FindAllUserMerchant(w http.ResponseWriter, r *http.Request) {
	userMerchants, err := a.service.FindAllUserMerchant()
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, status.Success, status.StatusText(status.Success), userMerchants)
}