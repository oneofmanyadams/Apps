package main

import (
	"fmt"
	
	"derp/ManufactureOrder"
)

func main() {
	mos_csv_file_location := "..\\MOs\\mos.csv"
	mos_json_file_location := "..\\MOs\\mos.json"
	classifications_folder := "..\\MO_Classifications"

	var existing_mos ManufactureOrder.Collection
	var new_mos ManufactureOrder.Collection
	var mo_classifier ManufactureOrder.Classifier


	// Load older MO data from json file 
	existing_mos.LoadFromJsonFile(mos_json_file_location)
	
	// Load MO data from Csv file
	new_mos.LoadFromCsvFile(mos_csv_file_location)

	// Import any new MOs from csv data into existing data
	existing_mos.ImportOtherCollection(new_mos)

	// Save updated data back to json file.
	existing_mos.SaveToJsonFile(mos_json_file_location)

	// Load Classifications
	mo_classifier.LoadClassificationsFromFolder(classifications_folder)

	if len(mo_classifier.Classifications) < 1 {
		ManufactureOrder.GenerateClassificationTemplateInFolder(classifications_folder)
	}

	fmt.Println(mo_classifier)
}