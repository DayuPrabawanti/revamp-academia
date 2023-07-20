package dbContext

import (
	"context"
	"database/sql"

	"codeid.revampacademy/models"
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

const listUsers = `-- name: ListUsers :many
SELECT user_entity_id, user_name, user_password, user_first_name, user_last_name, user_birth_date, user_email_promotion, user_demographic, user_modified_date, user_photo, user_current_role FROM users.users
ORDER BY user_name
`

func (q *Queries) ListUsers(ctx context.Context) ([]models.UsersUser, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUser
	for rows.Next() {
		var i models.UsersUser
		if err := rows.Scan(
			&i.UserEntityID,
			&i.UserName,
			&i.UserPassword,
			&i.UserFirstName,
			&i.UserLastName,
			&i.UserBirthDate,
			&i.UserEmailPromotion,
			&i.UserDemographic,
			&i.UserModifiedDate,
			&i.UserPhoto,
			&i.UserCurrentRole,
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
