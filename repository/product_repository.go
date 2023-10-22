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

type ProductRepository interface {
	BaseRepository[model.Product]
	BaseRepositoryPaging[model.Product]
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(payload *model.Product) (*model.Product, error) {

	stmt, err := r.db.Prepare(constant.PRODUCT_INSERT)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.Price, payload.Uom)

	if err != nil {
		return nil, err
	}

	return payload, nil

}
func (r *productRepository) List() ([]*model.Product, error) {
	products := []*model.Product{}

	stmt, err := r.db.Prepare(constant.PRODUCT_LIST)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := model.Product{}

		rows.Scan(&product.ID, &product.Name, &product.Price, &product.Uom, product)

		products = append(products, &product)

		fmt.Println(product)
	}

	return products, nil
}

func (r *productRepository) Get(id string) (*model.Product, error) {
	product := model.Product{}

	stmt, err := r.db.Prepare(constant.PRODUCT_GET)

	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Uom); err != nil {
		return nil, exception.ErrNotFound
	}

	return &product, nil
}

func (r *productRepository) Update(payload *model.Product) (*model.Product, error) {

	stmt, err := r.db.Prepare(constant.PRODUCT_UPDATE)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(payload.ID, payload.Name, payload.Price, payload.Uom)

	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (r *productRepository) Delete(id string) error {

	stmt, err := r.db.Prepare(constant.PRODUCT_DELETE)

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

func (r *productRepository) Paging(requestPaging dto.PaginationParam, params ...string) ([]*model.Product, dto.Paging, error) {
	paginationQuery := dto.PaginationQuery{}

	paginationQuery = common.GetPaginationParams(requestPaging)

	query := constant.PRODUCT_LIST

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
	var products []*model.Product

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Uom)
		if err != nil {
			return nil, dto.Paging{}, err
		}
		products = append(products, &product)
	}

	// count total rows
	var totalRows int

	row := r.db.QueryRow(constant.PRODUCT_COUNT)

	err = row.Scan(&totalRows)

	if err != nil {
		return nil, dto.Paging{}, err
	}
	return products, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}
