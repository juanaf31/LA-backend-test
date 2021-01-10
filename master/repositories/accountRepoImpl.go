package repositories

import (
	"database/sql"
	"linkAja/models"
	"linkAja/utils"
	"log"
)

type AccountRepoImpl struct{
	db *sql.DB
}

func InitAccountRepoImpl(db *sql.DB) AccountRepository {
	return &AccountRepoImpl{db: db}
}

func (rp *AccountRepoImpl)GetSaldo(accNum string)(*models.Account,error){

	row := rp.db.QueryRow(utils.CHECK_SALDO,accNum)
	
	var acc = models.Account{}

	err := row.Scan(&acc.AccountNumber,&acc.CustName,&acc.Balance)

	if err!=nil{
		return nil,err
	}

	return &acc,nil
}

func (rp *AccountRepoImpl)Transfer(accNum string,input *models.Transfer)(error){
	
	tx,err := rp.db.Begin()

	if err != nil {
		log.Println(err)
		return err
	}


	_,err = tx.Exec(utils.SENDTRANSFER,&input.Amount,accNum)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	_,err = tx.Exec(utils.RECEIVETRANSFER,&input.Amount,&input.Receiver)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	tx.Commit()

	return nil
}