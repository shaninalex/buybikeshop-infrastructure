package employee

import (
	"buybikeshop/apps/datasource/app/models"
	"context"
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

var (
	ErrEmployeeNotCreated = errors.New("no rows affected")
)

type Repository struct {
	db *sql.DB
}

func ProvideRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (s *Repository) EmployeeGet(ctx context.Context, id uuid.UUID) (*models.Employee, error) {
	q, _, err := goqu.From("admin.employees_departments").
		Select("employee_id", "department_id").
		Where(goqu.Ex{"employee_id": id}).
		ToSQL()
	if err != nil {
		return nil, err
	}
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var p models.Employee
	if err := rows.Scan(&p.EmployeeId, &p.DepartmentId); err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Repository) EmployeeSave(ctx context.Context, e models.Employee) (*models.Employee, error) {
	record := goqu.Record{
		"employee_id":   e.EmployeeId,
		"department_id": e.DepartmentId,
	}

	query, _, err := goqu.Insert("admin.employees_departments").Rows(record).ToSQL()
	if err != nil {
		return nil, err
	}

	result, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rows == 0 {
		return nil, ErrEmployeeNotCreated
	}

	return &e, nil
}

func (s *Repository) DepartmentGet(ctx context.Context, title string) (*models.Department, error) {
	q, _, err := goqu.From("admin.departments").
		Select("id", "title").
		Where(goqu.Ex{"title": title}).
		ToSQL()
	if err != nil {
		return nil, err
	}
	row := s.db.QueryRowContext(ctx, q)
	if err != nil {
		return nil, err
	}

	if err = row.Err(); err != nil {
		return nil, err
	}

	var p models.Department
	if err = row.Scan(&p.Id, &p.Title); err != nil {
		return nil, err
	}

	return &p, nil
}
