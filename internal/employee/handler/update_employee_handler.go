package handler

import (
	"context"
	"encoding/json"
	"github.com/Rochakb/go-stater-project/internal/model"
	net_http "net/http"

	"github.com/Rochakb/go-stater-project/internal/employee/models"
	product_svc "github.com/Rochakb/go-stater-project/internal/employee/service"
	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/endpoint"
	"github.com/unbxd/go-base/kit/transport/http"
)

func updateEmployeeDecoderFunc(_ context.Context, req *net_http.Request) (interface{}, error) {
	var request model.Employee
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, errors.Wrap(
			errBadRequest, "failed to cast object to IndexKeys",
		)
	}
	return &request, nil
}

func makeUpdateEmployeeEndpoint(service product_svc.Service) endpoint.Endpoint {
	return func(cx context.Context, req interface{}) (interface{}, error) {
		rq, ok := req.(*model.Employee)
		if !ok {
			return nil, errors.Wrap(
				errInternalServer, "failed to cast object to GetEmpByIdRequest",
			)
		}

		s, err := service.UpdateEmployee(cx, rq)
		if err != nil {
			return nil, err
		}

		return s, nil
	}
}

func updateEmployeeEncoderFunc(_ context.Context, rw net_http.ResponseWriter, result interface{}) error {
	resp := models.GetEmpByIdResponse{Response: result.(model.Employee)}
	bt, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(net_http.StatusOK)
	rw.Write(bt)
	return nil
}

func NewMakeUpdateEmployeeEndpointHandler(service product_svc.Service) http.Handler {
	return http.Handler(makeUpdateEmployeeEndpoint(service))
}

func NewUpdateEmployeeEndpointHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(updateEmployeeDecoderFunc),
		http.HandlerWithEncoder(updateEmployeeEncoderFunc),
		http.HandlerWithErrorEncoder(errorEncoder),
	}, opts...)
}
