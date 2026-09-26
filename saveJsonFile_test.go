package sscraper
import (
	"encoding/json"
	"testing"
)

func Test_Encoding_data(t *testing.T) {
	//input: string,string //output: []byte
	url := "http://example.com"
	price := "100,00"
	byte_v := Encoding_data(url, price)

	//define type struct
	type data struct {
		Site  string `json:"site"`
		Price string `json:"price"`
	}
	var result data

	//decoding byte_v and save them in result
	err := json.Unmarshal(byte_v, &result)
	if err != nil {
		print("error in json.Unmarshal funciton in encoding test package \n")

	}

	if result.Price == price && result.Site == url {
		print("Encoding_data passed \n")

	} else {
		print("error exist in Encoding_data function \n")
	}

}


func Test_Save(t *testing.T) {
	//input []byte
	//output error

	price := "100,000"
	url := "example.com"
	filename := "new.txt"
	if err := SaveOnFile(url, price, filename); err == nil {
		println("Save_on_file passed")
	} else {
		println("error exist in Save_on_file function")
	}
}