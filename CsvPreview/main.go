// Read all csv files in a directory and saves priviews to file with "preview" suffix
package main

import (	
	"fmt"
	"strings"

	"io/ioutil"
	"path/filepath"

	"data/csv"
)

const (
	MAX_READ_ROWS = 4
)

func main() {
	// Determine all files in active directory
	all_files, read_dir_err := ioutil.ReadDir("./")
	if read_dir_err != nil {
		fmt.Println(read_dir_err)
	}

	// Loop through files in active directory
	for _, active_file := range all_files {
		
		// Get file extension of file
		file_extension := filepath.Ext(active_file.Name())
		
		// skip to next file if current file is not .csv
		if file_extension != ".csv" {
			continue
		}

		// set preview file name
		fmt.Println(strings.Replace(active_file.Name(), ".csv", "_preview.csv", -1))

		// read file data
		csv_data := csv.New(active_file.Name())
		read_record_count := 0
		for csv_data.LoadData(); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {


			
			if read_record_count > MAX_READ_ROWS {
				break
			}
			read_record_count++
			fmt.Println(csv_data.ActiveRecord)
		}
			

	}
}