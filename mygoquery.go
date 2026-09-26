package sscraper
import(
	"io"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"github.com/alishadan/sscraper"
	//"CLIScraper/codes/kind"
	"strconv"
)


var books []Book

func MyQuery(body io.ReadCloser, url string) []Book{
	//url:="https://www.noon.com/uae-en/books/health-and-personal-development/mind-body-and-spirit/mind-body-spirit-self-help/?f%5BisCarousel%5D=true"
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

	showBooks(books)

	return books
}

func myfunc (i int, s *goquery.Selection){
	imgurl,_:=s.Find("img").Attr("src")
	price:=s.Find("strong").Text()
	title:=s.Find("h2._title_i1yaq_19").Text()
	price1,_:=strconv.ParseFloat(price,64)
	var book1 Book
	book1.Title=title
	book1.UrlImage=imgurl
	book1.Price=price1

	books=append(books,book1)

}
func showBooks(books []Book){
	for i,_:=range(books){
		fmt.Printf("%s \n %f  \n ____________ \n",books[i].Title,books[i].Price)
	}
}