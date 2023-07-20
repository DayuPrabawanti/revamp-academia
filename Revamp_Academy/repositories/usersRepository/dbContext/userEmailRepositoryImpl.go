package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	model "codeid.revampacademy/models/usersModel"
)

type CreateEmailParams struct {
	PmailEntityID     int32        `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32        `db:"pmail_id" json:"pmailId"`
	PmailAddress      string       `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate sql.NullTime `db:"pmail_modified_date" json:"pmailModifiedDate"`
}

const listEmail = `-- name: ListEmail :many
SELECT pmail_entity_id, pmail_id, pmail_address, pmail_modified_date FROM users.users_email
ORDER BY pmail_id
`

func (q *Queries) ListEmail(ctx context.Context) ([]model.UsersUsersEmail, error) {
	rows, err := q.db.QueryContext(ctx, listEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []model.UsersUsersEmail
	for rows.Next() {
		var i model.UsersUsersEmail
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

const getEmail = `-- name: GetEmail :one

SELECT pmail_entity_id, pmail_id, pmail_address, pmail_modified_date FROM users.users_email
WHERE pmail_id = $1
`

// Get User Email
func (q *Queries) GetEmail(ctx context.Context, pmailID int32) (model.UsersUsersEmail, error) {
	row := q.db.QueryRowContext(ctx, getEmail, pmailID)
	var i model.UsersUsersEmail
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)
	return i, err
}

// Create User Email

const createEmail = `-- name: CreateEmail :one

INSERT INTO users.users_email
(pmail_entity_id, pmail_id, pmail_address, pmail_modified_date)
VALUES($1,$2,$3,$4)
RETURNING * 
`

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (*model.UsersUsersEmail, *model.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmail,
		arg.PmailEntityID,
		arg.PmailID,
		arg.PmailAddress,
		arg.PmailModifiedDate,
	)
	i := model.UsersUsersEmail{}
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)

	if err != nil {
		return nil, &model.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &model.UsersUsersEmail{
		PmailEntityID:     i.PmailEntityID,
		PmailID:           i.PmailID,
		PmailAddress:      i.PmailAddress,
		PmailModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

// Mengambil sintaks Update pada file query.sql_gen.go
const updateEmail = `-- name: UpdateEmail :exec
UPDATE users.users_email
  set pmail_entity_id = $2,
  pmail_address   = $3,
  pmail_modified_date = $4
WHERE pmail_id = $1
`

func (q *Queries) UpdateEmail(ctx context.Context, arg CreateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail,
		arg.PmailID,
		arg.PmailEntityID,
		arg.PmailAddress,
		sql.NullTime{Time: time.Now(), Valid: true},
	)
	return err
}

// Mengambil sintaks Delete pada file query.sql_gen.go
const deleteEmail = `-- name: DeleteEmail :exec
DELETE FROM users.users_email
WHERE pmail_id = $1
`

func (q *Queries) DeleteEmail(ctx context.Context, pmailID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEmail, pmailID)
	return err
}
