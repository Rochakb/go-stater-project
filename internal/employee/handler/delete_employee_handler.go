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

func deleteEmployeeDecoderFunc(_ context.Context, req *net_http.Request) (interface{}, error) {
	var request model.Employee
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, errors.Wrap(
			errBadRequest, "failed to cast object to IndexKeys",
		)
	}
	return &request, nil
}

func makeDeleteEmployeeEndpoint(service product_svc.Service) endpoint.Endpoint {
	return func(cx context.Context, req interface{}) (interface{}, error) {
		rq, ok := req.(*model.Employee)
		if !ok {
			return nil, errors.Wrap(
				errInternalServer, "failed to cast object to GetEmpByIdRequest",
			)
		}

		s, err := service.DeleteEmployee(cx, rq)
		if err != nil {
			return nil, err
		}

		return s, nil
	}
}

func deleteEmployeeEncoderFunc(_ context.Context, rw net_http.ResponseWriter, result interface{}) error {
	resp := models.DeleteEmpResponse{
		Response: map[string]interface{}{
			"result": result,
		},
	}
	bt, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(net_http.StatusOK)
	rw.Write(bt)
	return nil
}

func NewMakeDeleteEmployeeEndpointHandler(service product_svc.Service) http.Handler {
	return http.Handler(makeDeleteEmployeeEndpoint(service))
}

func NewDeleteEmployeeEndpointHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(deleteEmployeeDecoderFunc),
		http.HandlerWithEncoder(deleteEmployeeEncoderFunc),
		http.HandlerWithErrorEncoder(errorEncoder),
	}, opts...)
}
