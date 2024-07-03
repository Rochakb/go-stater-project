package repository

import (
	"context"
	"github.com/Rochakb/go-stater-project/internal/model"
)

// Repository defines the methods that must be implemented by a repository.
type Repository interface {
	GetEmployeeByID(ctx context.Context, empId int) (model.Employee, error)
	CreateEmployee(ctx context.Context, employee model.Employee) (bool, error)
	DeleteEmployee(ctx context.Context, empId int) (bool, error)
	UpdateEmployee(ctx context.Context, empId int, employee model.Employee) (bool, error)
	//GetEmployeesUnderBoss(ctx context.Context, bossId int) ([]model.Employee, error)
}
