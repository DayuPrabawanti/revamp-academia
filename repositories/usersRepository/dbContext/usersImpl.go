package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateUsersParams struct {
	UserEntityID       int32                 `db:"user_entity_id" json:"userEntityId"`
	UserName           sql.NullString        `db:"user_name" json:"userName"`
	UserPassword       sql.NullString        `db:"user_password" json:"userPassword"`
	UserFirstName      sql.NullString        `db:"user_first_name" json:"userFirstName"`
	UserLastName       sql.NullString        `db:"user_last_name" json:"userLastName"`
	UserBirthDate      sql.NullTime          `db:"user_birth_date" json:"userBirthDate"`
	UserEmailPromotion sql.NullInt32         `db:"user_email_promotion" json:"userEmailPromotion"`
	UserDemographic    sql.NullString `db:"user_demographic" json:"userDemographic"`
	UserModifiedDate   sql.NullTime          `db:"user_modified_date" json:"userModifiedDate"`
	UserPhoto          sql.NullString        `db:"user_photo" json:"userPhoto"`
	UserCurrentRole    sql.NullInt32         `db:"user_current_role" json:"userCurrentRole"`
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

const getUsers = `-- name: GetUsers :one

SELECT user_entity_id, user_name, user_password, user_first_name, user_last_name, user_birth_date, user_email_promotion, user_demographic, user_modified_date, user_photo, user_current_role FROM users.users
WHERE user_entity_id = $1
`

// users
func (q *Queries) GetUser(ctx context.Context, userEntityID int32) (models.UsersUser, error) {
	row := q.db.QueryRowContext(ctx, getUsers, userEntityID)
	var i models.UsersUser
	err := row.Scan(
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
	)
	return i, err
}

const createUsers = `-- name: CreateUsers :one

WITH inserted_entity AS (
  INSERT INTO users.business_entity 
  (entity_modified_date)
  VALUES (Now())
  RETURNING entity_id
)
INSERT INTO users.users 
(user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role)
SELECT  entity_id, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10 FROM inserted_entity
RETURNING user_entity_id, user_name, user_password, user_first_name, 
user_last_name, user_birth_date, user_email_promotion, user_demographic, 
user_modified_date, user_photo, user_current_role
`

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		sql.NullTime{Time:time.Now(), Valid:true},
		arg.UserPhoto,
		sql.NullInt64{Int64: 1, Valid: true},
	)
	i := models.UsersUser{}
	err := row.Scan(
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
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUser{
		UserEntityID: i.UserEntityID,
		UserName:     i.UserName,
		UserPassword: i.UserPassword,
		UserFirstName:    i.UserFirstName,
		UserLastName:   i.UserLastName,
		UserBirthDate:      i.UserBirthDate,
		UserEmailPromotion:       i.UserEmailPromotion,
		UserDemographic:         i.UserDemographic,
		UserModifiedDate:        i.UserModifiedDate,
		UserPhoto:          i.UserPhoto,
		UserCurrentRole:           i.UserCurrentRole,
	}, nil
}

const updateUsers = `-- name: UpdateUsers :exec
UPDATE users.users
  set user_name = $2,
  user_password=$3,
  user_first_name= $4,
  user_last_name =$5,
  user_birth_date=$6,
  user_email_promotion=$7,
  user_demographic=$8,
  user_modified_date=$9,
  user_photo=$10,
  user_current_role=$11
WHERE user_entity_id = $1
`
func (q *Queries) UpdateUser(ctx context.Context, arg CreateUsersParams) error {
	_, err := q.db.ExecContext(ctx, updateUsers, 
		arg.UserEntityID,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		sql.NullTime{Time:time.Now(), Valid:true},
		arg.UserPhoto,
		arg.UserCurrentRole,)
	return err
}


const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users.users
WHERE user_entity_id = $1
`

func (q *Queries) DeleteUsers(ctx context.Context, userEntityID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUsers, userEntityID)
	return err
}

