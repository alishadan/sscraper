package sscraper
import(
	//"io"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	
)
// you need edit this function for your uses

func MyQuery(url string,wfind string,myfunc func (i int, s *goquery.Selection)) error {
	body,err:=Myhttp(url)
	if err!=nil{
		return fmt.Errorf("happend error in connect to site %w",err)
	}
	defer body.Close()
	doc,err:=goquery.NewDocumentFromReader(body)
	if err!=nil{
		return fmt.Errorf("error in goquery.NewDocument function %w",err)
	}
	myselection:=doc.Find(wfind)
	if myselection.Length()==0{
		fmt.Printf("doc.find dont find anything \n")
		return nil
	}

	myselection.Each(myfunc)
	return nil
}


