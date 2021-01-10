package usecases

import (
	"database/sql"
	"fmt"
	"linkAja/models"
	"linkAja/utils"
	"log"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountUsecaseMock struct{
	mock.Mock
}

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

func (m *AccountUsecaseMock) GetSaldo(id string) (*models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Account), args.Error(1)
}

func (m *AccountUsecaseMock) Transfer(id string, input *models.Transfer) error {
	args := m.Called(id, input)


	return args.Error(0)
}

func TestGetSaldoSuccess (t *testing.T){

	db,_ := NewMock()

	uc := new(AccountUsecaseMock)

	acc := models.Account{AccountNumber: "123", CustName:"juan",Balance: 1000}

	uc.On("GetSaldo","123").Return(&acc,nil)
	
	testInitAccountUsecase := InitAccountUsecase(uc,db)

	result,_ := testInitAccountUsecase.GetSaldo("123")

	uc.AssertExpectations(t)

	assert.Equal(t,"123",result.AccountNumber)
	assert.Equal(t,"juan",result.CustName)
	assert.Equal(t,1000,result.Balance)
}
func TestGetSaldoFailure (t *testing.T){
	db,_ := NewMock()

	uc := new(AccountUsecaseMock)

	acc := models.Account{AccountNumber: "123", CustName:"juan",Balance: 1000}

	uc.On("GetSaldo","123").Return(&acc, fmt.Errorf("some error"))
	
	testInitAccountUsecase := InitAccountUsecase(uc,db)

	_, err := testInitAccountUsecase.GetSaldo("123")

	uc.AssertExpectations(t)

	if assert.NotNil(t, err.Error()){
		assert.Equal(t,"some error",err.Error())
	}
}

func TestTransferSuccess (t *testing.T){
	db,mock := NewMock()
	uc := new(AccountUsecaseMock)
	repo := &AccountUsecaseImpl{
		// accountRepo: uc,
		db: db}

	defer func ()  {
		repo.db.Close()
	}()

	e := strconv.Itoa(tf.Amount)

	acc := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
	found := sqlmock.NewRows([]string{"account_number"}).AddRow(f.AccountNumber)
	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnRows(acc)
	mock.ExpectQuery(utils.CHECK_USER_FOUND).WithArgs(tf.Receiver).WillReturnRows(found)

	post := models.Transfer{Receiver: "123",Amount: 1000}

	uc.On("Transfer","123",&post).Return(nil)

	testInitAccountUsecase := InitAccountUsecase(uc,db)

	err := testInitAccountUsecase.Transfer("123",&post)

	uc.AssertExpectations(t)

	assert.Equal(t,err,nil)	

}

func TestTransferFailure(t *testing.T){
	db,mock := NewMock()
	uc := new(AccountUsecaseMock)
	repo := &AccountUsecaseImpl{
		// accountRepo: uc,
		db: db}

	defer func ()  {
		repo.db.Close()
	}()

	// e := strconv.Itoa(tf.Amount)

	// acc := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
	// found := sqlmock.NewRows([]string{"account_number"}).AddRow(f.AccountNumber)
	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnError(fmt.Errorf("some error"))
	// mock.ExpectQuery(utils.CHECK_USER_FOUND).WithArgs(tf.Receiver).WillReturnRows(found)

	post := models.Transfer{Receiver: "123",Amount: 1000}

	uc.On("Transfer","123",&post).Return(nil)

	testInitAccountUsecase := InitAccountUsecase(uc,db)

	err := testInitAccountUsecase.Transfer("123",&post)

	if assert.NotNil(t,err){
		assert.Equal(t,"some error",err.Error())
	}
}
func TestTransferFailureNoFound(t *testing.T){
	db,mock := NewMock()
	uc := new(AccountUsecaseMock)
	repo := &AccountUsecaseImpl{
		// accountRepo: uc,
		db: db}

	defer func ()  {
		repo.db.Close()
	}()

	e := strconv.Itoa(tf.Amount)

	acc := sqlmock.NewRows([]string{"account_number","customer_name","balance"}).AddRow(u.AccountNumber,u.CustName,e)
	// found := sqlmock.NewRows([]string{"account_number"}).AddRow(f.AccountNumber)
	mock.ExpectQuery(utils.CHECK_SALDO).WithArgs(u.AccountNumber).WillReturnRows(acc)
	mock.ExpectQuery(utils.CHECK_USER_FOUND).WithArgs(tf.Receiver).WillReturnError(fmt.Errorf("some error"))

	post := models.Transfer{Receiver: "123",Amount: 1000}

	uc.On("Transfer","123",&post).Return(nil)

	testInitAccountUsecase := InitAccountUsecase(uc,db)

	err := testInitAccountUsecase.Transfer("123",&post)

	if assert.NotNil(t,err){
		assert.Equal(t,"some error",err.Error())
	}
}