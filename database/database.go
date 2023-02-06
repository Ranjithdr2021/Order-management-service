package database

import (
	"database/sql"
	"fmt"
)

// ConnectToDB connects to the database --> orders
func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Ranjith@tcp(localhost:3306)/orders")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Connected to DB Successfully....... ")
	return db
}

// NewTable creates new table if the table not exist
func NewTable() {
	db := ConnectToDB()
	defer db.Close()
	_, err := db.Query("CREATE TABLE IF NOT EXISTS orders(O_id varchar(20) UNIQUE NOT NULL, Status varchar(20) NOT NULL, Total int(6) NOT NULL, CurrencyUnit varchar(10) NOT NULL, PRIMARY KEY(O_id))")
	if err != nil {
		fmt.Println(err)
	}
	_, e := db.Query("CREATE TABLE IF NOT EXISTS items(I_id varchar(10) UNIQUE NOT NULL, Description varchar(20), Price int(10) NOT NULL, Quantity int(6) NOT NULL, PRIMARY KEY(I_id), FOREIGN KEY(O_id) REFERENCES orders(O_id))")
	if e != nil {
		fmt.Println(e)
	}
}
