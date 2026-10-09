package sscraper

import (
	"testing"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"fmt"
)

type book struct{
	Title		string
	UrlImage	string
	Price 		float64
}
var Data1 []book
func Test_MyQuery(t *testing.T){
	
	url:="https://duckduckgo.com"
	body,err:=Myhttp(url)
	if err!=nil{
		panic("happend error in connect to site \n")
	}
	defer body.Close()

	wfind:="div._linkWrapper_1ts6x_1"

	err=MyQuery(url,wfind,myfunc)
	if err!=nil{
		fmt.Printf("Error exist in MyQuery function %v \n",err)
	}else{
		fmt.Println("MyQuery passed")
	}

}

func myfunc (i int, s *goquery.Selection){
	imgurl,_:=s.Find("img").Attr("src")
	price:=s.Find("strong").Text()
	title:=s.Find("h2._title_i1yaq_19").Text()
	price1,_:=strconv.ParseFloat(price,64)

	var book1 book 
	book1.Title=title
	book1.UrlImage=imgurl
	book1.Price=price1

	Data1=append(Data1,book1)
}