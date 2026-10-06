package sscraper

import (
	"testing"
)
//for your uses, we neet to edit this functions







func Test_sq(t *testing.T) {
	books:=make([]book,1)
	books[0].Title="ha ha ha"
	books[0].Price=150000
	books[0].UrlImage="https://example.com"
	filename:="sqlite"
	queryCreateTable := `CREATE TABLE IF NOT EXISTS Books(
		id INTEGER PRIMARY KEY,
		title TEXT UNIQE NOT NULL,
		price INTEGER NOT NULL,
		image_url TEXT NOT NULL
);`

queryInsertRecord := `INSERT OR IGNORE INTO Books (title,price,image_url)
VALUES (?,?,?)`

	Sq(books,filename, queryCreateTable,queryInsertRecord,"ha ha ha",150000,"https://example.com")

}