package dbContext

import "context"

type AddSkillParams struct {
	UskiSktyName string `db:"uski_skty_name" json:"uskiSktyName"`
}

const addSkill = `-- name: addSkill :one

INSERT INTO users.users_skill(
	uski_entity_id, uski_modified_date, uski_skty_name
)
VALUES ($1, Now(), $2)
`

func (q *Queries) AddSkill(ctx context.Context, arg AddSkillParams, id int32) error {
	_, err := q.db.ExecContext(ctx, addSkill, id,
		arg.UskiSktyName,
	)
	return err
}
