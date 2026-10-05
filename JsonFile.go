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
func Encoding_data(data any) []byte {

	encodedData, err := json.Marshal(data)

	if err != nil {
		print("some errors happen in encoding_data function \n")
		return nil
	}
	return encodedData
}

func Decoder(filename string,data any) error {

	//open file
	data1, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error exist in os.ReadFile function")
		return err
	}

	if err := json.Unmarshal(data1, &data); err != nil {
        panic(err)
        return err
    }
    return nil
}