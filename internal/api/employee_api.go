package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Rochakb/go-stater-project/internal/model"
	"github.com/Rochakb/go-stater-project/internal/service"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Request and Response structs
type getEmployeeRequest struct {
	EmployeeId int `json:"employeeId"`
}

type getEmployeeResponse struct {
	Employee model.Employee `json:"employee,omitempty"`
	Error    string         `json:"error,omitempty"`
}

type createEmployeeRequest struct {
	Employee model.Employee `json:"employee"`
}

type createEmployeeResponse struct {
	Created bool   `json:"created"`
	Error   string `json:"error,omitempty"`
}

type updateEmployeeRequest struct {
	EmployeeId int            `json:"employeeId"`
	Employee   model.Employee `json:"employee"`
}

type updateEmployeeResponse struct {
	Updated bool   `json:"updated"`
	Error   string `json:"error,omitempty"`
}

type deleteEmployeeRequest struct {
	EmployeeId int `json:"employeeId"`
}

type deleteEmployeeResponse struct {
	Deleted bool   `json:"deleted"`
	Error   string `json:"error,omitempty"`
}

// Request and Response structs
type getEmployeesUnderBossRequest struct {
	BossId int `json:"bossId"`
}

type getEmployeesUnderBossResponse struct {
	Employees []model.Employee `json:"employees"`
	Error     string           `json:"error,omitempty"`
}

//func makeGetEmployeesUnderBossEndpoint(svc service.EmployeeService) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(getEmployeesUnderBossRequest)
//		employees, err := svc.GetEmployeesUnderBoss(ctx, req.BossId)
//		if err != nil {
//			return getEmployeesUnderBossResponse{Error: err.Error()}, nil
//		}
//		return getEmployeesUnderBossResponse{Employees: employees}, nil
//	}
//}

// Request decoding function
func decodeGetEmployeesUnderBossRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getEmployeesUnderBossRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Service Endpoints
func makeGetEmployeeEndpoint(svc service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getEmployeeRequest)
		employee, err := svc.GetEmployee(ctx, req.EmployeeId)
		if err != nil {
			return getEmployeeResponse{Error: err.Error()}, nil
		}
		return getEmployeeResponse{Employee: employee}, nil
	}
}

func makeCreateEmployeeEndpoint(svc service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createEmployeeRequest)
		created, err := svc.CreateEmployee(ctx, req.Employee)
		if err != nil {
			return createEmployeeResponse{Created: false, Error: err.Error()}, nil
		}
		return createEmployeeResponse{Created: created}, nil
	}
}

func makeUpdateEmployeeEndpoint(svc service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateEmployeeRequest)
		updated, err := svc.UpdateEmployee(ctx, req.EmployeeId, req.Employee)
		if err != nil {
			return updateEmployeeResponse{Updated: false, Error: err.Error()}, nil
		}
		return updateEmployeeResponse{Updated: updated}, nil
	}
}

func makeDeleteEmployeeEndpoint(svc service.EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteEmployeeRequest)
		deleted, err := svc.DeleteEmployee(ctx, req.EmployeeId)
		if err != nil {
			return deleteEmployeeResponse{Deleted: false, Error: err.Error()}, nil
		}
		return deleteEmployeeResponse{Deleted: deleted}, nil
	}
}

// Request decoding functions
func decodeGetEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCreateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDeleteEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// Response encoding functions
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// Endpoints struct to group all endpoints
type Endpoints struct {
	GetEmployeeEndpoint    endpoint.Endpoint
	CreateEmployeeEndpoint endpoint.Endpoint
	UpdateEmployeeEndpoint endpoint.Endpoint
	DeleteEmployeeEndpoint endpoint.Endpoint
}

// NewEndpoints initializes all endpoints
func NewEndpoints(svc service.EmployeeService) Endpoints {
	return Endpoints{
		GetEmployeeEndpoint:    makeGetEmployeeEndpoint(svc),
		CreateEmployeeEndpoint: makeCreateEmployeeEndpoint(svc),
		UpdateEmployeeEndpoint: makeUpdateEmployeeEndpoint(svc),
		DeleteEmployeeEndpoint: makeDeleteEmployeeEndpoint(svc),
	}
}

// MakeHTTPHandler creates a new HTTP handler using Go Kit's transport
func MakeHTTPHandler(endpoints Endpoints) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/employee/get", httptransport.NewServer(
		endpoints.GetEmployeeEndpoint,
		decodeGetEmployeeRequest,
		encodeResponse,
	))

	mux.Handle("/employee/create", httptransport.NewServer(
		endpoints.CreateEmployeeEndpoint,
		decodeCreateEmployeeRequest,
		encodeResponse,
	))

	mux.Handle("/employee/update", httptransport.NewServer(
		endpoints.UpdateEmployeeEndpoint,
		decodeUpdateEmployeeRequest,
		encodeResponse,
	))

	mux.Handle("/employee/delete", httptransport.NewServer(
		endpoints.DeleteEmployeeEndpoint,
		decodeDeleteEmployeeRequest,
		encodeResponse,
	))

	return mux
}
