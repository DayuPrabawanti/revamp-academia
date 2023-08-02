package dbContext

import "context"

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

// type PaymentService struct {
// 	db *sql.DB
// }

// type TopupDetail struct {
// 	SourceName    sql.NullString
// 	SourceAccount sql.NullString
// 	SourceSaldo   float64
// 	TargetName    sql.NullString
// 	TargetAccount sql.NullString
// 	TargetSaldo   float64
// }

// const listTopupDetail = `-- name: ListTopupDetail :many

// SELECT
// 		b.bank_code,
// 		ub.usac_account_number AS bank_account_number,
// 		ub.usac_saldo AS bank_saldo,
// 		f.fint_code,
// 		uf.usac_account_number AS fintech_account_number,
// 		uf.usac_saldo AS fintech_saldo
// FROM
// 		users.users AS u
// LEFT JOIN
// 		payment.users_account AS ub
// 		ON
// 		u.user_entity_id = ub.usac_user_entity_id
// 		AND
// 		ub.usac_type = 'Bank'
// LEFT JOIN
// 		payment.bank AS b
// 		ON
// 		ub.usac_bank_entity_id = b.bank_entity_id
// LEFT JOIN
// 		payment.users_account AS uf
// 		ON
// 		u.user_entity_id = uf.usac_user_entity_id
// 		AND
// 		uf.usac_type = 'Fintech'
// LEFT JOIN
// 		payment.fintech AS f
// 		ON
// 		uf.usac_bank_entity_id = f.fint_entity_id
// WHERE
// 		b.bank_code IS NOT NULL
// 		AND
// 		f.fint_code IS NOT NULL
// GROUP BY
// 		u.user_entity_id,
// 		b.bank_code,
// 		ub.usac_account_number,
// 		ub.usac_saldo,
// 		f.fint_code,
// 		uf.usac_account_number,
// 		uf.usac_saldo;

// `

// func (q *Queries) ListTopupDetail(ctx context.Context) ([]TopupDetail, error) {
// 	rows, err := q.db.QueryContext(ctx, listTopupDetail)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []TopupDetail
// 	for rows.Next() {
// 		var i TopupDetail
// 		if err := rows.Scan(
// 			&i.SourceName,
// 			&i.SourceAccount,
// 			&i.SourceSaldo,
// 			&i.TargetName,
// 			&i.TargetAccount,
// 			&i.TargetSaldo,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// const getTopupDetailById = `-- name: ListTopupDetail :many

// SELECT
// 		b.bank_code,
// 		ub.usac_account_number AS bank_account_number,
// 		ub.usac_saldo AS bank_saldo,
// 		f.fint_code,
// 		uf.usac_account_number AS fintech_account_number,
// 		uf.usac_saldo AS fintech_saldo
// FROM
// 		users.users AS u
// LEFT JOIN
// 		payment.users_account AS ub
// 		ON
// 		u.user_entity_id = ub.usac_user_entity_id
// 		AND
// 		ub.usac_type = 'Bank'
// LEFT JOIN
// 		payment.bank AS b
// 		ON
// 		ub.usac_bank_entity_id = b.bank_entity_id
// LEFT JOIN
// 		payment.users_account AS uf
// 		ON
// 		u.user_entity_id = uf.usac_user_entity_id
// 		AND uf.usac_type = 'Fintech'
// LEFT JOIN
// 		payment.fintech AS f
// 		ON
// 		uf.usac_bank_entity_id = f.fint_entity_id
// WHERE
// 		u.user_entity_id = $1
// 		AND
// 		b.bank_code IS NOT NULL
// 		AND
// 		f.fint_code IS NOT NULL
// GROUP BY
// 		u.user_entity_id,
// 		b.bank_code,
// 		ub.usac_account_number,
// 		ub.usac_saldo,
// 		f.fint_code,
// 		uf.usac_account_number,
// 		uf.usac_saldo;
// `

// func (q *Queries) GetTopupDetailById(ctx context.Context, id int32) ([]TopupDetail, error) {
// 	rows, err := q.db.QueryContext(ctx, getTopupDetailById, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var items []TopupDetail
// 	for rows.Next() {
// 		var i TopupDetail
// 		if err := rows.Scan(
// 			&i.SourceName,
// 			&i.SourceAccount,
// 			&i.SourceSaldo,
// 			&i.TargetName,
// 			&i.TargetAccount,
// 			&i.TargetSaldo,
// 		); err != nil {
// 			return nil, err
// 		}
// 		items = append(items, i)
// 	}
// 	if err := rows.Close(); err != nil {
// 		return nil, err
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return items, nil
// }

