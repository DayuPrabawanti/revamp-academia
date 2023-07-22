package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateLicenseParams struct {
	UsliID           int32          `db:"usli_id" json:"usliId"`
	UsliLicenseCode  sql.NullString `db:"usli_license_code" json:"usliLicenseCode"`
	UsliModifiedDate sql.NullTime   `db:"usli_modified_date" json:"usliModifiedDate"`
	UsliStatus       string         `db:"usli_status" json:"usliStatus"`
	UsliEntityID     int32          `db:"usli_entity_id" json:"usliEntityId"`
}

// GetList Users License
const listLicense = `-- name: ListLicense :many
SELECT usli_id, usli_license_code, usli_modified_date, usli_status, usli_entity_id FROM users.users_license
ORDER BY usli_id
`

func (q *Queries) ListLicense(ctx context.Context) ([]models.UsersUsersLicense, error) {
	rows, err := q.db.QueryContext(ctx, listLicense)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersLicense
	for rows.Next() {
		var i models.UsersUsersLicense
		if err := rows.Scan(
			&i.UsliID,
			&i.UsliLicenseCode,
			&i.UsliModifiedDate,
			&i.UsliStatus,
			&i.UsliEntityID,
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

// Get License
const getLicense = `-- name: GetLicense :one

SELECT usli_id, usli_license_code, usli_modified_date, usli_status, usli_entity_id FROM users.users_license
WHERE usli_id = $1
`

// Get Users License
func (q *Queries) GetLicense(ctx context.Context, usliID int32) (models.UsersUsersLicense, error) {
	row := q.db.QueryRowContext(ctx, getLicense, usliID)
	var i models.UsersUsersLicense
	err := row.Scan(
		&i.UsliID,
		&i.UsliLicenseCode,
		&i.UsliModifiedDate,
		&i.UsliStatus,
		&i.UsliEntityID,
	)
	return i, err
}

// Create User License
const createLicense = `-- name: CreateLicense :one

INSERT INTO users.users_license
(usli_id, usli_license_code, usli_modified_date, usli_status, usli_entity_id)
VALUES($1,$2,$3,$4,$5)
RETURNING *
`

func (q *Queries) CreateLicense(ctx context.Context, arg CreateLicenseParams) (*models.UsersUsersLicense, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createLicense,
		arg.UsliID,
		arg.UsliLicenseCode,
		arg.UsliModifiedDate,
		arg.UsliStatus,
		arg.UsliEntityID,
	)
	i := models.UsersUsersLicense{}
	err := row.Scan(
		&i.UsliID,
		&i.UsliLicenseCode,
		&i.UsliModifiedDate,
		&i.UsliStatus,
		&i.UsliEntityID,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersLicense{
		UsliID:           i.UsliID,
		UsliLicenseCode:  i.UsliLicenseCode,
		UsliModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
		UsliStatus:       i.UsliStatus,
		UsliEntityID:     i.UsliEntityID,
	}, nil
}

// Update User License
const updateLicense = `-- name: UpdateLicense :exec
UPDATE users.users_license
  set usli_license_code = $2,
  usli_modified_date = $3,
  usli_status = $4,
  usli_entity_id = $5
WHERE usli_id = $1
`

func (q *Queries) UpdateLicense(ctx context.Context, arg CreateLicenseParams) error {
	_, err := q.db.ExecContext(ctx, updateLicense,
		arg.UsliID,
		arg.UsliLicenseCode,
		arg.UsliModifiedDate,
		arg.UsliStatus,
		arg.UsliEntityID,
	)
	return err
}

// Delete Table
const deleteLicense = `-- name: DeleteLicense :exec
DELETE FROM users.users_license
WHERE usli_id = $1
`

func (q *Queries) DeleteLicense(ctx context.Context, usliID int32) error {
	_, err := q.db.ExecContext(ctx, deleteLicense, usliID)
	return err
}
