package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateMediaParams struct {
	UsmeID           int32          `db:"usme_id" json:"usmeId"`
	UsmeEntityID     int32          `db:"usme_entity_id" json:"usmeEntityId"`
	UsmeFileLink     sql.NullString `db:"usme_file_link" json:"usmeFileLink"`
	UsmeFilename     sql.NullString `db:"usme_filename" json:"usmeFilename"`
	UsmeFilesize     int32          `db:"usme_filesize" json:"usmeFilesize"`
	UsmeFiletype     sql.NullString `db:"usme_filetype" json:"usmeFiletype"`
	UsmeNote         sql.NullString `db:"usme_note" json:"usmeNote"`
	UsmeModifiedDate sql.NullTime   `db:"usme_modified_date" json:"usmeModifiedDate"`
}

// GetList
const listMedia = `-- name: ListMedia :many
SELECT usme_id, usme_entity_id, usme_file_link, usme_filename, usme_filesize, usme_filetype, usme_note, usme_modified_date FROM users.users_media
ORDER BY usme_id
`

func (q *Queries) ListMedia(ctx context.Context) ([]models.UsersUsersMedia, error) {
	rows, err := q.db.QueryContext(ctx, listMedia)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersMedia
	for rows.Next() {
		var i models.UsersUsersMedia
		if err := rows.Scan(
			&i.UsmeID,
			&i.UsmeEntityID,
			&i.UsmeFileLink,
			&i.UsmeFilename,
			&i.UsmeFilesize,
			&i.UsmeFiletype,
			&i.UsmeNote,
			&i.UsmeModifiedDate,
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

// GetMedia
const getMedia = `-- name: GetMedia :one

SELECT usme_id, usme_entity_id, usme_file_link, usme_filename, usme_filesize, usme_filetype, usme_note, usme_modified_date FROM users.users_media
WHERE usme_id = $1
`

// Users Media
func (q *Queries) GetMedia(ctx context.Context, usmeID int32) (models.UsersUsersMedia, error) {
	row := q.db.QueryRowContext(ctx, getMedia, usmeID)
	var i models.UsersUsersMedia
	err := row.Scan(
		&i.UsmeID,
		&i.UsmeEntityID,
		&i.UsmeFileLink,
		&i.UsmeFilename,
		&i.UsmeFilesize,
		&i.UsmeFiletype,
		&i.UsmeNote,
		&i.UsmeModifiedDate,
	)
	return i, err
}

// Create User Media
const createMedia = `-- name: CreateMedia :one

INSERT INTO users.users_media
(usme_id, usme_entity_id, usme_file_link, usme_filename,
usme_filesize, usme_filetype, usme_note, usme_modified_date)
VALUES($1,$2,$3,$4,$5,$6,$7,$8)
RETURNING *
`

func (q *Queries) CreateMedia(ctx context.Context, arg CreateMediaParams) (*models.UsersUsersMedia, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createMedia,
		arg.UsmeID,
		arg.UsmeEntityID,
		arg.UsmeFileLink,
		arg.UsmeFilename,
		arg.UsmeFilesize,
		arg.UsmeFiletype,
		arg.UsmeNote,
		arg.UsmeModifiedDate,
	)
	i := models.UsersUsersMedia{}
	err := row.Scan(
		&i.UsmeID,
		&i.UsmeEntityID,
		&i.UsmeFileLink,
		&i.UsmeFilename,
		&i.UsmeFilesize,
		&i.UsmeFiletype,
		&i.UsmeNote,
		&i.UsmeModifiedDate,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersMedia{
		UsmeID:           i.UsmeID,
		UsmeEntityID:     i.UsmeEntityID,
		UsmeFileLink:     i.UsmeFileLink,
		UsmeFilename:     i.UsmeFilename,
		UsmeFilesize:     i.UsmeFilesize,
		UsmeFiletype:     i.UsmeFiletype,
		UsmeNote:         i.UsmeNote,
		UsmeModifiedDate: sql.NullTime{Time: time.Now(), Valid: true},
	}, nil
}

// Update Media
const updateMedia = `-- name: UpdateMedia :exec
UPDATE users.users_media
  set usme_entity_id = $2,
  usme_file_link = $3,
  usme_filename = $4,
  usme_filesize = $5,
  usme_filetype = $6,
  usme_note = $7,
  usme_modified_date = $8
WHERE usme_id = $1
`

func (q *Queries) UpdateMedia(ctx context.Context, arg CreateMediaParams) error {
	_, err := q.db.ExecContext(ctx, updateMedia,
		arg.UsmeID,
		arg.UsmeEntityID,
		arg.UsmeFileLink,
		arg.UsmeFilename,
		arg.UsmeFilesize,
		arg.UsmeFiletype,
		arg.UsmeNote,
		arg.UsmeModifiedDate,
	)
	return err
}

// Delete Users Media
const deleteMedia = `-- name: DeleteMedia :exec
DELETE FROM users.users_media
WHERE usme_id = $1
`

func (q *Queries) DeleteMedia(ctx context.Context, usmeID int32) error {
	_, err := q.db.ExecContext(ctx, deleteMedia, usmeID)
	return err
}
