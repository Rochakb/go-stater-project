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
