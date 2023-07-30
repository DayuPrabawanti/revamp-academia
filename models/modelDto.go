package models

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

// type ListCategoryProductDto []struct {
// 	CategoryID   int    `json:"category_id"`
// 	CategoryName string `json:"category_name"`
// 	Description  string `json:"description"`
// 	Picture      string `json:"picture"`
// 	Products     []struct {
// 		ProductID       int     `json:"product_id"`
// 		ProductName     string  `json:"product_name"`
// 		SupplierID      int     `json:"supplier_id"`
// 		CategoryID      int     `json:"category_id"`
// 		QuantityPerUnit string  `json:"quantity_per_unit"`
// 		UnitPrice       float64 `json:"unit_price"`
// 		UnitsInStock    int     `json:"units_in_stock"`
// 		UnitsOnOrder    int     `json:"units_on_order"`
// 		ReorderLevel    int     `json:"reorder_level"`
// 		Discontinued    int     `json:"discontinued"`
// 	} `json:"products"`
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
