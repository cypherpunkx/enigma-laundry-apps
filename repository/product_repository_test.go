package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"enigmacamp.com/enigma-laundry-apps/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dmProducts []model.Product = []model.Product{
	{
		ID:    "1232",
		Name:  "rasdsad",
		Price: 20000,
		Uom:   "KG",
	},
}

type ProductRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ProductRepository
}

func (suite *ProductRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Printf("Error : %s", err)
	}

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewProductRepository(db)
}

func (suite *ProductRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func TestProductRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ProductRepositoryTestSuite))
}

func (suite *ProductRepositoryTestSuite) TestCreateNewProductSuccess() {
	product := dmProducts[0]
	suite.mockSql.ExpectExec("INSERT INTO products (.+)").WithArgs(product.ID, product.Name, product.Price, product.Uom).WillReturnResult(sqlmock.NewResult(1, 1))

	data, actualError := suite.repo.Create(&product)

	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
	assert.Equal(suite.T(), product, data)
}
func (suite *ProductRepositoryTestSuite) TestCreateNewProductFail() {
	product := dmProducts[0]
	suite.mockSql.ExpectExec("INSERT INTO products (.+)").WithArgs(product.ID, product.Name, product.Price, product.Uom).WillReturnError(fmt.Errorf("error"))

	_, actualError := suite.repo.Create(&product)

	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)

}
