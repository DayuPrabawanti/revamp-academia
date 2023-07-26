package curriculumRepositories

import "database/sql"

type RepositoryManager struct {
	ProgEntityRepository
	ProgEntityDescRepository
	ProgReviewsRepository
	SectionDetailMaterialRepository
}

func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewProgEntityRepository(dbHandler),
		*NewProgEntityDescRepository(dbHandler),
		*NewProgReviewsRepository(dbHandler),
		*NewSectionDetailMaterialRepository(dbHandler),
	}
}
