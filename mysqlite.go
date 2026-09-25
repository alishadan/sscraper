package mysqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	"CLIScraper/codes/kind"

)

func Sq(Books []kind.Book) {
	filename := "./Book.db"
	db := connect_sqlite3(filename)
	if db == nil {
		print("dont create connect \n")
	}
	defer db.Close()

	_,err:=create_table(db)
	if err!=nil{
		fmt.Printf("something is wrong in create_table function \n %v \n",err)
		panic("ttt")
	}

	_,err=insert_record(db,Books)
	if err!=nil{
		fmt.Printf("error ocuured in insert_record \n %v \n",err)
		panic("fff")
	}

	fmt.Printf("data saved in Books.db \n")

}

func connect_sqlite3(filename string) (db *sql.DB) {
	//connect to the SQLITE
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	println("connected to db successfull")

	//how to use:
	// filename := "./my.db"
	// db := connect_sqlite3(filename)
	// if db==nil{
	// 	print("dont create connect \n")
	// }
	// defer db.Close()
	return db

}
func create_table(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS Books(
		id INTEGER PRIMARY KEY,
		title TEXT UNIQE NOT NULL,
		price INTEGER NOT NULL,
		image_url TEXT NOT NULL
	);`
	return db.Exec(sql)

	//HOW to use:
	//_, err = create_table(db)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//} else {
	//	println("table created successfully")
	//}
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
	//Get the version of SQLITE
	//get_version(db)

}
func insert_record(db *sql.DB, Books []kind.Book) (int64, error) {
	query := `INSERT OR IGNORE INTO Books (title,price,image_url)
	VALUES (?,?,?)`

	var result sql.Result
	var err error

	for i,_:=range(Books){
		result, err = db.Exec(query, Books[i].Title, Books[i].Price, Books[i].UrlImage)
		if err != nil {
			fmt.Println("error occured in db.Exec(query) funciton")
			return 0, err
		}
	}
	
	return result.LastInsertId()

}
