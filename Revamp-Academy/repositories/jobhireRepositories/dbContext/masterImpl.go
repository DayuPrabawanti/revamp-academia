package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

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

const listMasterCity = `-- name: ListMasterCity :many
SELECT city_id, city_name, city_modified_date, city_prov_id FROM master.city
ORDER BY city_id
`

func (q *Queries) ListMasterCity(ctx context.Context) ([]models.MasterCity, error) {
	rows, err := q.db.QueryContext(ctx, listMasterCity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.MasterCity
	for rows.Next() {
		var i models.MasterCity
		if err := rows.Scan(
			&i.CityID,
			&i.CityName,
			&i.CityModifiedDate,
			&i.CityProvID,
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
