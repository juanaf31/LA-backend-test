package master

import (
	"database/sql"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func DB_Test() *sql.DB{
	db,_ := sql.Open("sqlite3",":memory:")
	return db
}

func mockRouter() *mux.Router{
	r:= mux.NewRouter()
	return r
}

func TestInitData(t *testing.T){
	db := DB_Test()
	r := mockRouter()

	InitData(r,db)
}