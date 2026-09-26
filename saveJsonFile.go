package sscraper

import(
	"os"
	"encoding/json"
)
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
	type data struct {
		Site  string `json:"site"`
		Price string `json:"price"`
	}
	data1 := data{uRl, price}
	encodedData, err := json.Marshal(data1)

	if err != nil {
		print("some errors happen in encoding_data function \n")
		return nil
	}
	return encodedData
}