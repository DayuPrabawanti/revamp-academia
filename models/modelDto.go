package models

import "time"

type CreatePaymentBankDto struct {
	BankEntityID     int32     `db:"bank_entity_id" json:"bankEntityId"`
	BankCode         string    `db:"bank_code" json:"bankCode"`
	BankName         string    `db:"bank_name" json:"bankName"`
	BankModifiedDate time.Time `db:"bank_modified_date" json:"bankModifiedDate"`
}

type CreatePaymentFintechDto struct {
	FintEntityID     int32     `db:"fint_entity_id" json:"fintEntityId"`
	FintCode         string    `db:"fint_code" json:"fintCode"`
	FintName         string    `db:"fint_name" json:"fintName"`
	FintModifiedDate time.Time `db:"fint_modified_date" json:"fintModifiedDate"`
}

type CreatePaymentTransactionPaymentDto struct {
	TrpaID           int32     `db:"trpa_id" json:"trpaId"`
	TrpaCodeNumber   string    `db:"trpa_code_number" json:"trpaCodeNumber"`
	TrpaOrderNumber  string    `db:"trpa_order_number" json:"trpaOrderNumber"`
	TrpaDebit        string    `db:"trpa_debit" json:"trpaDebit"`
	TrpaCredit       string    `db:"trpa_credit" json:"trpaCredit"`
	TrpaType         string    `db:"trpa_type" json:"trpaType"`
	TrpaNote         string    `db:"trpa_note" json:"trpaNote"`
	TrpaModifiedDate time.Time `db:"trpa_modified_date" json:"trpaModifiedDate"`
	TrpaFromID       string    `db:"trpa_source_id" json:"trpaSourceId"`
	TrpaToID         string    `db:"trpa_target_id" json:"trpaTargetId"`
	TrpaUserEntityID int32     `db:"trpa_user_entity_id" json:"trpaUserEntityId"`
}

type CreatePaymentUsersAccountDto struct {
	UsacBankEntityID  int32     `db:"usac_bank_entity_id" json:"usacBankEntityId"`
	UsacUserEntityID  int32     `db:"usac_user_entity_id" json:"usacUserEntityId"`
	UsacAccountNumber string    `db:"usac_account_number" json:"usacAccountNumber"`
	UsacSaldo         string    `db:"usac_saldo" json:"usacSaldo"`
	UsacType          string    `db:"usac_type" json:"usacType"`
	UsacStartDate     time.Time `db:"usac_start_date" json:"usacStartDate"`
	UsacEndDate       time.Time `db:"usac_end_date" json:"usacEndDate"`
	UsacModifiedDate  time.Time `db:"usac_modified_date" json:"usacModifiedDate"`
	UsacStatus        string    `db:"usac_status" json:"usacStatus"`
}

type TransferDto struct {
	AccountFrom struct {
		BankCode          string `db:"bank_code" json:"bankCode"`
		UsacAccountNumber string `db:"usac_account_number" json:"usacAccountNumber"`
		UsacSaldo         string `db:"usac_saldo" json:"usacSaldo"`
	} `json:"accountFrom"`
	AccountTo struct {
		FintCode          string `db:"fint_code" json:"fintCode"`
		UsacAccountNumber string `db:"usac_account_number" json:"usacAccountNumber"`
		UsacSaldo         string `db:"usac_saldo" json:"usacSaldo"`
	} `json:"accountTo"`
	Transaction struct {
		TrpaNote string `db:"trpa_note" json:"trpaNote"`
	} `json:"transaction"`
	Amount float64 `json:"amount"`
}
