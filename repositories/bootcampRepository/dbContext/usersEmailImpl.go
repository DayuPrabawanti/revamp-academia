package dbContext

import (
	"context"

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
