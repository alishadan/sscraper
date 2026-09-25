package mysqlite

import (
	"testing"
)
type Book struct{
	Title		string
	UrlImage	string
	Price 		float32
}


func Test_sq(t *testing.T) {
	books:=make([]Book,1)
	books[0].Title="ha ha ha"
	books[0].Price=150000
	books[0].UrlImage="https://example.com"

	sq(books)

}
