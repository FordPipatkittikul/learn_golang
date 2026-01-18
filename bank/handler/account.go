package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accountService service.AccountService
}

func NewAccountHanlder(accountService service.AccountService) accountHandler {
	return accountHandler{accountService : accountService}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	// validate header
	if r.Header.Get("content-type") != "application/json"{
		handleError(w, errs.NewValidationError("request header incorrect format"))
		return
	}

	// validate body
	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handleError(w, errs.NewValidationError("request body incorrect format"))
		return
	}

	res, err := h.accountService.NewAccount(customerID, request)
	if err != nil {
		handleError(w, err)
		return
	}
	
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) { 
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	res, err := h.accountService.GetAccounts(customerID)
	if err != nil {
		handleError(w, err)
		return 
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}