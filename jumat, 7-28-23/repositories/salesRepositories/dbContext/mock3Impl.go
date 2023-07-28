package dbContext

import (
	"context"
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
)

type CreateMergeMock struct {
	CreateUsersParams
	CreateEducationParams
	CreateMediaParams
}

const createUsers = `-- name: CreateUsers :one

INSERT INTO users.users 
(user_first_name, user_last_name, user_birth_date, user_photo)
VALUES($1,$2,$3,$4)
RETURNING *
`

type CreateUsersParams struct {
	UserFirstName string       `db:"user_first_name" json:"userFirstName"`
	UserLastName  string       `db:"user_last_name" json:"userLastName"`
	UserBirthDate sql.NullTime `db:"user_birth_date" json:"userBirthDate"`
	UserPhoto     string       `db:"user_photo" json:"userPhoto"`
}

func (q *Queries) CreateUsersParams(ctx context.Context, arg CreateUsersParams) (*CreateUsersParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.UserFirstName,
		arg.UserLastName,
		arg.UserBirthDate,
		arg.UserPhoto,
	)
	i := CreateUsersParams{}
	err := row.Scan(
		&i.UserFirstName,
		&i.UserLastName,
		&i.UserBirthDate,
		&i.UserPhoto,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateUsersParams{
		UserFirstName: i.UserFirstName,
		UserLastName:  i.UserLastName,
		UserBirthDate: i.UserBirthDate,
		UserPhoto:     i.UserPhoto,
	}, nil
}

const createEducation = `-- name: CreateEducation :one

INSERT INTO users.users_education
(usdu_school, usdu_degree, usdu_field_study,usdu_description)
VALUES($1,$2,$3,$4)
RETURNING *
`

type CreateEducationParams struct {
	UsduSchool      string `db:"usdu_school" json:"usduSchool"`
	UsduDegree      string `db:"usdu_degree" json:"usduDegree"`
	UsduFieldStudy  string `db:"usdu_field_study" json:"usduFieldStudy"`
	UsduDescription string `db:"usdu_description" json:"usduDescription"`
}

func (q *Queries) CreateEducationParams(ctx context.Context, arg CreateEducationParams) (*CreateEducationParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createEducation,
		arg.UsduSchool,
		arg.UsduDegree,
		arg.UsduFieldStudy,
		arg.UsduDescription,
	)
	i := CreateEducationParams{}
	err := row.Scan(
		&i.UsduSchool,
		&i.UsduDegree,
		&i.UsduFieldStudy,
		&i.UsduDescription,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateEducationParams{
		UsduSchool:      i.UsduDegree,
		UsduDegree:      i.UsduDegree,
		UsduFieldStudy:  i.UsduFieldStudy,
		UsduDescription: i.UsduDescription,
	}, nil
}

const createMedia = `-- name: CreateMedia :one

INSERT INTO users.users_media
(usme_id, usme_entity_id, usme_file_link, usme_filename,
usme_filesize, usme_filetype, usme_note, usme_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING usme_id
`

type CreateMediaParams struct {
	UsmeFilename string `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize int32  `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype string `db:"usme_filetype" json:"usmeFiletype"`
}

func (q *Queries) CreateMediaParams(ctx context.Context, arg CreateMediaParams) (*CreateMediaParams, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createMedia,
		arg.UsmeFilename,
		arg.UsmeFilesize,
		arg.UsmeFiletype,
	)
	i := CreateMediaParams{}
	err := row.Scan(
		&i.UsmeFilename,
		&i.UsmeFilesize,
		&i.UsmeFiletype,
	)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &CreateMediaParams{
		UsmeFilename: arg.UsmeFilename,
		UsmeFilesize: arg.UsmeFilesize,
		UsmeFiletype: arg.UsmeFiletype,
	}, nil
}

const getUsers = `-- name: GetUsers :one

SELECT user_entity_id, user_name, user_password, user_first_name, user_last_name, user_birth_date, user_email_promotion, user_demographic, user_modified_date, user_photo, user_current_role FROM users.users
WHERE user_entity_id = $1
`

// users
func (q *Queries) GetUsers(ctx context.Context, userEntityID int32) (models.UsersUser, error) {
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
