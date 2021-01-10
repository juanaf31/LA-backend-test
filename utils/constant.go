package utils

const (
	CHECK_USER_FOUND = `select account_number from account where account_number=?`
	CHECK_SALDO = `select acc.account_number,cus.name as customer_name,acc.balance from account acc join customer cus on acc.customer_number = cus.customer_number where acc.account_number=?`
	SENDTRANSFER = `update account set balance = balance - ?  where account_number = ?`
	RECEIVETRANSFER = `update account set balance = balance + ? where account_number = ?`
)