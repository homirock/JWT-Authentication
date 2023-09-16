package database

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
func AddDataToDatabase(){
// Establish a connection to the MySQL database
db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/keyvalue_db")
if err != nil {
	panic(err.Error())
}
defer db.Close()

// Insert a key-value pair
key := "myKey"
value := "myValue"

_, err = db.Exec("INSERT INTO keyvalue (`key`, `value`) VALUES (?, ?)", key, value)
if err != nil {
	panic(err.Error())
}

fmt.Printf("Inserted key-value pair: %s:%s\n", key, value)
}