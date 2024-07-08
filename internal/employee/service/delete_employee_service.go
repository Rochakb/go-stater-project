package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/model"
	"github.com/unbxd/go-base/utils/log"
)

func (s *svc) DeleteEmployee(ctx context.Context, req interface{}) (interface{}, error) {
	deleteEmpRequest := req.(*model.Employee)

	resp, er := s.repository.DeleteEmployee(ctx, deleteEmpRequest.EmpId)

	if er != nil {
		s.logger.Error("Error deleting an employee",
			log.String("service", "DeleteEmployee"),
			log.String("err", er.Error()))
		return nil, er
	}
	return resp, nil
}
