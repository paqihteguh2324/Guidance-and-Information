package repository

import (
	"context"
	"database/sql"
	"sadhelx-be-guidelines/datastruct"

	"github.com/go-kit/kit/log"
)

type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo handle all db operation
func NewRepo(db *sql.DB, logger log.Logger) datastruct.GuidelinesRepository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}
}

const (
	queryGetGuidelines = "SELECT row_number()over() id, * FROM vw_active_guidelines where guidelines_type = 'document'"
)

func (repo *repo) GetGuidelinesDocument(ctx context.Context) ([]datastruct.Guidelines, error) {
	var guidelines []datastruct.Guidelines
	var err error

	rows, err := repo.db.QueryContext(ctx, queryGetGuidelines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var guideline datastruct.Guidelines

		if err = rows.Scan(
			&guideline.ID,
			&guideline.GuidelinesName,
			&guideline.GuidelinesDescription,
			&guideline.GuidelinesType,
			&guideline.GuidelinesLink,
		); err != nil {
			return nil, err
		}
		guidelines = append(guidelines, guideline)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return guidelines, nil
}
