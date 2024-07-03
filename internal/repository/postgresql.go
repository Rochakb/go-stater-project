package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Rochakb/go-stater-project/internal/model"
)

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) *PostgreSQLRepository {
	return &PostgreSQLRepository{
		db: db,
	}
}

func (r *PostgreSQLRepository) GetEmployeeByID(ctx context.Context, empId int) (model.Employee, error) {
	query := `SELECT EmpId, Name, DOB, Department, Salary, BossId FROM Employee WHERE EmpId = $1`
	row := r.db.QueryRowContext(ctx, query, empId)

	var employee model.Employee
	err := row.Scan(&employee.EmployeeId, &employee.Name, &employee.DOB, &employee.Department, &employee.Salary, &employee.BossId)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Employee{}, fmt.Errorf("employee not found")
		}
		return model.Employee{}, fmt.Errorf("failed to scan employee row: %v", err)
	}
	return employee, nil
}

func (r *PostgreSQLRepository) CreateEmployee(ctx context.Context, employee model.Employee) (bool, error) {
	query := `
		INSERT INTO Employee (empid, name, dob, department, salary, bossId)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query, employee.EmployeeId, employee.Name, employee.DOB, employee.Department, employee.Salary, employee.BossId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostgreSQLRepository) UpdateEmployee(ctx context.Context, empId int, employee model.Employee) (bool, error) {
	query := `
		UPDATE Employee
		SET name = $1, dob = $2, department = $3, salary = $4, bossId = $5
		WHERE empid = $6
	`

	_, err := r.db.ExecContext(ctx, query, employee.Name, employee.DOB, employee.Department, employee.Salary, employee.BossId, empId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *PostgreSQLRepository) DeleteEmployee(ctx context.Context, empId int) (bool, error) {
	query := `DELETE FROM Employee WHERE empid = $1`

	_, err := r.db.ExecContext(ctx, query, empId)
	if err != nil {
		return false, err
	}
	return true, nil
}
