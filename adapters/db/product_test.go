package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/filipegorges/hex/adapters/db"
	"github.com/filipegorges/hex/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		status string,
		price float
	);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product Test", "disabled", 0)`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "abc", product.GetID())
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 0.0, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	// create
	productDb := db.NewProductDb(Db)
	product := application.NewProduct()
	product.ID = "def"
	product.Name = "Product Test 2"
	product.Status = "enabled"
	product.Price = 10.0
	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "def", result.GetID())
	require.Equal(t, "Product Test 2", result.GetName())
	require.Equal(t, "enabled", result.GetStatus())
	require.Equal(t, 10.0, result.GetPrice())
	// update
	product.ID = "ghi"
	product.Name = "Product Test 3"
	product.Status = "disabled"
	product.Price = 20.0
	result, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "ghi", result.GetID())
	require.Equal(t, "Product Test 3", result.GetName())
	require.Equal(t, "disabled", result.GetStatus())
	require.Equal(t, 20.0, result.GetPrice())
}
