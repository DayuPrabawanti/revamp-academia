package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreatePhonesParams struct {
	UspoEntityID     int32          `db:"uspo_entity_id" json:"uspoEntityId"`
	UspoNumber       string         `db:"uspo_number" json:"uspoNumber"`
	UspoModifiedDate sql.NullTime   `db:"uspo_modified_date" json:"uspoModifiedDate"`
	UspoPontyCode    sql.NullString `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

const listPhones = `-- name: ListPhones :many
SELECT uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code FROM users.users_phones
ORDER BY uspo_entity_id
`

func (q *Queries) ListPhones(ctx context.Context) ([]models.UsersUsersPhone, error) {
	rows, err := q.db.QueryContext(ctx, listPhones)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersPhone
	for rows.Next() {
		var i models.UsersUsersPhone
		if err := rows.Scan(
			&i.UspoEntityID,
			&i.UspoNumber,
			&i.UspoModifiedDate,
			&i.UspoPontyCode,
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

const getPhones = `-- name: GetPhones :one

SELECT uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code FROM users.users_phones
WHERE uspo_entity_id = $1
`

// Users Phones
func (q *Queries) GetPhones(ctx context.Context, uspoEntityID int32) (models.UsersUsersPhone, error) {
	row := q.db.QueryRowContext(ctx, getPhones, uspoEntityID)
	var i models.UsersUsersPhone
	err := row.Scan(
		&i.UspoEntityID,
		&i.UspoNumber,
		&i.UspoModifiedDate,
		&i.UspoPontyCode,
	)
	return i, err
}

const createPhones = `-- name: CreatePhones :one
WITH inserted_entity AS (
	SELECT * FROM users.users
	ORDER BY user_entity_id DESC
	LIMIT 1
)
INSERT INTO users.users_phones
(uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code)
SELECT user_entity_id, $1, $2, $3 FROM inserted_entity
RETURNING uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code
`


func (q *Queries) CreatePhones(ctx context.Context, arg CreatePhonesParams) (*models.UsersUsersPhone, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPhones,
		arg.UspoNumber,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UspoPontyCode,
	)
	i := models.UsersUsersPhone{}
	err := row.Scan(
		&i.UspoEntityID,
		&i.UspoNumber,
		&i.UspoModifiedDate,
		&i.UspoPontyCode,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersPhone{
		UspoEntityID: i.UspoEntityID,
		UspoNumber: i.UspoNumber,
		UspoModifiedDate: i.UspoModifiedDate,
		UspoPontyCode: i.UspoPontyCode,
	}, nil
}

const updatePhones = `-- name: UpdatePhones :exec
UPDATE users.users_phones
  set uspo_number = $2,
  uspo_modified_date = $3,
  uspo_ponty_code=$4
WHERE uspo_entity_id = $1
`
func (q *Queries) UpdatePhones(ctx context.Context, arg CreatePhonesParams) error {
	_, err := q.db.ExecContext(ctx, updatePhones,
		arg.UspoEntityID,
		arg.UspoNumber,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UspoPontyCode,
	)
	return err
}


const deletePhones = `-- name: DeletePhones :exec
DELETE FROM users.users_phones
WHERE uspo_entity_id = $1
`

func (q *Queries) DeletePhones(ctx context.Context, uspoEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deletePhones, uspoEntityID)
	return err
}