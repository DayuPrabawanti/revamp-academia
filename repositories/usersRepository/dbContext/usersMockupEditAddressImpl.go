package dbContext

import "context"

type AddAddressParams struct {
	AddrLine1      string `db:"addr_line1" json:"addrLine1"`
	AddrLine2      string `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode string `db:"addr_postal_code" json:"addrPostalCode"`
	AddrCityID     int32  `db:"addr_city_id" json:"addrCityId"`
}

type AddressParams struct {
	EtadAddrID   int32 `db:"etad_addr_id" json:"etadAddrId"`
	EtadEntityID int32 `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID   int32 `db:"etad_adty_id" json:"etadAdtyId"`
}

type Address struct {
	Address AddAddressParams
	Addr 	AddressParams
}

const addAddress = `-- name: addAddress :one

INSERT INTO master.address
(addr_line1, addr_line2, addr_postal_code, addr_modified_date, addr_city_id)
VALUES ($1, $2, $3, Now(), $4)
RETURNING addr_line1, addr_line2, addr_postal_code, addr_city_id`

func (q *Queries) AddAddress(ctx context.Context, arg AddAddressParams) error {
	_, err := q.db.ExecContext(ctx, addAddress,
		arg.AddrLine1,
		arg.AddrLine2,
		arg.AddrPostalCode,
		arg.AddrCityID,
	)
	return err
}

const addAddress2 = `-- name: addAddress2 :one
WITH inserted_entity AS (
	SELECT * FROM master.address
	ORDER BY addr_id DESC
	LIMIT 1
)
INSERT INTO users.users_address
(etad_addr_id, etad_modified_date, etad_entity_id, etad_adty_id)
SELECT addr_id, Now(), $1, $2 FROM inserted_entity
`

func (q *Queries) AddAddress2(ctx context.Context, arg AddressParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addAddress2, id,
		arg.EtadAdtyID,
	)
	return err
}