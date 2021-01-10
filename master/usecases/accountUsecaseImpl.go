package usecases

import (
	"database/sql"
	"errors"
	"linkAja/master/repositories"
	"linkAja/models"
	"linkAja/utils"
	"log"
)

type AccountUsecaseImpl struct{
	accountRepo repositories.AccountRepository
	db *sql.DB
}

func InitAccountUsecase(accountRepo repositories.AccountRepository, db *sql.DB) AccountUsecase {
	return &AccountUsecaseImpl{
		accountRepo: accountRepo,
		db: db,
	}
}

func(u *AccountUsecaseImpl)GetSaldo(accNum string)(*models.Account,error){
	account,err := u.accountRepo.GetSaldo(accNum)
	if err!=nil{
		return nil,err
	}
	return account,nil
}

func (u *AccountUsecaseImpl)Transfer(accNum string,input *models.Transfer)error{
	row := u.db.QueryRow(utils.CHECK_SALDO,accNum)

	var acc = models.Account{}

	err := row.Scan(&acc.AccountNumber,&acc.CustName,&acc.Balance)

	if err != nil {
		log.Println(err)
		return err
	}

	if input.Amount > acc.Balance{
		return errors.New("Not enough balance")
	}

	rows, err := u.db.Query(utils.CHECK_USER_FOUND, input.Receiver)

    if err != nil {
		log.Println(err)
		return err
    }

    var RowsAffected int
    for rows.Next() {
        RowsAffected ++
	}
	
	if RowsAffected == 0 {
		return errors.New("destination account number not found")
	}

	err = u.accountRepo.Transfer(accNum,input)
	if err != nil{
		return err
	}
	return nil
}