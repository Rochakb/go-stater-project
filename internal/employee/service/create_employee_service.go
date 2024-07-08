package service

import (
	"context"
	"encoding/json"
	"github.com/Rochakb/go-stater-project/internal/model"
	"github.com/unbxd/go-base/utils/log"
)

func (s *svc) CreateEmployee(ctx context.Context, req interface{}) (interface{}, error) {
	createEmpRequest := req.(*model.Employee)
	reqString, _ := json.Marshal(createEmpRequest)
	s.logger.Debug("Received Create request:", log.String("a", string(reqString)))
	// Log the received DOB value to debug
	s.logger.Debug("Received DOB:", log.String("dob", createEmpRequest.DOB.String()))

	resp, er := s.repository.CreateEmployee(ctx, *createEmpRequest)

	if er != nil {
		s.logger.Error("Error creating a new employee",
			log.String("service", "CreateEmployee"),
			log.String("err", er.Error()))
		return nil, er
	}
	return resp, nil
}
