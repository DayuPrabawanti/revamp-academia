package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models/usersModel"
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

func (q *Queries) ListEmail(ctx context.Context) ([]usersModel.UsersUsersEmail, error) {
	rows, err := q.db.QueryContext(ctx, listEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []usersModel.UsersUsersEmail
	for rows.Next() {
		var i usersModel.UsersUsersEmail
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

// Users Email
func (q *Queries) GetEmail(ctx context.Context, pmailID int32) (usersModel.UsersUsersEmail, error) {
	row := q.db.QueryRowContext(ctx, getEmail, pmailID)
	var i usersModel.UsersUsersEmail
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)
	return i, err
}

const createEmail = `-- name: CreateEmail :one

INSERT INTO users.users_email
(pmail_entity_id, pmail_id, pmail_address, pmail_modified_date)
VALUES($1,$2,$3,$4)
RETURNING * `

func (q *Queries) CreateEmail(ctx context.Context, arg CreateEmailParams) (*usersModel.UsersUsersEmail, *usersModel.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmail,
		arg.PmailEntityID,
		arg.PmailID,
		arg.PmailAddress,
		arg.PmailModifiedDate,
	)
	i := usersModel.UsersUsersEmail{}
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)

	if err != nil {
		return nil, &usersModel.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &usersModel.UsersUsersEmail{
		PmailEntityID: i.PmailEntityID,
		PmailID:       i.PmailID,
		PmailAddress:  i.PmailAddress,
		PmailModifiedDate: sql.NullTime{Time:time.Now(), Valid:true},
	}, nil
}