package dbContext

import (
	"context"
	"time"
)

type EditUsernameParams struct {
	UserEntityID       int32  `db:"user_entity_id" json:"userEntityId"`
	UserName           string `db:"user_name" json:"userName"`
	UserFirstName      string `db:"user_first_name" json:"userFirstName"`
	UserLastName       string `db:"user_last_name" json:"userLastName"`
	UserBirthDate      time.Time   `db:"user_birth_date" json:"userBirthDate"`
	UserModifiedDate   time.Time  `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          string`db:"user_photo" json:"userPhoto"`
}

type GetUsernameParams struct {
	UserName           string `db:"user_name" json:"userName"`
	UserFirstName      string `db:"user_first_name" json:"userFirstName"`
	UserLastName       string `db:"user_last_name" json:"userLastName"`
	UserBirthDate      time.Time   `db:"user_birth_date" json:"userBirthDate"`
	UserPhoto          string`db:"user_photo" json:"userPhoto"`
}

const Getusername = `-- name: GetUsername :one

SELECT user_name, user_first_name, 
		user_last_name, user_birth_date, user_photo 
FROM users.users
WHERE user_entity_id = $1
`

// users
func (q *Queries) GetUsername(ctx context.Context, userEntityID int32) (GetUsernameParams, error) {
	row := q.db.QueryRowContext(ctx, Getusername, userEntityID)
	var i GetUsernameParams
	err := row.Scan(
		&i.UserName,
		&i.UserFirstName,
		&i.UserLastName,
		&i.UserBirthDate,
		&i.UserPhoto,
	)
	return i, err
}

const EditUsername = `-- name: EditUsername :exec
UPDATE users.users
  set user_name = $2,
  user_first_name= $3,
  user_last_name = $4,
  user_birth_date = $5,
  user_modified_date = Now(),
  user_photo = $6
WHERE user_entity_id = $1
`
func (q *Queries) EditUsername(ctx context.Context, arg EditUsernameParams) error {

	_, err := q.db.ExecContext(ctx, EditUsername, 
		arg.UserEntityID,
		arg.UserName,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserPhoto,
	)
	return err
}
