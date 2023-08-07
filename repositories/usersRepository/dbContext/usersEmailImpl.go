package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateEmailParams struct {
	PmailEntityID     int32        `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32        `db:"pmail_id" json:"pmailId"`
	PmailAddress      sql.NullString       `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate sql.NullTime `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

type Email struct {
	PmailAddress      string       `db:"pmail_address" json:"pmailAddress"`
}




const listEmail = `-- name: ListEmail :many
SELECT pmail_entity_id, pmail_id, pmail_address, pmail_modified_date FROM users.users_email
ORDER BY pmail_id
`

func (q *Queries) ListEmail(ctx context.Context) ([]models.UsersUsersEmail, error) {
	rows, err := q.db.QueryContext(ctx, listEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersEmail
	for rows.Next() {
		var i models.UsersUsersEmail
		if err := rows.Scan(
			&i.PmailEntityID,
			&i.PmailID,
			&i.PmailAddress,
			&i.PmailModifiedDate,
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

// const getEmail = `-- name: GetEmail :one

// SELECT pmail_entity_id, pmail_id, pmail_address, pmail_modified_date FROM users.users_email
// WHERE pmail_entity_id = $1
// `
const getEmail = `-- name: GetEmail :one

SELECT pmail_address FROM users.users_email
WHERE pmail_entity_id = $1
`

// Users Email
func (q *Queries) GetEmail(ctx context.Context, pmailID int32) ([]Email, error) {
	row, err := q.db.QueryContext(ctx, getEmail, pmailID)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	var email []Email
	for row.Next() {
		var i Email
		if err := row.Scan(
			&i.PmailAddress,
		); err != nil {
			return nil, err
		}
		email = append(email, i)
	}
	if err := row.Close(); err != nil {
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return email, nil
}


const createEmail = `-- name: CreateEmail :one
WITH inserted_entity AS (
	SELECT * FROM users.users
	ORDER BY user_entity_id DESC
	LIMIT 1
  )
INSERT INTO users.users_email
(pmail_entity_id, pmail_address, pmail_modified_date)
SELECT user_entity_id,$1,$2 FROM inserted_entity
RETURNING pmail_entity_id, pmail_address, pmail_modified_date `

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (*models.UsersUsersEmail, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmail,
		arg.PmailAddress,
		sql.NullTime{Time:time.Now(), Valid:true},
	)
	i := models.UsersUsersEmail{}
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersEmail{
		PmailEntityID: i.PmailEntityID,
		PmailID:       i.PmailID,
		PmailAddress:  i.PmailAddress,
		PmailModifiedDate: i.PmailModifiedDate,
	}, nil
}

const updateEmail = `-- name: UpdateEmail :exec
UPDATE users.users_email
  set pmail_entity_id = $2,
  pmail_address   = $3,
  pmail_modified_date = $4
WHERE pmail_id = $1`

func (q *Queries) UpdateEmail(ctx context.Context, arg CreateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail,
		arg.PmailID,
		arg.PmailEntityID,
		arg.PmailAddress,
		arg.PmailModifiedDate == sql.NullTime{Time:time.Now(), Valid:true},
	)
	return err
}

const deleteEmail = `-- name: DeleteEmail :exec
DELETE FROM users.users_email
WHERE pmail_id = $1
`

func (q *Queries) DeleteEmail(ctx context.Context, pmailID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmail, pmailID)
	return err
}