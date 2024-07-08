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

func filterEmployeeDecoderFunc(_ context.Context, req *net_http.Request) (interface{}, error) {
	var request models.FilterRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode filter request")
	}
	return &request, nil
}

func makeFilterEmployeeEndpoint(service product_svc.Service) endpoint.Endpoint {
	return func(cx context.Context, req interface{}) (interface{}, error) {
		rq, ok := req.(*models.FilterRequest)
		if !ok {
			return nil, errors.Wrap(
				errInternalServer, "failed to cast object to FilterRequest",
			)
		}

		s, err := service.FilterEmployee(cx, rq)
		if err != nil {
			return nil, err
		}

		return s, nil
	}
}

func filterEmployeeEncoderFunc(_ context.Context, rw net_http.ResponseWriter, result interface{}) error {
	resp := models.FilterResponse{Response: result.([]model.Employee)}
	bt, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(net_http.StatusOK)
	rw.Write(bt)
	return nil
}

func NewMakeFilterEmployeeEndpointHandler(service product_svc.Service) http.Handler {
	return http.Handler(makeFilterEmployeeEndpoint(service))
}

func NewFilterEmployeeEndpointHandlerOption(opts []http.HandlerOption) []http.HandlerOption {
	return append([]http.HandlerOption{
		http.HandlerWithDecoder(filterEmployeeDecoderFunc),
		http.HandlerWithEncoder(filterEmployeeEncoderFunc),
		http.HandlerWithErrorEncoder(errorEncoder),
	}, opts...)
}
