package repository

import (
	"database/sql"
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
)

type EmployeeRepository interface {
	BaseRepository[model.Employee]
	BaseRepositoryPaging[model.Employee]
}

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}

func (r *employeeRepository) Create(payload *model.Employee) (*model.Employee, error) {
	employee := &model.Employee{
		ID:          payload.ID,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		Address:     payload.Address,
	}

	stmt, err := r.db.Prepare(constant.EMPLOYEE_INSERT)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.PhoneNumber, payload.Address)

	if err != nil {
		return nil, err
	}

	return employee, nil

}
func (r *employeeRepository) List() ([]*model.Employee, error) {
	employees := []*model.Employee{}

	stmt, err := r.db.Prepare(constant.EMPLOYEE_LIST)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		employee := model.Employee{}

		rows.Scan(&employee.ID, &employee.Name, &employee.PhoneNumber, &employee.Address)

		employees = append(employees, &employee)

		fmt.Println(employee)
	}

	return employees, nil
}

func (r *employeeRepository) Get(id string) (*model.Employee, error) {
	employee := model.Employee{}

	stmt, err := r.db.Prepare(constant.EMPLOYEE_GET)

	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&employee.ID, &employee.Name, &employee.PhoneNumber, &employee.Address); err != nil {
		return nil, exception.ErrNotFound
	}

	return &employee, nil
}

func (r *employeeRepository) Update(payload *model.Employee) (*model.Employee, error) {
	employee := &model.Employee{
		ID:          payload.ID,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		Address:     payload.Address,
	}

	stmt, err := r.db.Prepare(constant.EMPLOYEE_UPDATE)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.PhoneNumber, payload.Address)

	if err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) Delete(id string) error {

	stmt, err := r.db.Prepare(constant.EMPLOYEE_DELETE)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	count, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	return nil
}

func (r *employeeRepository) Paging(requestPaging dto.PaginationParam, params ...string) ([]*model.Employee, dto.Paging, error) {
	paginationQuery := dto.PaginationQuery{}

	paginationQuery = common.GetPaginationParams(requestPaging)

	query := constant.EMPLOYEE_LIST

	if params[0] != "" {
		query += ` WHERE name ILIKE '%` + params[0] + `%'`
	}

	if params[1] != "" {
		query += ` WHERE phone_number ILIKE '%` + params[1] + `%'`
	}

	if params[2] != "" {
		query += ` WHERE address ILIKE '%` + params[2] + `%'`
	}

	query += ` LIMIT $1 OFFSET $2`

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	rows, err := stmt.Query(paginationQuery.Take, paginationQuery.Skip)

	if err != nil {
		return nil, dto.Paging{}, err
	}
	var employees []*model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.PhoneNumber, &employee.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		employees = append(employees, &employee)
	}

	// count total rows
	var totalRows int

	row := r.db.QueryRow(constant.EMPLOYEE_COUNT)

	err = row.Scan(&totalRows)

	if err != nil {
		return nil, dto.Paging{}, err
	}
	return employees, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}
