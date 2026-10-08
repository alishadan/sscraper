package sscraper

import(
	"os"
	"encoding/json"
	"fmt"
)
//get data variable and filename and save datan on a json file

func SaveOnFile(data any, filename string) error {
	encodedData :=Encoding_data(data)
	//create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error in opening file: %w",err)
	}
	defer file.Close()

	//write data to the file
	_, err = file.Write(encodedData)
	if err != nil {
		return fmt.Errorf("error in file.WriteString() function :%w",err)
	}
	fmt.Println("data save on", filename, " successfully")
	return nil
}
func Encoding_data(data any) []byte {
	encodedData, err := json.Marshal(data)
	if err != nil {
		fmt.Errorf("some errors happen in encoding_data function")
		return nil
	}
	return encodedData
}

func Decoder[T any](filename string) ( []T,  error) {
	//open file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil,fmt.Errorf("error exist in os.ReadFile function: %w",err)
	}
	var products []T

	if err := json.Unmarshal(data, &products); err != nil {
		return nil,fmt.Errorf("error exist in json.Unmarshal function: %w",err)
    }
    return products,nil
}