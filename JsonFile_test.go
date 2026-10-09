package sscraper
import (
	"encoding/json"
	"testing"
	"fmt"
	
)
//define type struct
type Data struct {
	Site  string `json:"site"`
	Price string `json:"price"`
}

func Test_Encoding_data(t *testing.T) {
	var data1 Data
	data1.Site="http://example.com"
	data1.Price="100,00"

	byte_v := Encoding_data(data1)


	var result Data

	//decoding byte_v and save them in result
	err := json.Unmarshal(byte_v, &result)
	if err != nil {
		print("error in json.Unmarshal funciton in encoding test package \n")

	}

	if result.Price == data1.Price && result.Site == data1.Site {
		print("Encoding_data passed \n")

	} else {
		print("error exist in Encoding_data function \n")
	}

}


func Test_Save(t *testing.T) {
	//input []byte
	//output error

	var data1 Data
	data1.Site="http://example.com"
	data1.Price="100,00"
	filename := "new.txt"
	if err := SaveOnFile(data1, filename); err == nil {
		println("Save_on_file passed")
	} else {
		println("error exist in Save_on_file function")
	}
}

func Test_decoder(t *testing.T){
	//var data2 Data

	_,err:=Decoder[Data]("new.txt")
	if err!=nil{
		fmt.Printf("we have error in Decoder function: %v \n",err)
	}else{
		fmt.Println("Decoder Passed")
	}

}