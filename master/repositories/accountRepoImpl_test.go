package repositories

import (
	"database/sql"
	"fmt"
	"linkAja/models"
	"linkAja/utils"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = &models.Account{
	AccountNumber : "123",
	CustName : "Juan",
	Balance : 1000,
}

var f = &models.Found{
	AccountNumber: "123",
}

var tf = &models.Transfer{
	Receiver: "123",
	Amount: 1000,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetSaldoSuccess(t *testing.T){
	db,mock := NewMock()
	repo := &AccountRepoImpl{db: db}
	defer func ()  {
		repo.db.Close()
	}()
	e := strconv.Itoa(u.Balance)

	row := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
	
	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnRows(row)

	data,err:=repo.GetSaldo(u.AccountNumber)
	assert.NotNil(t,data)
	assert.NoError(t,err)
}

func TestGetSaldoFailure(t *testing.T){
	db,mock := NewMock()
	repo := &AccountRepoImpl{db: db}
	defer func ()  {
		repo.db.Close()
	}()

	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnError(fmt.Errorf("some error"))
	_,err := repo.GetSaldo(u.AccountNumber)

	if assert.NotNil(t,err){
		assert.Equal(t,"some error",err.Error())
	}
}

func TestTransferSuccess(t *testing.T){
	db,mock := NewMock()
	repo := &AccountRepoImpl{db: db}

	defer func ()  {
		repo.db.Close()
	}()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).WithArgs(tf.Amount,u.AccountNumber).WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectExec(regexp.QuoteMeta(utils.RECEIVETRANSFER)).WithArgs(tf.Amount,tf.Receiver).WillReturnResult(sqlmock.NewResult(1,1))
	
	mock.ExpectCommit()

	err := repo.Transfer(u.AccountNumber,tf)
	assert.NoError(t,err)
}



// func TestDestinationNotFound(t *testing.T){
// 	db,mock := NewMock()
// 	repo := &AccountRepoImpl{db: db}
// 	defer func ()  {
// 		repo.db.Close()
// 	}()
// 	e := strconv.Itoa(tf.Amount)
// 	acc := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
// 	// found := sqlmock.NewRows([]string{"account_number"}).AddRow(f.AccountNumber)
// 	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnRows(acc)

// 	mock.ExpectQuery(utils.CHECK_USER_FOUND).WithArgs(tf.Receiver).WillReturnError(fmt.Errorf("some error"))
// 	err := repo.Transfer(u.AccountNumber,tf)

// 	if assert.NotNil(t,err){
// 		assert.Equal(t,"some error",err.Error())
// 	}
// }

// func TestTransferDBBeginFailure(t *testing.T){
// 	db,mock := NewMock()
// 	repo := &AccountRepoImpl{db: db}

// 	defer func ()  {
// 		repo.db.Close()
// 	}()

// 	e := strconv.Itoa(tf.Amount)

// 	acc := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
// 	found := sqlmock.NewRows([]string{"account_number"}).AddRow(f.AccountNumber)
// 	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnRows(acc)
// 	mock.ExpectQuery(utils.CHECK_USER_FOUND).WithArgs(tf.Receiver).WillReturnRows(found)


// 	mock.ExpectBegin().WillReturnError(fmt.Errorf("some error"))

// 	err:=repo.Transfer(u.AccountNumber,tf)

// 	if assert.NotNil(t, err){
// 		assert.Equal(t,"some error",err.Error())
// 	}
// }

func TestTransferSendFailure(t *testing.T){
	db,mock := NewMock()
	repo := &AccountRepoImpl{db: db}

	defer func ()  {
		repo.db.Close()
	}()
	
	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).WithArgs(tf.Amount,u.AccountNumber).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	err := repo.Transfer(u.AccountNumber,tf)

	if assert.NotNil(t,err){
		assert.Equal(t,"some error",err.Error())
	}

}
func TestTransferReceiveFailure(t *testing.T){
	db,mock := NewMock()
	repo := &AccountRepoImpl{db: db}

	defer func ()  {
		repo.db.Close()
	}()

	mock.ExpectBegin()

	mock.ExpectExec(regexp.QuoteMeta(utils.SENDTRANSFER)).WithArgs(tf.Amount,u.AccountNumber).WillReturnResult(sqlmock.NewResult(1,1))
	mock.ExpectExec(regexp.QuoteMeta(utils.RECEIVETRANSFER)).WithArgs(tf.Amount,u.AccountNumber).WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback()

	err := repo.Transfer(u.AccountNumber,tf)

	if assert.NotNil(t,err){
		assert.Equal(t,"some error",err.Error())
	}
}

