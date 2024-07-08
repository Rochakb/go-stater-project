package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/employee/models"
	"github.com/unbxd/go-base/utils/log"
)

func (s *svc) FilterEmployee(ctx context.Context, req interface{}) (interface{}, error) {
	filterRequest := req.(*models.FilterRequest)

	resp, er := s.repository.FilterEmployee(ctx, filterRequest.BossId)

	if er != nil {
		s.logger.Error("Error in filter employee api",
			log.String("service", "FilterEmployee"),
			log.String("err", er.Error()))
		return nil, er
	}
	return resp, nil
}
