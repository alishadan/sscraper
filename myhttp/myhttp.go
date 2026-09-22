package myhttp

import (
	"io"
	"net/http"
	"os"
)

//how to use: example
// in your function:
//body, _ := myhttp("https://google.com")
// defer body.Close()
//proccess body

func Myhttp(url string) (io.ReadCloser, error) {
	//url := "https://coinmarketcap.com/"
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		print("error in NewRequest \n")
		return nil, err1
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		print("download", url, " failed")
		return nil, err
	}
	if resp.StatusCode > 299 {
		print("error in get", url, "Status", resp.Status)
		return nil, err
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
