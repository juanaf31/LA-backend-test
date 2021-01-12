package controllers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"linkAja/master/repositories"
	"linkAja/master/usecases"
	"linkAja/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type WrongTransfer struct{
	Amount int `json:"amount"`
}

var mockTrans = &models.Transfer{
	Receiver: "555111",
	Amount: 1000,
}

var exceedMockTrans = &models.Transfer{
	Receiver: "555112",
	Amount: 100000,
}
var badMockTrans = &WrongTransfer{
	Amount: 100,
}



func DB_Test() *sql.DB{
	db,_ := sql.Open("sqlite3",":memory:")

	db.Exec("CREATE TABLE IF NOT EXISTS customer (customer_number TEXT PRIMARY KEY, name TEXT)")
	db.Exec("CREATE TABLE IF NOT EXISTS account (account_number TEXT PRIMARY KEY, customer_number TEXT, balance INTEGER)")

	db.Exec("INSERT INTO customer VALUES ('2002', 'Dummy 2')")
	db.Exec("INSERT INTO customer VALUES ('2001', 'Dummy 1'")
	db.Exec("INSERT INTO account VALUES ('555112','2002', 10000)")
	db.Exec("INSERT INTO account VALUES ('555111','2001', 20000)")
	return db
}

func mockRouter() *mux.Router{
	db:=DB_Test()
	r:=mux.NewRouter()
	repo := repositories.InitAccountRepoImpl(db)
	uc := usecases.InitAccountUsecase(repo,db)

	controller := &AccountHandler{accountUsecase: uc}

	acc := r.PathPrefix("/account/{accNum}").Subrouter()
	acc.HandleFunc("",controller.GetSaldo).Methods(http.MethodGet)
	acc.HandleFunc("/transfer",controller.Transfer).Methods(http.MethodPost)	

	return r
}

func TestGetSaldo(t *testing.T) {
	router := mockRouter()
	req, err := http.NewRequest("GET", "/account/555112", nil)
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	log.Println(response.Body)
	assert.Equal(t, 200, response.Code, "Response 200 is expected")
}
func TestGetSaldoFailure(t *testing.T) {
	router := mockRouter()
	req, err := http.NewRequest("GET", "/account/5551111", nil)
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 404, response.Code, "Response 404 is expected")
}


func TestAccController_Transfer(t *testing.T){
	router := mockRouter()
	body,_ := json.Marshal(mockTrans)
	req, err := http.NewRequest("POST", "/account/555112/transfer", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 201, response.Code, "Response 201 is expected")
}
func TestAccController_TransferFailUser(t *testing.T){
	router := mockRouter()
	body,_ := json.Marshal(mockTrans)
	req, err := http.NewRequest("POST", "/account/5551123/transfer", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("error occurred %v", err)
	}
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)
	assert.Equal(t, 400, response.Code, "Response 400 is expected")
}