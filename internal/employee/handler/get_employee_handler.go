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

func getEmployeeDecoderFunc(_ context.Context, req *net_http.Request) (interface{}, error) {
	var (
		id = req.URL.Query().Get("empId")
	)
	request := models.GetEmpByIdRequest{EmpId: id}
	//err := json.NewDecoder(req.Body).Decode(&request)

	//if err != nil {
	//	return nil, errors.Wrap(
	//		errBadRequest, "failed to cast object to IndexKeys",
	//	)
	//}
	return &request, nil
}

func makeGetEmployeeEndpoint(service product_svc.Service) endpoint.Endpoint {
	return func(cx context.Context, req interface{}) (interface{}, error) {
		rq, ok := req.(*models.GetEmpByIdRequest)
		if !ok {
			return nil, errors.Wrap(
				errInternalServer, "failed to cast object to GetEmpByIdRequest",
			)
		}

		s, err := service.GetEmployee(cx, rq)
		if err != nil {
			return nil, err
		}

		return s, nil
	}
}

func getEmployeeEncoderFunc(_ context.Context, rw net_http.ResponseWriter, result interface{}) error {
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

func NewMakeGetEmployeeEndpointHandler(service product_svc.Service) http.Handler {
	return http.Handler(makeGetEmployeeEndpoint(service))
}

func NewGetEmployeeEndpointHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(getEmployeeDecoderFunc),
		http.HandlerWithEncoder(getEmployeeEncoderFunc),
		http.HandlerWithErrorEncoder(errorEncoder),
	}, opts...)
}
