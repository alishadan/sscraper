package sscraper

import (
	"my_library/myhttp"
	"testing"
)

func Test_myhtml(t *testing.T) {
	//input: body io.ReadCloser, price *string
	//output

	body, err := myhttp.Myhttp(`https://google.com`)
	if err != nil {
		t.Skipf("Skipping test - network error: %v", err)
		return
	}
	defer body.Close()
	price := "100,000"

	Myhtml(body, &price)
	println("myhtml function passed ")

}
