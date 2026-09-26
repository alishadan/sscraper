package sscraper

import (
	"testing"
)

func Test_Sendmail(t *testing.T) {
	//input: (extracted_price string, url string, product string)
	//output : error
	extracted_price := "100,000"
	url := "https://example.com"
	product := "paper A4"
	err := SendMail(extracted_price, url, product)
	if err == nil {
		println("SendMail passed")
	} else {
		println("errors exist in SendMail function")
	}

}
