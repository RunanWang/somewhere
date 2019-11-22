package model

import (
	"github.com/somewhere/db"
)

type TStores struct {
	ID    int    `json:"store_id"`
	Name  string `json:"store_name"`
	Level int    `json:"store_level"`
}

func (t *TStores) AddStore() (int, error) {

	// Prepare statement for inserting data
	stmtIns, err := db.SqlDb.Prepare("INSERT INTO stores (name,level) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		return -1, err
	}
	defer stmtIns.Close()

	rs, err := stmtIns.Exec(t.Name, t.Level)
	if err != nil {
		return -1, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (t *TStores) GetStoreByName() (stores []*TStores, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM stores where name = ?", t.Name)
	if err != nil {
		return
	}

	var aStore TStores
	err = row.Scan(&aStore.ID, &aStore.Name, &aStore.Level)
	if err != nil {
		return
	}
	stores = append(stores, &aStore)

	return
}

func (t *TStores) GetStoreByID() (stores []*TStores, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM stores where id = ?", t.ID)
	if err != nil {
		return
	}
	var aStore TStores
	err = row.Scan(&aStore.ID, &aStore.Name, &aStore.Level)
	if err != nil {
		return
	}
	stores = append(stores, &aStore)

	return
}

func GetAllStores() (stores []*TStores, err error) {

	rows, err := db.SqlDb.Query("SELECT * from stores")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aStore TStores
		err = rows.Scan(&aStore.ID, &aStore.Name, &aStore.Level)
		if err != nil {
			return
		}
		stores = append(stores, &aStore)
	}
	return stores, nil
}

func (t *TStores) UpdateStore() (int, error) {

	stmt, err := db.SqlDb.Prepare("UPDATE stores SET name=?,level=? WHERE id=?")
	if err != nil {

		return -1, err
	}
	rs, err := stmt.Exec(t.Name, t.Level, t.ID)
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

func (t *TStores) DeleteStore() (int, error) {

	stmt, err := db.SqlDb.Prepare("DELETE FROM stores WHERE id=?")
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
