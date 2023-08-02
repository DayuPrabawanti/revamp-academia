package dbContext

import (
	"context"
)

const getBankAccount = `-- name: GetBankAccount :one
SELECT 
    b.bank_code, 
    ua.usac_account_number, 
    ua.usac_saldo
FROM 
    payment.bank AS b
JOIN 
    payment.users_account AS ua
ON 
    b.bank_entity_id = ua.usac_bank_entity_id
WHERE 
    ua.usac_type = 'Bank' AND
    b.bank_code = $1 AND
    ua.usac_account_number = $2
ORDER BY 
    b.bank_code, 
    ua.usac_account_number;
`

type BankAccount struct {
	BankCode          string
	BankAccountNumber string
	BankSaldo         float64
}

func (q *Queries) GetBankAccount(ctx context.Context, bankCode string, usacAccountNumber string) (BankAccount, error) {
	row := q.db.QueryRowContext(ctx, getBankAccount, bankCode, usacAccountNumber)
	var i BankAccount
	err := row.Scan(
		&i.BankCode,
		&i.BankAccountNumber,
		&i.BankSaldo,
	)
	return i, err
}

const getFintechAccount = `-- name: GetFintechAccount :one
SELECT 
    f.fint_code,
    ua.usac_account_number,
    ua.usac_saldo
FROM 
    payment.fintech f
JOIN 
    payment.users_account ua ON f.fint_entity_id = ua.usac_bank_entity_id
WHERE
    ua.usac_type = 'Fintech' AND
    f.fint_code = $1 AND
    ua.usac_account_number = $2
ORDER BY 
    f.fint_code, 
    ua.usac_account_number;
`

type FintechAccount struct {
	FintCode          string
	FintAccountNumber string
	FintSaldo         float64
}

func (q *Queries) GetFintechAccount(ctx context.Context, fintCode string, usacAccountNumber string) (FintechAccount, error) {
	row := q.db.QueryRowContext(ctx, getFintechAccount, fintCode, usacAccountNumber)
	var i FintechAccount
	err := row.Scan(
		&i.FintCode,
		&i.FintAccountNumber,
		&i.FintSaldo,
	)
	return i, err
}
