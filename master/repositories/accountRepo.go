package repositories

import "linkAja/models"

type AccountRepository interface{
	GetSaldo(string)(*models.Account,error)
	Transfer(string,*models.Transfer)(error)
}