package dbContext

import (
	"context"
	"database/sql"
	"net/http"

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

//create address

const createMasterAddress = `-- name: CreateMasterAddress :one
INSERT INTO master.address (addr_id, addr_line1, addr_line2, addr_postal_code, addr_spatial_location, addr_modified_date, addr_city_id)
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING addr_id
`

type CreateMasterAddressParams struct {
	AddrID              int32          `db:"addr_id" json:"addrId"`
	AddrLine1           string         `db:"addr_line1" json:"addrLine1"`
	AddrLine2           string         `db:"addr_line2" json:"addrLine2"`
	AddrPostalCode      string         `db:"addr_postal_code" json:"addrPostalCode"`
	AddrSpatialLocation sql.NullString `db:"addr_spatial_location" json:"addrSpatialLocation"`
	AddrModifiedDate    sql.NullTime   `db:"addr_modified_date" json:"addrModifiedDate"`
	AddrCityID          int32          `db:"addr_city_id" json:"addrCityId"`
}

func (q *Queries) CreateMasterAddress(ctx context.Context, arg CreateMasterAddressParams) (*models.MasterAddress, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createMasterAddress,
		arg.AddrID,
		arg.AddrLine1,
		arg.AddrLine2,
		arg.AddrPostalCode,
		arg.AddrSpatialLocation,
		arg.AddrModifiedDate,
		arg.AddrCityID,
	)
	i := models.MasterAddress{}
	err := row.Scan(
		&i.AddrID,
		&i.AddrLine1,
		&i.AddrLine2,
		&i.AddrPostalCode,
		&i.AddrSpatialLocation,
		&i.AddrModifiedDate,
		&i.AddrCityID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.MasterAddress{
		AddrID:              i.AddrID,
		AddrLine1:           i.AddrLine1,
		AddrLine2:           i.AddrLine2,
		AddrPostalCode:      i.AddrPostalCode,
		AddrSpatialLocation: i.AddrSpatialLocation,
		AddrModifiedDate:    i.AddrModifiedDate,
		AddrCityID:          i.AddrCityID,
	}, nil
}
