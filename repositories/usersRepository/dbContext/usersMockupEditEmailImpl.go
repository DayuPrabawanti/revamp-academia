package dbContext

import (
	"context"
)

type AddEmailParams struct {
	PmailAddress string `db:"pmail_address" json:"pmailAddress"`
}

const addEmail = `-- name: addEmail :one

INSERT INTO users.users_email
(pmail_entity_id, pmail_address, pmail_modified_date)
VALUES ($1, $2, Now())
RETURNING pmail_entity_id, pmail_address, pmail_modified_date`

func (q *Queries) AddEmail(ctx context.Context, arg AddEmailParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addEmail, id,
	arg.PmailAddress,
	)
	return err
}


