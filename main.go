package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// func hello() {
// 	fmt.Println("Hello, World!")
// }

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func ReadJSONFileMarshal() {

	// Open our jsonFile
	jsonFile, err := os.Open("javacpp.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened jacacpp.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["payload"])

}

func readJSONToken(fileName string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(fileName)
	defer file.Close()

	decoder := json.NewDecoder(file)

	filteredData := []map[string]interface{}{}

	// Read the array open bracket
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)

		if filter(data) {
			filteredData = append(filteredData, data)
		}
	}

	return filteredData
}

func LicensesReadToken(b *testing.B) {
	for n := 0; n < b.N; n++ {
		readJSONToken("javacpp.json", func(data map[string]interface{}) bool {
			return data["licenses"].(string) != "null"
		})
	}
}

func main() {
	// hello()
	//ReadJSONFileMarshal()
	LicensesReadToken(b * testing.B)

	records := readCsvFile("licenses.csv")
	fmt.Println(records)
}
