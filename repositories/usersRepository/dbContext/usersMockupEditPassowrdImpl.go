package dbContext

import (
	"context"
	"time"
)

type EditUsersPasswordParams struct {
	UserEntityID     int32    `db:"user_entity_id" json:"userEntityId"`
	UserPassword     string   `db:"user_password" json:"userPassword"`
	UserModifiedDate time.Time `db:"user_modified_date" json:"userModifiedDate"`
}
type GetUsersPasswordParams struct {
	UserPassword       string        `db:"user_password" json:"userPassword"`
}

const GetPassword = `-- name: GetPassword :one

SELECT user_password
FROM users.users
WHERE user_entity_id = $1
`

// users
func (q *Queries) GetPassword(ctx context.Context, userEntityID int32) (GetUsersPasswordParams, error) {
	row := q.db.QueryRowContext(ctx, GetPassword, userEntityID)
	var i GetUsersPasswordParams
	err := row.Scan(
		&i.UserPassword,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users.users
  set user_password = $2,
  user_modified_date = Now()
WHERE user_entity_id = $1
`
func (q *Queries) UpdatePassword(ctx context.Context, userEntityId int32, arg EditUsersPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, userEntityId,
		arg.UserPassword,
		)
	return err
}