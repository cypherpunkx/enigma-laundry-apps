package repository

import (
	"database/sql"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
)

type BillRepository interface {
	Create(payload model.Bill) error
	Get(id string) (*model.Bill, error)
	List(requestPaging dto.PaginationParam, params ...string) ([]*dto.BillResponse, dto.Paging, error)
	GetBillDetailByBill(id string) ([]*model.BillDetails, error)
}

type billRepository struct {
	db *sql.DB
}

func NewBillRepository(db *sql.DB) BillRepository {
	return &billRepository{db: db}
}

func (r *billRepository) List(requestPaging dto.PaginationParam, params ...string) ([]*dto.BillResponse, dto.Paging, error) {
	paginationQuery := common.GetPaginationParams(requestPaging)

	query := constant.BILL_LIST

	stmt, err := r.db.Prepare(query)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, dto.Paging{}, err
	}

	var billResponses []*dto.BillResponse

	for rows.Next() {
		var currentBill dto.BillResponse
		var currentBillDetails dto.BillDetailsResponse

		err := rows.Scan(
			&currentBill.ID,
			&currentBill.BillDate,
			&currentBill.EntryDate,
			&currentBill.Employee.ID,
			&currentBill.Employee.Name,
			&currentBill.Employee.PhoneNumber,
			&currentBill.Employee.Address,
			&currentBill.Customer.ID,
			&currentBill.Customer.Name,
			&currentBill.Customer.PhoneNumber,
			&currentBill.Customer.Address,
			&currentBillDetails.ID,
			&currentBillDetails.BillID,
			&currentBillDetails.Product.ID,
			&currentBillDetails.Product.Name,
			&currentBillDetails.Product.Price,
			&currentBillDetails.Product.Uom,
			&currentBillDetails.Qty,
			&currentBillDetails.FinishDate,
		)

		if err != nil {
			return nil, dto.Paging{}, err
		}

		if currentBill.ID == currentBillDetails.BillID {
			currentBillDetails.ProductPrice = currentBillDetails.Product.Price
			currentBill.TotalBill += currentBillDetails.ProductPrice * currentBillDetails.Qty
			currentBill.BillDetails = append(currentBill.BillDetails, currentBillDetails)
		}

		billResponses = append(billResponses, &currentBill)
	}

	// count total rows
	var totalRows int

	row := r.db.QueryRow(constant.BILL_COUNT)

	err = row.Scan(&totalRows)

	if err != nil {
		return nil, dto.Paging{}, err
	}

	return billResponses, common.Paginate(paginationQuery.Page, paginationQuery.Take, totalRows), nil
}

func (r *billRepository) Create(payload model.Bill) error {

	tx, err := r.db.Begin()

	if err != nil {
		tx.Rollback()
		return err
	}

	// INSERT BILL
	_, err = tx.Exec(constant.BILL_CREATE, payload.ID, payload.BillDate, payload.EntryDate, payload.EmployeeID, payload.CustomerID)

	if err != nil {
		tx.Rollback()
		return err
	}

	// INSERT BILL DETAILS
	for _, item := range payload.BillDetails {
		_, err := tx.Exec(constant.BILL_DETAIL_CREATE, item.ID, item.BillID, item.ProductID, item.ProductPrice, item.Qty, item.FinishDate)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *billRepository) Get(id string) (*model.Bill, error) {
	var bill model.Bill

	err := r.db.QueryRow(constant.BILL_GET, id).Scan(&bill.ID, &bill.BillDate, &bill.EntryDate, &bill.EmployeeID, &bill.CustomerID)

	if err != nil {
		return nil, err
	}

	return &bill, nil
}

func (r *billRepository) GetBillDetailByBill(id string) ([]*model.BillDetails, error) {
	var billDetails []*model.BillDetails

	rows, err := r.db.Query(constant.BILL_DETAIL_GET, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var billDetail model.BillDetails
		err := rows.Scan(&billDetail.ID, &billDetail.BillID, &billDetail.ProductID, &billDetail.ProductPrice, &billDetail.Qty, &billDetail.FinishDate)

		if err != nil {
			return nil, err
		}

		billDetails = append(billDetails, &billDetail)
	}

	return billDetails, nil
}
