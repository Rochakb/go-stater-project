package control

import (
	"github.com/Rochakb/go-stater-project/internal/downstream/repository"
	"github.com/Rochakb/go-stater-project/internal/employee/handler"
	"github.com/Rochakb/go-stater-project/internal/employee/service"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

type EmployeeControl struct {
	service service.Service
}

func (c *EmployeeControl) Service() service.Service { return c.service }

func (c *EmployeeControl) Bind(ht *http.Transport, opts ...http.HandlerOption) {
	ht.GET(
		"/employee",
		handler.NewMakeGetEmployeeEndpointHandler(c.service),
		handler.NewGetEmployeeEndpointHandlerOption(opts)...,
	)
	ht.PUT(
		"/employee/create",
		handler.NewMakeCreateEmployeeEndpointHandler(c.service),
		handler.NewCreateEmployeeEndpointHandlerOption(opts)...,
	)
	ht.DELETE(
		"/employee",
		handler.NewMakeDeleteEmployeeEndpointHandler(c.service),
		handler.NewDeleteEmployeeEndpointHandlerOption(opts)...,
	)
	ht.PUT(
		"/employee/update",
		handler.NewMakeUpdateEmployeeEndpointHandler(c.service),
		handler.NewUpdateEmployeeEndpointHandlerOption(opts)...,
	)
	ht.POST(
		"/employee/filter",
		handler.NewMakeFilterEmployeeEndpointHandler(c.service),
		handler.NewFilterEmployeeEndpointHandlerOption(opts)...,
	)
}

func NewEmployeeControl(
	l log.Logger,
	r repository.Repository,
) *EmployeeControl {
	svc := service.NewSvc(l, r)
	return &EmployeeControl{
		svc,
	}
}
