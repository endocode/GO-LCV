package main

import (
	"encoding/csv"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"os"
)

type CsvLine struct {
	LicenseName string
	Classpath   string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

//Still to implement the logic Licenses Array against ComplianceMatrix
func LicenseValidation(s []string, CSVfilePath string) {

	lines, err := ReadCsv(CSVfilePath)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := CsvLine{
			LicenseName: line[0],
			Classpath:   line[1],
		}
		fmt.Println(data.LicenseName + " " + data.Classpath)
		fmt.Println(data.Classpath)
	}
}

type valuesIAmInterestedIn struct {
	Payload struct {
		FileMetadata []*struct {
			Licenses *[]string `json:"licenses"`
		} `json:"fileMetadata"`
	} `json:"payload"`
}

func LicensesExtractor(jsonFileName string) []string {
	var s []string
	var values valuesIAmInterestedIn
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		fmt.Println(err)
	} else {
		err = json.Unmarshal(jsonFile, &values)
		for _, gl := range values.Payload.FileMetadata {
			if gl.Licenses != nil {
				var licenses = *gl.Licenses
				for _, license := range licenses {
					if (contains(s, license)) == false {
						s = append(s, license)
					}
				}
			}
		}
	}
	return s
}

func main() {
	var jsonFileName = "javacpp_full.json"
	var CSVfilePath = "licenses.csv"
	var LicenseArray []string = LicensesExtractor(jsonFileName)
	// var records [][]string = ReadCSV(CSVfilePath)
	LicenseValidation(LicenseArray, CSVfilePath)
}
