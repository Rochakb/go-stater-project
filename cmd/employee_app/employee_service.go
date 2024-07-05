package main

import (
	"context"
	"fmt"
	"github.com/Rochakb/go-stater-project/internal/downstream/repository"
	"github.com/Rochakb/go-stater-project/internal/employee/control"
	"os"
	"os/signal"
	"reflect"
	"runtime"

	"github.com/pkg/errors"
	"github.com/unbxd/go-base/kit/transport/http"
	"github.com/unbxd/go-base/utils/log"
)

var employeeService *EmployeeService

type (
	Option func(*EmployeeService) error

	EmployeeService struct {
		httpTransport   *http.Transport
		repository      repository.Repository
		logger          log.Logger
		employeeControl *control.EmployeeControl
	}
)

func (r *EmployeeService) listen(transport *http.Transport, errch chan error) {
	err := transport.Open()
	if err != nil {
		errch <- errors.Wrap(err, "failed to start transport")
	}
}

func (r *EmployeeService) Close(cx context.Context) error {
	r.logger.Flush()
	return nil
}

func (r *EmployeeService) Open(cx context.Context) error {
	r.logger.Info("-- Starting EmployeeService!")

	var (
		intchan = make(chan os.Signal, 1)
		errchan = make(chan error)
	)

	go r.listen(r.httpTransport, errchan)
	go signal.Notify(intchan, os.Interrupt)

	for {
		select {
		case <-intchan:
			r.logger.Info("Recieved os.Interrupt. Signal: ",
				log.String("signal", os.Interrupt.String()))
			r.logger.Info("Shutting down gracefully!!!")

			err := r.httpTransport.Close()
			if err != nil {
				panic(err)
			}

			r.logger.Info("Done!")
			return nil

		case err := <-errchan:
			r.logger.Error("Failed to start employeeService server:", log.Error(err))
			return err
		}
	}
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetEmployeeServiceInstance(options ...Option) (*EmployeeService, error) {
	var (
		tr, _ = http.NewTransport("0.0.0.0", "8000")
		lg, _ = log.NewZapLogger()
	)

	o := &EmployeeService{
		httpTransport: tr,
		logger:        lg,
	}

	for _, ofn := range options {
		fmt.Println(">> ----- Initializing: ", getFunctionName(ofn))
		err := ofn(o)
		if err != nil {
			return nil, err
		}
	}

	return o, nil
}
