package dbContext

import "context"

type AddPhonesParams struct {
	UspoNumber    string `db:"uspo_number" json:"uspoNumber"`
	UspoPontyCode string `db:"uspo_ponty_code" json:"uspoPontyCode"`
}

const addPhones = `-- name: addPhones :one

INSERT INTO users.users_phones
(uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code)
VALUES ($1, $2, Now(), $3)
RETURNING uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code`

func (q *Queries) AddPhone(ctx context.Context, arg AddPhonesParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addPhones, id,
		arg.UspoNumber,
		arg.UspoPontyCode,
	)
	return err
}
