package dbContext

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"codeid.revampacademy/models"
)

type CreateSkillParams struct {
	UskiID           int32          `db:"uski_id" json:"uskiId"`
	UskiEntityID     int32          `db:"uski_entity_id" json:"uskiEntityId"`
	UskiModifiedDate sql.NullTime   `db:"uski_modified_date" json:"uskiModifiedDate"`
	UskiSktyName     sql.NullString `db:"uski_skty_name" json:"uskiSktyName"`
}

// GetList Users Skill
const listSkill = `-- name: ListSkill :many
SELECT uski_id, uski_entity_id, uski_modified_date, uski_skty_name FROM users.users_skill
ORDER BY uski_id
`

func (q *Queries) ListSkill(ctx context.Context) ([]models.UsersUsersSkill, error) {
	rows, err := q.db.QueryContext(ctx, listSkill)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.UsersUsersSkill
	for rows.Next() {
		var i models.UsersUsersSkill
		if err := rows.Scan(
			&i.UskiID,
			&i.UskiEntityID,
			&i.UskiModifiedDate,
			&i.UskiSktyName,
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

// Get User Skill

const getSkill = `-- name: GetSkill :one

SELECT uski_id, uski_entity_id, uski_modified_date, uski_skty_name FROM users.users_skill
WHERE uski_id = $1
`

// Get Users Skill
func (q *Queries) GetSkill(ctx context.Context, uskiID int32) (models.UsersUsersSkill, error) {
	row := q.db.QueryRowContext(ctx, getSkill, uskiID)
	var i models.UsersUsersSkill
	err := row.Scan(
		&i.UskiID,
		&i.UskiEntityID,
		&i.UskiModifiedDate,
		&i.UskiSktyName,
	)
	return i, err
}

// Create User License
const createSkill = `-- name: CreateSkill :one

INSERT INTO users.users_skill
(uski_id, uski_entity_id, uski_modified_date, uski_skty_name)
VALUES($1,$2,$3,$4)
RETURNING *
`

func (q *Queries) CreateSkill(ctx context.Context, arg CreateSkillParams) (*models.UsersUsersSkill, *models.ResponseError) {
	row := q.db.QueryRowContext(ctx, createSkill,
		arg.UskiID,
		arg.UskiEntityID,
		sql.NullTime{Time: time.Now(), Valid: true},
		arg.UskiSktyName,
	)
	i := models.UsersUsersSkill{}
	err := row.Scan(
		&i.UskiID,
		&i.UskiEntityID,
		&i.UskiModifiedDate,
		&i.UskiSktyName,
	)

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	return &models.UsersUsersSkill{
		UskiID:           i.UskiID,
		UskiEntityID:     i.UskiEntityID,
		UskiModifiedDate: i.UskiModifiedDate,
		UskiSktyName:     i.UskiSktyName,
	}, nil
}

// Update Skill
const updateSkill = `-- name: UpdateSkill :exec
UPDATE users.users_skill
  set uski_entity_id = $2,
  uski_modified_date = $3,
  uski_skty_name= $4
WHERE uski_id = $1
`

func (q *Queries) UpdateSkill(ctx context.Context, arg CreateSkillParams) error {
	_, err := q.db.ExecContext(ctx, updateSkill,
		arg.UskiID,
		arg.UskiEntityID,
		arg.UskiModifiedDate,
		arg.UskiSktyName,
	)
	return err
}

// Delete Users Skill
const deleteSkill = `-- name: DeleteSkill :exec
DELETE FROM users.users_skill
WHERE uski_id = $1
`

func (q *Queries) DeleteSkill(ctx context.Context, uskiID int32) error {
	_, err := q.db.ExecContext(ctx, deleteSkill, uskiID)
	return err
}
