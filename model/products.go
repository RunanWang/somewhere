package model

import (
	"github.com/somewhere/db"
)

type TProduct struct {
	ID      int    `json:"product_id"`
	StoreID int    `json:"store_id"`
	Name    string `json:"product_name"`
	Price   int    `json:"product_price"`
}

func (t *TProduct) AddProduct() (int, error) {

	// Prepare statement for inserting data
	stmtIns, err := db.SqlDb.Prepare("INSERT INTO products (store_id,name,price) VALUES(?, ?, ? )") // ? = placeholder
	if err != nil {
		return -1, err
	}
	defer stmtIns.Close()

	rs, err := stmtIns.Exec(t.StoreID, t.Name, t.Price)
	if err != nil {
		return -1, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (t *TProduct) GetProductByName() (Products []*TProduct, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM products where name = ?", t.Name)
	if err != nil {
		return
	}

	var aProduct TProduct
	err = row.Scan(&aProduct.ID, &aProduct.Name, &aProduct.Price)
	if err != nil {
		return
	}
	Products = append(Products, &aProduct)

	return
}

func (t *TProduct) GetProductByID() (Products []*TProduct, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM products where id = ?", t.ID)
	if err != nil {
		return
	}
	var aProduct TProduct
	err = row.Scan(&aProduct.ID, &aProduct.Name, &aProduct.Price)
	if err != nil {
		return
	}
	Products = append(Products, &aProduct)

	return
}

func GetAllProducts() (Products []*TProduct, err error) {

	rows, err := db.SqlDb.Query("SELECT * from products")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aProduct TProduct
		err = rows.Scan(&aProduct.ID, &aProduct.StoreID, &aProduct.Name, &aProduct.Price)
		if err != nil {
			return
		}
		Products = append(Products, &aProduct)
	}
	return Products, nil
}

func (t *TProduct) UpdateProduct() (int, error) {

	stmt, err := db.SqlDb.Prepare("UPDATE products SET name=?,store_id=?,price=? WHERE id=?")
	if err != nil {

		return -1, err
	}
	rs, err := stmt.Exec(t.Name, t.StoreID, t.Price, t.ID)
	if err != nil {

		return -1, err
	}

	row, err := rs.RowsAffected()
	if err != nil {

		return -1, err
	}
	defer stmt.Close()

	return int(row), nil
}

func (t *TProduct) DeleteProduct() (int, error) {

	stmt, err := db.SqlDb.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {

		return -1, err
	}

	rs, err := stmt.Exec(t.ID)
	if err != nil {

		return -1, err
	}
	row, err := rs.RowsAffected()
	if err != nil {

		return -1, err
	}
	defer stmt.Close()

	return int(row), nil
}
