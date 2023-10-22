package repository

import (
	"database/sql"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
)

type CustomerRepository interface {
	BaseRepository[model.Customer]
	BaseRepositoryPaging[model.Customer]
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) Create(payload *model.Customer) (*model.Customer, error) {

	stmt, err := r.db.Prepare(constant.CUSTOMER_INSERT)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.PhoneNumber, payload.Address)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	return payload, nil

}
func (r *customerRepository) List() ([]*model.Customer, error) {
	customers := []*model.Customer{}

	stmt, err := r.db.Prepare(constant.CUSTOMER_LIST)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		customer := model.Customer{}

		rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.Address)

		customers = append(customers, &customer)

	}

	return customers, nil
}

func (r *customerRepository) Get(id string) (*model.Customer, error) {
	Customer := model.Customer{}

	stmt, err := r.db.Prepare(constant.CUSTOMER_GET)

	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&Customer.ID, &Customer.Name, &Customer.PhoneNumber, &Customer.Address); err != nil {
		return nil, exception.ErrNotFound
	}

	return &Customer, nil
}

func (r *customerRepository) Update(payload *model.Customer) (*model.Customer, error) {

	stmt, err := r.db.Prepare(constant.CUSTOMER_UPDATE)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.PhoneNumber, payload.Address)

	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (r *customerRepository) Delete(id string) error {

	stmt, err := r.db.Prepare(constant.CUSTOMER_DELETE)

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

func (r *customerRepository) Paging(requestPaging dto.PaginationParam, params ...string) ([]*model.Customer, dto.Paging, error) {
	paginationQuery := dto.PaginationQuery{}

	paginationQuery = common.GetPaginationParams(requestPaging)

	query := constant.CUSTOMER_LIST

	if params[0] != "" {
		query += ` WHERE name ILIKE '%` + params[0] + `%'`
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
	var customers []*model.Customer

	for rows.Next() {
		var customer model.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.Address)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		customers = append(customers, &customer)
	}

	// count total rows
	var totalRows int

	row := r.db.QueryRow(constant.CUSTOMER_COUNT)

	err = row.Scan(&totalRows)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	return customers, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}
