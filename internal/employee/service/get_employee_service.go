package service

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/employee/models"
	"github.com/unbxd/go-base/utils/log"
)

func (s *svc) GetEmployee(
	ctx context.Context, req interface{},
) (interface{}, error) {

	getEmpByIdRequest := req.(*models.GetEmpByIdRequest)
	empId := getEmpByIdRequest.EmpId

	resp, er := s.repository.GetEmployeeByID(ctx, empId)

	if er != nil {
		s.logger.Error("creating pg index",
			log.String("service", "getEmployee"),
			log.String("err", er.Error()))
		return nil, er
	}
	return resp, nil

}
