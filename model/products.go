package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/somewhere/db"
)

type TProduct struct {
	ID        bson.ObjectId `json:"item_id" bson:"_id"`
	StoreID   bson.ObjectId `json:"store_id" bson:"store_id"`
	Name      string        `json:"item_name" bson:"item_name"`
	Price     float64       `json:"item_price" bson:"item_price"`
	Score     float64       `json:"item_score" bson:"item_score"`
	SaleCount int           `json:"item_salecount" bson:"item_salecount"`
	Brand     string        `json:"item_brand" bson:"item_brand"`
	Timestamp int64         `json:"item_timestamp" bson:"item_timestamp"`
}

func (t *TProduct) AddProduct() error {
	col := db.MgoDb.C("items")
	err := col.Insert(t)
	return err
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

func GetAllProducts() (Products []TProduct, err error) {
	col := db.MgoDb.C("items")
	var ret []TProduct
	err = col.Find(bson.M{}).All(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
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
