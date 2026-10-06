package sscraper

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)
func Sq[T any](data1 []T, filename string, queryCreateTable string,queryInsertRecord string, insertData ... any )error {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer db.Close()
	println("connected to db successfull")

	_,err=db.Exec(queryCreateTable)
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	//inser record
	for range(data1){
		_, err1 := db.Exec(queryInsertRecord, insertData... )
		if err1 != nil {
			fmt.Println("error occured in db.Exec(query) funciton")
			return err
		}
	}

	fmt.Println("data saved in",filename)

	return nil

}


func get_version(db *sql.DB) {
	var sqlite_v string
	err := db.QueryRow("select sqlite_version()").Scan(&sqlite_v)
	if err != nil {
		println(err)
		return
	}
	fmt.Printf("%s \n", sqlite_v)

	//how to use:
	//get_version(db)

}

