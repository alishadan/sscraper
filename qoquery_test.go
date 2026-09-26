package sscraper

import (
	"testing"
	//"github.com/alishadan/sscraper"
)
func Test_MyQuery(t *testing.T){
	url:="https://google.com"
	body,err:=Myhttp(url)
	if err!=nil{
		panic("happend error in connect to site \n")
	}
	defer body.Close()

	Books:=MyQuery(body,url)

	if Books!=nil{
		println(Books[0].Title)

	}




}