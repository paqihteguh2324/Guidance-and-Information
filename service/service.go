package service

import (
	"context"

	"github.com/go-kit/kit/log"

	"sadhelx-be-guidelines/datastruct"
)

type (
	// Service ...
	Service interface {
		GetGuidelinesDocument(ctx context.Context) ([]datastruct.Guidelines, error)
		GetDocument(ctx context.Context, filename string) error
	}

	service struct {
		repository datastruct.GuidelinesRepository
		// configs    *util.Configurations
		logger log.Logger
	}
)

// NewService ...
func NewService(repo datastruct.GuidelinesRepository, logger log.Logger) Service {
	return &service{
		repository: repo,
		logger:     log.With(logger, "repo", "service"),
	}
}

func (s *service) GetGuidelinesDocument(ctx context.Context) ([]datastruct.Guidelines, error) {
	return s.repository.GetGuidelinesDocument(ctx)
}

func (s *service) GetDocument(ctx context.Context, filename string) error {
	return nil

}

// func HelloWorld(name string) string {
// 	var helloOutput string
// 	helloOutput = fmt.Sprintf("Hello, %s ", name)
// 	return helloOutput
// }
