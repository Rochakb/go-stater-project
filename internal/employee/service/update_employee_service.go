package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/model"
	"github.com/unbxd/go-base/utils/log"
)

func (s *svc) UpdateEmployee(ctx context.Context, req interface{}) (interface{}, error) {
	updateEmpRequest := req.(*model.Employee)

	resp, er := s.repository.UpdateEmployee(ctx, updateEmpRequest.EmpId, *updateEmpRequest)

	if er != nil {
		s.logger.Error("Error updating an employee",
			log.String("service", "UpdateEmployee"),
			log.String("err", er.Error()))
		return nil, er
	}
	return resp, nil
}
