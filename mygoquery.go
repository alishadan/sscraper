package sscraper
import(
	"io"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"github.com/alishadan/sscraper"
	//"CLIScraper/codes/kind"
	"strconv"
)
//example of goquery for scrapping
// you need edit this function for your uses
var data []any


func MyQuery(body io.ReadCloser, url string) []any{
	body,err:=Myhttp(url)
	if err!=nil{
		panic("happend error in connect to site \n")
	}
	defer body.Close()
	doc,err:=goquery.NewDocumentFromReader(body)
	if err!=nil{
		panic("error in goquery.NewDocument function")
	}
	myselection:=doc.Find("div._linkWrapper_1ts6x_1")
	if myselection.Length()==0{
		fmt.Println("doc.find dont find anything")
		return nil

	}

	myselection.Each(myfunc)
	return data
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

	data=append(data,book1)

}
