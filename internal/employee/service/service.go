package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/downstream/repository"
	"github.com/unbxd/go-base/utils/log"
)

type (
	Service interface {
		GetEmployee(context.Context, interface{}) (interface{}, error)
	}
)

type svc struct {
	logger     log.Logger
	repository repository.Repository
}

func NewSvc(
	l log.Logger,
	r repository.Repository,
) Service {
	return &svc{
		l,
		r,
	}
}
