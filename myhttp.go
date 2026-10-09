package sscraper

import (
	"io"
	"net/http"
	"os"
	"fmt"
)

//how to use: example
// in your function:
//body, _ := myhttp("https://google.com")
// defer body.Close()
//proccess body

func Myhttp(url string) (io.ReadCloser, error) {
	//url := "https://coinmarketcap.com/"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil,fmt.Errorf("error in NewRequest %q: %w",url,err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("download %q failed", url)
	}
	if resp.StatusCode > 299 {
		resp.Body.Close()
		return nil,fmt.Errorf("error in get %q : %s",url,resp.Status)
	}
	return resp.Body, nil

}

//how to use: example
// in your function:
//body, _ := myfile("crypt.html")
// defer body.Close()
//proccess body

func Myfile(filename string) (io.ReadCloser, error) {
	//fileName := "crypt.htm"

	file, err := os.Open(filename)
	if err != nil {
		print("error in opening ", filename, "\n")
		return nil, err
	}
	return file, err

}
