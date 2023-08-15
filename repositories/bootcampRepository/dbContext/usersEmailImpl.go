package dbContext

import (
	"context"
	"database/sql"
	"time"

	"codeid.revampacademy/models"
)

type CreateEmailParams struct {
	PmailEntityID     int32  `db:"pmail_entity_id" json:"pmailEntityId"`
	PmailID           int32  `db:"pmail_id" json:"pmailId"`
	PmailAddress      string `db:"pmail_address" json:"pmailAddress"`
	PmailModifiedDate string `db:"pmail_modified_date" json:"pmailModifiedDate"`
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
		sql.NullTime{Time: time.Now(), Valid: true},
	)
	return err
}

const GetUSerEmail = `-- name:  GetUSerEmail :one
SELECT pmail_entity_id, pmail_id, pmail_address, pmail_modified_date FROM users.users_email
, picture FROM pmailentityid
WHERE pmail_id = $1
`

func (q *Queries) GetUSerEmail(ctx context.Context, pmailEntityId int32) (models.UsersUsersEmail, error) {
	row := q.db.QueryRowContext(ctx, GetUSerEmail, pmailEntityId)
	var i models.UsersUsersEmail
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)
	return i, err
}
