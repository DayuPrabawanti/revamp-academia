package dbContext

import (
	"context"
	"database/sql"

	"codeid.revampacademy/models"
)

type CreateMasterAddressParams struct {
	AddrID              int32          `db:"addr_id" json:"addrId"`
	AddrLine1           sql.NullString `db:"addr_line1" json:"addrLine1"`
	AddrLine2           sql.NullString `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode      sql.NullString `db:"addr_postal_code" json:"addrPostalCode"`
	AddrSpatialLocation sql.NullString `db:"addr_spatial_location" json:"addrSpatialLocation"`
	AddrModifiedDate    sql.NullTime   `db:"addr_modified_date" json:"addrModifiedDate"`
	AddrCityID          sql.NullInt32  `db:"addr_city_id" json:"addrCityId"`
}

const listMasterAddress = `-- name: ListMasterAddress :many
SELECT addr_id, addr_line1, addr_line2, addr_postal_code, addr_spatial_location, addr_modified_date, addr_city_id FROM master.address
ORDER BY addr_id
`

func (q *Queries) ListMasterAddress(ctx context.Context) ([]models.MasterAddress, error) {
	rows, err := q.db.QueryContext(ctx, listMasterAddress)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MasterAddress
	for rows.Next() {
		var i models.MasterAddress
		if err := rows.Scan(
			&i.AddrID,
			&i.AddrLine1,
			&i.AddrLine2,
			&i.AddrPostalCode,
			&i.AddrSpatialLocation,
			&i.AddrModifiedDate,
			&i.AddrCityID,
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