package controllers

import (
	"encoding/json"
	"fmt"
	"linkAja/master/usecases"
	"linkAja/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct{
	accountUsecase usecases.AccountUsecase
}

func AccountController(r *mux.Router,service usecases.AccountUsecase){
	accountHandler := AccountHandler{accountUsecase: service}

	account := r.PathPrefix("/account/{accNum}").Subrouter()
	account.HandleFunc("",accountHandler.GetSaldo).Methods(http.MethodGet)
	account.HandleFunc("/transfer",accountHandler.Transfer).Methods(http.MethodPost)
}

func (s *AccountHandler)GetSaldo(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	accNum := params["accNum"]

	account,err:=s.accountUsecase.GetSaldo(accNum)
	if err!= nil{
		log.Println(err)
	}
	var msg models.Message
	byteData,err := json.Marshal(account)

	
	if err!=nil{
		byteMsg,_ := json.Marshal(msg)
		msg.Msg="Something Wrong on Marshalling Data"
		w.Write(byteMsg)
		return
	}
	if account == nil{
		msg.Msg=fmt.Sprintf("Account Number %s Not Found",accNum)
		byteMsg,_ := json.Marshal(msg)	
		
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(byteMsg)
	}else{
		w.Header().Set("Content-type", "application/json")
		w.Write(byteData)
	}
}

func (s *AccountHandler)Transfer(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	accNum := params["accNum"]

	var msg models.Message

	var input *models.Transfer
	_ = json.NewDecoder(r.Body).Decode(&input)
	err := s.accountUsecase.Transfer(accNum,input)

	if err!=nil{
		msg.Msg = err.Error()
		byteMsg,_ := json.Marshal(msg)
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(byteMsg)
	}else{
		w.WriteHeader(http.StatusCreated)
	}
}