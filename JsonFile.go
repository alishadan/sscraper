package sscraper

import(
	"os"
	"encoding/json"
	"fmt"
)
//get url, price and filename and save url and price on a json file
//for your uses , you need to edit this funcitons

type Data struct {
	Site  string `json:"site"`
	Price string `json:"price"`
}

func SaveOnFile(uRl string, price string, filename string) error {
	encodedData :=Encoding_data(uRl, price)
	//create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		print("error in opening file \n")
		return err
	}
	defer file.Close()

	//write data to the file
	_, err = file.Write(encodedData)
	if err != nil {
		print("error in file.WriteString() function \n")
		return err
	}

	print("data save on", filename, " successfully \n")
	return nil
}
func Encoding_data(uRl string, price string) []byte {

	data1 := Data{uRl, price}
	encodedData, err := json.Marshal(data1)

	if err != nil {
		print("some errors happen in encoding_data function \n")
		return nil
	}
	return encodedData
}

func Decoder(filename string,data2 any) error {

	//open file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error exist in os.ReadFile function")
		return err
	}

	if err := json.Unmarshal(data, &data2); err != nil {
        panic(err)
        return err
    }
    return nil
}