package main

import (
	"github.com/Rochakb/go-stater-project/internal/downstream/repository"
	"github.com/Rochakb/go-stater-project/internal/employee/control"
	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

func WithPostgres() Option {
	return func(empS *EmployeeService) (err error) {
		if db, err := repository.GetPostgresSQLRepositoryInstance(PostgresURL); err != nil {
			return errors.Wrap(err, "create postgress failed")
		} else {
			empS.repository = db
		}
		return
	}
}

func WithCustomLogger(
	level string,
	encoding string,
	output string,
) Option {
	return func(emp *EmployeeService) (err error) {
		logger, err := log.NewZapLogger(
			log.ZapWithLevel(level),
			log.ZapWithEncoding(encoding),
			log.ZapWithOutput([]string{output}),
		)

		if err != nil {
			return errors.Wrap(err, "failed to create logger")
		}

		emp.logger = logger
		return
	}
}

func WithHTTPTransport(
	host, port string,
	monitor []string,
) Option {
	return func(empS *EmployeeService) (err error) {
		tr, err := http.NewTransport(
			host,
			port,
			http.WithLogger(empS.logger),
			http.WithFullDefaults(),
			http.WithMonitors(monitor),
		)
		if err != nil {
			return err
		}

		empS.httpTransport = tr
		return
	}
}
func WithEmployeeServiceControlPlane() Option {
	return func(h *EmployeeService) (err error) {
		// check if logger is set, if not return error
		if h.logger == nil {
			return errors.New("Logger is required for products control plane")
		}

		h.logger.Debug(
			"Initializing EmployeeServiceControlPlane",
		)

		// create control object
		h.employeeControl = control.NewEmployeeControl(h.logger, h.repository)
		h.employeeControl.Bind(
			h.httpTransport,
			[]http.HandlerOption{
				http.NewSetRequestHeader("unx-server", "employeeService"),
				http.NewSetResponseHeader("unx-server", "employeeService"),
			}...,
		)
		return
	}
}

//func WithEmployeeServiceControlPlane() Option {
//	return func(empS *EmployeeService) (err error) {
//		empc := employeeInternal.NewControl(empS.logger, empS.repository)
//
//		empS.employeeControl = empc
//
//		empS.employeeControl.Bind(
//			empS.httpTransport,
//			[]http.HandlerOption{
//				http.HandlerWithFilter(
//					metrics_filter.NewHttpMetricFilter(
//						"rerank",
//						false,
//					),
//				),
//				http.HandlerWithErrorhandler(
//					http.ErrorHandler(err_hn.NewErrorHandler(
//						empS.logger,
//					)),
//				),
//				http.HandlerWithErrorEncoder(err_enc.ErrorEncoder),
//			}...,
//		)
//
//		return
//	}
//}
