package dbContext

import (
	"context"
	"database/sql"

	"codeid.revampacademy/models"
)

type CreateAddreesParams struct {
	EtadAddrID       int32         `db:"etad_addr_id" json:"etadAddrId"`
	EtadModifiedDate sql.NullTime  `db:"etad_modified_date" json:"etadModifiedDate"`
	EtadEntityID     sql.NullInt32 `db:"etad_entity_id" json:"etadEntityId"`
	EtadAdtyID       sql.NullInt32 `db:"etad_adty_id" json:"etadAdtyId"`
}

const listAddress = `-- name: ListAddress :many
SELECT etad_addr_id, etad_modified_date, etad_entity_id, etad_adty_id FROM users.users_address
ORDER BY etad_modified_date
`

func (q *Queries) ListAddress(ctx context.Context) ([]models.UsersUsersAddress, error) {
	rows, err := q.db.QueryContext(ctx, listAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersAddress
	for rows.Next() {
		var i models.UsersUsersAddress
		if err := rows.Scan(
			&i.EtadAddrID,
			&i.EtadModifiedDate,
			&i.EtadEntityID,
			&i.EtadAdtyID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}