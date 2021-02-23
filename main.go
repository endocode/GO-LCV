package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

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
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// TO IMPLEMENT THE VALIDATION AGAINST THE MATRIX
// func LicenseValidation(s []string) {
// 	for _, license := range s {
// 	}
// 	return false
// }

type valuesIAmInterestedIn struct {
	Payload struct {
		FileMetadata []*struct {
			Licenses *[]string `json:"licenses"`
		} `json:"fileMetadata"`
	} `json:"payload"`
}

func main() {
	var s []string
	var values valuesIAmInterestedIn
	jsonFile, err := ioutil.ReadFile("javacpp_full.json")
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
	for _, license := range s {
		fmt.Printf(license)
		fmt.Println("")
	}
	// LicenseValidation(s)
}