// // Merekam Transaksi yang tersimpan di aplikasi
// type RecordTransaction struct {
// 	TrpaType         string
// 	TrpaDebit        float64
// 	TrpaCredit       float64
// 	TrpaNote         string
// 	TrpaFromID       string
// 	TrpaToID         string
// 	TrpaUserEntityID int
// }

// const transferFunds = `-- name: TransferFunds :exec
// 	UPDATE payment.users_account
// 	SET usac_saldo = usac_saldo - $2
// 	WHERE usac_account_number = $1;
// 	UPDATE payment.users_account
// 	SET usac_saldo = usac_saldo + $2
// 	WHERE usac_account_number = $3;
// `

// func (ps *PaymentService) TransferFunds(ctx context.Context, fromAccount string, amount float64, toAccount string, fromUserID int, toUserID int) error {
// 	tx, err := ps.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	// debit dan credit
// 	if _, err = tx.ExecContext(ctx, transferFunds, fromAccount, amount, toAccount); err != nil {
// 		tx.Rollback()
// 		return err
// 	}

// 	// menyimpan transaksi debit
// 	debitTransaction := &RecordTransaction{
// 		TrpaType:         "TR",
// 		TrpaDebit:        amount,
// 		TrpaCredit:       0,
// 		TrpaNote:         "Transfer to " + toAccount,
// 		TrpaFromID:       fromAccount,
// 		TrpaToID:         toAccount,
// 		TrpaUserEntityID: fromUserID,
// 	}
// 	if err = ps.CreateTransactionPayment(ctx, tx, *debitTransaction); err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	// menyimpan transaksi credit
// 	creditTransaction := &RecordTransaction{
// 		TrpaType:         "TR",
// 		TrpaDebit:        0,
// 		TrpaCredit:       amount,
// 		TrpaNote:         "Received transfer from  " + fromAccount,
// 		TrpaFromID:       fromAccount,
// 		TrpaToID:         toAccount,
// 		TrpaUserEntityID: fromUserID,
// 	}
// 	if err = ps.CreateTransactionPayment(ctx, tx, *creditTransaction); err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	// jika transaksi & debit berhasil semua, maka commit transaksi
// 	if err = tx.Commit(); err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return nil
// }

// func (q *PaymentService) CreateTransactionPayment(ctx context.Context, tx *sql.Tx, transaction RecordTransaction) error {
// 	const insertTransaction = `
// 		INSERT INTO payment.transactions (trpa_type, trpa_debit, trpa_credit, trpa_note, trpa_from_id, trpa_to_id, trpa_user_entity_id)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7);
// 	`
// 	_, err := tx.ExecContext(ctx, insertTransaction,
// 		transaction.TrpaType,
// 		transaction.TrpaDebit,
// 		transaction.TrpaCredit,
// 		transaction.TrpaNote,
// 		transaction.TrpaFromID,
// 		transaction.TrpaToID,
// 		transaction.TrpaUserEntityID,
// 	)
// 	return err
// }

// func (ps *PaymentService) Topup(ctx context.Context, sourceBankEntityID int32, targetFintechEntityID int32, amount float64) error {
// 	tx, err := ps.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	// check source balance
// 	var sourceBalance float64
// 	err = tx.QueryRowContext(ctx, `SELECT usac_saldo FROM payment.users_account WHERE usac_bank_entity_id = ?`, sourceBankEntityID).Scan(&sourceBalance)
// 	if err != nil {
// 		return err
// 	}

// 	if sourceBalance < amount {
// 		return errors.New("insufficient funds")
// 	}

// 	// deduct amount from source account
// 	_, err = tx.ExecContext(ctx, `UPDATE payment.users_account SET usac_saldo = usac_saldo - ? WHERE usac_bank_entity_id = ?`, amount, sourceBankEntityID)
// 	if err != nil {
// 		return err
// 	}

// 	// add amount to target account
// 	_, err = tx.ExecContext(ctx, `UPDATE payment.users_account SET usac_saldo = usac_saldo + ? WHERE usac_bank_entity_id = ?`, amount, targetFintechEntityID)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
