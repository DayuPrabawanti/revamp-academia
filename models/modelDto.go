package models

// import (
// 	"database/sql"
// 	"time"
// )

// type CreatePaymentUsersAccountDto struct {
// 	UsacBankEntityID  int32     `db:"usac_bank_entity_id" json:"usacBankEntityId"`
// 	UsacUserEntityID  int32     `db:"usac_user_entity_id" json:"usacUserEntityId"`
// 	UsacAccountNumber string    `db:"usac_account_number" json:"usacAccountNumber"`
// 	UsacSaldo         string    `db:"usac_saldo" json:"usacSaldo"`
// 	UsacType          string    `db:"usac_type" json:"usacType"`
// 	UsacStartDate     time.Time `db:"usac_start_date" json:"usacStartDate"`
// 	UsacEndDate       time.Time `db:"usac_end_date" json:"usacEndDate"`
// 	UsacModifiedDate  time.Time `db:"usac_modified_date" json:"usacModifiedDate"`
// 	UsacStatus        string    `db:"usac_status" json:"usacStatus"`
// }

// type CreatePaymentBankDto struct {
// 	BankEntityID     int32     `db:"bank_entity_id" json:"bankEntityId"`
// 	BankCode         string    `db:"bank_code" json:"bankCode"`
// 	BankName         string    `db:"bank_name" json:"bankName"`
// 	BankModifiedDate time.Time `db:"bank_modified_date" json:"bankModifiedDate"`
// }

// type CreatePaymentFintechDto struct {
// 	FintEntityID     int32     `db:"fint_entity_id" json:"fintEntityId"`
// 	FintCode         string    `db:"fint_code" json:"fintCode"`
// 	FintName         string    `db:"fint_name" json:"fintName"`
// 	FintModifiedDate time.Time `db:"fint_modified_date" json:"fintModifiedDate"`
// }

// type ListTopupDto []struct {
// 	SourceName    sql.NullString `json:"sourceName"`
// 	SourceAccount sql.NullString `json:"sourceAccount"`
// 	SourceSaldo   float64        `json:"sourceSaldo"`
// 	TargetName    sql.NullString `json:"targetName"`
// 	TargetAccount sql.NullString `json:"targetAccount"`
// 	TargetSaldo   float64        `json:"targetSaldo"`
// }

// type CreateCategoryDto struct {
// 	CategoryID   int16  `db:"category_id" json:"category_id"`
// 	CategoryName string `db:"category_name" json:"category_name"`
// 	Description  string `db:"description" json:"description,omitempty"`
// 	Picture      []byte `db:"picture" json:"picture,omitempty"`
// }

// type CreateCategoryProductDto struct {
// 	CreateCategoryDto
// 	Products []CreateProductDto
// }
