package main

import (
	"data/csv"
	"derp/ManufactureOrder"
)

func main() {
	mos_csv_file_location := "..\\MOs\\mos.csv"
	mos_json_file_location := "..\\MOs\\mos.json"
	classifications_csv_file := "..\\MOs\\mos_classified.csv"
	classifications_folder := "..\\MO_Classifications"

	var existing_mos ManufactureOrder.Collection
	var new_mos ManufactureOrder.Collection
	
	mo_classifier := ManufactureOrder.NewClassifier()


	// Load older MO data from json file 
	existing_mos.LoadFromJsonFile(mos_json_file_location)
	
	// Load MO data from Csv file
	new_mos.LoadFromCsvFile(mos_csv_file_location)

	// Load Classifications
	mo_classifier.LoadClassificationsFromFolder(classifications_folder)

	// Create template classification file if none exist
	if len(mo_classifier.Classifications) < 1 {
		ManufactureOrder.GenerateClassificationTemplateInFolder(classifications_folder)
	}

	// Import any new MOs from csv data into existing data
	existing_mos.ImportOtherCollection(new_mos)

	//	Loop through each MO, classifying each one.
	for mo_key, mo := range existing_mos.Mos {
		mo_c := mo_classifier.ClassifyMo(mo)
		if mo_c != "" {
			existing_mos.Mos[mo_key].Classification = mo_c
		} else {
			existing_mos.Mos[mo_key].Classification = "UNKOWN"
		}
	}

	// Save updated data back to json file.
	existing_mos.SaveToJsonFile(mos_json_file_location)

	// Write MO data to csv file
	classified_file := csv.WriteNewData(classifications_csv_file)
	classified_file.WriteRecord([]string{"SKU","MO NUMBER", "CLASSIFICATION"})
	for _, mo := range existing_mos.Mos {
		classified_file.WriteRecord([]string{mo.Sku, mo.Number, mo.Classification})
	}
	classified_file.WriteRecordsToFile()

}