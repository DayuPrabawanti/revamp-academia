package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"codeid.revampacademy/models"
)

type SignUpEmplloyeeParams struct {
	User  CreateUsersParams
	Email CreateEmailParams
	Phone CreatePhonesParams
}

const createEmployee = `-- name: CreateEmployee :one

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

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateUsersParams) (*models.UsersUser, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.UserName,
		arg.UserPassword,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserEmailPromotion,
		arg.UserDemographic,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UserPhoto,
		sql.NullInt64{Int64: 12, Valid: true},
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
		UserEntityID:       i.UserEntityID,
		UserName:           i.UserName,
		UserPassword:       i.UserPassword,
		UserFirstName:      i.UserFirstName,
		UserLastName:       i.UserLastName,
		UserBirthDate:      i.UserBirthDate,
		UserEmailPromotion: i.UserEmailPromotion,
		UserDemographic:    i.UserDemographic,
		UserModifiedDate:   i.UserModifiedDate,
		UserPhoto:          i.UserPhoto,
		UserCurrentRole:    i.UserCurrentRole,
	}, nil
}

const createEmailEmployee = `-- name: CreateEmailEmployee :one
WITH inserted_entity AS (
	SELECT * FROM users.users
	ORDER BY user_entity_id DESC
	LIMIT 1
  )
INSERT INTO users.users_email
(pmail_entity_id, pmail_address, pmail_modified_date)
SELECT user_entity_id,$1,$2 FROM inserted_entity
RETURNING pmail_entity_id, pmail_address, pmail_modified_date `

func (q *Queries) CreateEmailEmployee(ctx context.Context, arg CreateEmailParams) (*models.UsersUsersEmail, *models.ResponseError) {
	
	if !emailKantor(arg.PmailAddress.String, "@code.id") {
		return nil, &models.ResponseError{
			Message: "Email Domain Tidak Di Izinkan !",
			Status:  http.StatusBadRequest,
		}
	}
	
	row := q.db.QueryRowContext(ctx, createEmailEmployee,
		arg.PmailAddress,
		sql.NullTime{Time:time.Now(), Valid:true},
	)
	i := models.UsersUsersEmail{}
	err := row.Scan(
		&i.PmailEntityID,
		&i.PmailAddress,
		&i.PmailModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersEmail{
		PmailEntityID: i.PmailEntityID,
		PmailID:       i.PmailID,
		PmailAddress:  i.PmailAddress,
		PmailModifiedDate: i.PmailModifiedDate,
	}, nil
}

func emailKantor(email, emailKantor string) bool {
    return strings.HasSuffix(email, emailKantor)
}

const createPhonesEmployee = `-- name: CreatePhonesEmployee :one
WITH inserted_entity AS (
	SELECT * FROM users.users
	ORDER BY user_entity_id DESC
	LIMIT 1
)
INSERT INTO users.users_phones
(uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code)
SELECT user_entity_id, $1, $2, $3 FROM inserted_entity
RETURNING uspo_entity_id, uspo_number, uspo_modified_date, uspo_ponty_code
`


func (q *Queries) CreatePhonesEmployee(ctx context.Context, arg CreatePhonesParams) (*models.UsersUsersPhone, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createPhonesEmployee,
		arg.UspoNumber,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UspoPontyCode,
	)
	i := models.UsersUsersPhone{}
	err := row.Scan(
		&i.UspoEntityID,
		&i.UspoNumber,
		&i.UspoModifiedDate,
		&i.UspoPontyCode,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersPhone{
		UspoEntityID: i.UspoEntityID,
		UspoNumber: i.UspoNumber,
		UspoModifiedDate: i.UspoModifiedDate,
		UspoPontyCode: i.UspoPontyCode,
	}, nil
}