package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const GetMasterIndustryImpl = `-- name: GetMasterIndustryImpl :one
	SELECT Indu_code_id, indu_name FROM master.industry
	WHERE Indu_code_id = $1
`

	func (q *Queries) GetMasterIndustryImpl(ctx context.Context, InduCodeID int32) (models.MasterIndustry, error) {
		row := q.db.QueryRowContext(ctx, GetMasterIndustryImpl, InduCodeID)
		var i models.MasterIndustry
		err := row.Scan(
			&i.InduCodeID,
			&i.InduName,
		)
		return i, err
	}

const ListMasterIndustryImpl = `-- name: ListMasterIndustryImpl :many
SELECT Indu_Code_Id, Indu_Name  FROM master.industry
ORDER BY Indu_Code_id
`

	func (q *Queries) ListMasterIndustryImpl(ctx context.Context) ([]models.MasterIndustry, error) {
					rows, err := q.db.QueryContext(ctx, ListMasterIndustryImpl)
					if err != nil {
						return nil, err
					}
					defer rows.Close()
					var items []models.MasterIndustry
					for rows.Next() {
						var i models.MasterIndustry
						if err := rows.Scan(
							&i.InduCodeID,
							&i.InduName,
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

type CreateMasterIndustryParams struct {
						InduCodeID           int32          `db:"joca_id" json:"jocaId"`
						InduName         string `db:"joca_name" json:"jocaName"`
		}