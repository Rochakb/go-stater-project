package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Rochakb/go-stater-project/internal/model"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type PostgresSQLRepository struct {
	db *sql.DB
}

var (
	pc   *sql.DB
	lock *sync.Mutex = &sync.Mutex{}
)

func GetPostgresSQLRepositoryInstance(dbURI string) (Repository, error) {
	lock.Lock()
	defer lock.Unlock()

	if pc != nil {

		return &PostgresSQLRepository{pc}, nil
	}

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
		return nil, err
	}
	return &PostgresSQLRepository{db: db}, nil
}

func (r *PostgresSQLRepository) GetEmployeeByID(ctx context.Context, empId string) (model.Employee, error) {
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

func (r *PostgresSQLRepository) CreateEmployee(ctx context.Context, employee model.Employee) (bool, error) {
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

func (r *PostgresSQLRepository) UpdateEmployee(ctx context.Context, empId string, employee model.Employee) (bool, error) {
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

func (r *PostgresSQLRepository) DeleteEmployee(ctx context.Context, empId string) (bool, error) {
	query := `DELETE FROM Employee WHERE empid = $1`

	_, err := r.db.ExecContext(ctx, query, empId)
	if err != nil {
		return false, err
	}
	return true, nil
}
