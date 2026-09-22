package sscraper

import (
	"golang.org/x/net/html"
	"io"
)

//how to use
// for use of this funciton:
// we edit loop for or every IF: like in a page that calsses diffrent
// with this in here we change "sc-c1554bc0-0 eWrlhi coin-item-name"

func Myhtml(body io.ReadCloser, names *string) {
	tokenizer := html.NewTokenizer(body)
	if tokenizer.Err() != nil {
		print("error in html.NewTokenizer \n")
	}

	for {
		if tokenizer.Next() == html.ErrorToken {
			break
		}
		token1 := tokenizer.Token()
		if token1.Type == html.StartTagToken {
			if token1.Data == "p" {
				for _, value := range token1.Attr {
					if value.Val == "sc-c1554bc0-0 eWrlhi coin-item-name" {
						tokenizer.Next()
						token1 = tokenizer.Token()
						*names = token1.Data
					}
				}

			}
		}
	}

}
