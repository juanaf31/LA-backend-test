package usecases

import "linkAja/models"

type AccountUsecase interface{
	GetSaldo(string)(*models.Account,error)
	Transfer(string,*models.Transfer)(error)
}