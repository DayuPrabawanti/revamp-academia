package dbContext

import (
	"context"

	"codeid.revampacademy/models"
)

const listBatchTrinee = `-- name: ListBatchTrinee :many
SELECT batr_id, batr_status, batr_certificated, batre_certificate_link, batr_access_token, batr_access_grant, batr_review, batr_total_score, batr_modified_date, batr_trainee_entity_id, batr_batch_id FROM bootcamp.batch_trainee
ORDER BY batr_id
`

func (q *Queries) ListBatchTrinee(ctx context.Context) ([]models.BootcampBatchTrainee, error) {
	rows, err := q.db.QueryContext(ctx, listBatchTrinee)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.BootcampBatchTrainee
	for rows.Next() {
		var i models.BootcampBatchTrainee
		if err := rows.Scan(
			&i.BatrID,
			&i.BatrStatus,
			&i.BatrCertificated,
			&i.BatreCertificateLink,
			&i.BatrAccessToken,
			&i.BatrAccessGrant,
			&i.BatrReview,
			&i.BatrTotalScore,
			&i.BatrModifiedDate,
			&i.BatrTraineeEntityID,
			&i.BatrBatchID,
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
