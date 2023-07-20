package dbContext

import (
	"context"
	"database/sql"
)

type CreateUsersParams struct {
	UserEntityID       int32          `db:"user_entity_id" json:"userEntityId"`
	UserName           string         `db:"user_name" json:"userName"`
	UserPassword       string         `db:"user_password" json:"userPassword"`
	UserFirstName      string         `db:"user_first_name" json:"userFirstName"`
	UserLastName       string         `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime   `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion int64          `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime   `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          string         `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    int64          `db:"user_current_role" json:"userCurrentRole"`
}

const createUsers = `-- name: CreateUsers :one

INSERT INTO users.users 
(user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
RETURNING user_entity_id
`

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.UserEntityID,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		arg.UserModifiedDate,
		arg.UserPhoto,
		arg.UserCurrentRole,
	)
	var user_entity_id int32
	err := row.Scan(&user_entity_id)
	return user_entity_id, err
}
