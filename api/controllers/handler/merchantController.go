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

type merchantController struct {
	router    *mux.Router
	parseJson *httpParse.JsonParse
	responder httpResponse.IResponder
	service   usecase.IMerchantUseCase
}

func NewMerchantController(router *mux.Router, parse *httpParse.JsonParse, responder httpResponse.IResponder, service usecase.IMerchantUseCase) controllers.IDelivery {
	return &merchantController{
		router,
		parse,
		responder,
		service,
	}
}

func (a *merchantController) InitRoute(mdw ...mux.MiddlewareFunc) {
	u := a.router.PathPrefix("/merchants").Subrouter()
	u.Use(mdw...)
	u.HandleFunc("", a.SaveMerchant).Methods("POST")
	u.HandleFunc("", a.GetAllMerchant).Methods("GET")
}

func (a *merchantController) SaveMerchant(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var merchant models.Merchant

	err = json.Unmarshal(body, &merchant)
	if err != nil {
		a.responder.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	merchantCreated, err := a.service.SaveMerchant(&merchant)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		a.responder.Error(w, http.StatusInternalServerError, formattedError.Error())
		return
	}

	a.responder.Data(w, http.StatusCreated, status.StatusText(status.CREATED), merchantCreated)
}

func (a *merchantController) GetAllMerchant(w http.ResponseWriter, r *http.Request) {
	merchants, err := a.service.FindAllMerchant()
	if err != nil {
		a.responder.Error(w, http.StatusNotFound, err.Error())
		return
	}
	a.responder.Data(w, status.Success, status.StatusText(status.Success), merchants)
}
