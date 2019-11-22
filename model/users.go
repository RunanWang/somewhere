package model

import (
	"github.com/somewhere/db"
)

type TUser struct {
	ID   int    `json:"user_id"`
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}

func (t *TUser) AddUser() (int, error) {

	// Prepare statement for inserting data
	stmtIns, err := db.SqlDb.Prepare("INSERT INTO users (name,age) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		return -1, err
	}
	defer stmtIns.Close()

	rs, err := stmtIns.Exec(t.Name, t.Age)
	if err != nil {
		return -1, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (t *TUser) GetUserByName() (users []*TUser, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM users where name = ?", t.Name)
	if err != nil {
		return
	}

	var aUser TUser
	err = row.Scan(&aUser.ID, &aUser.Name, &aUser.Age)
	if err != nil {
		return
	}
	users = append(users, &aUser)

	return
}

func (t *TUser) GetUserByID() (users []*TUser, err error) {
	row := db.SqlDb.QueryRow("SELECT * FROM users where id = ?", t.ID)
	if err != nil {
		return
	}
	var aUser TUser
	err = row.Scan(&aUser.ID, &aUser.Name, &aUser.Age)
	if err != nil {
		return
	}
	users = append(users, &aUser)

	return
}

func GetAllUsers() (users []*TUser, err error) {

	rows, err := db.SqlDb.Query("SELECT * from users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var aUser TUser
		err = rows.Scan(&aUser.ID, &aUser.Name, &aUser.Age)
		if err != nil {
			return
		}
		users = append(users, &aUser)
	}
	return users, nil
}

func (t *TUser) UpdateUser() (int, error) {

	stmt, err := db.SqlDb.Prepare("UPDATE users SET name=?,age=? WHERE id=?")
	if err != nil {

		return -1, err
	}
	rs, err := stmt.Exec(t.Name, t.Age, t.ID)
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

func (t *TUser) DeleteUser() (int, error) {

	stmt, err := db.SqlDb.Prepare("DELETE FROM users WHERE id=?")
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
