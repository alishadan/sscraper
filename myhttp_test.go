package sscraper

import (
	"fmt"
	"io"
	"testing"
)

func Test_myhttp(t *testing.T) {
	//input: string
	//output: io.ReadCloser, error

	body, err := Myhttp(`https://google.com`)
	if err != nil {
		t.Skipf("Skipping test - network error: %v", err)
		return
	}
	defer body.Close()
	_, err = io.ReadAll(body)
	if err == nil {
		fmt.Println("myhtpp passed ")
	} else {
		t.Errorf("error exist in myhttp function")
	}

}
