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

type PreviewData struct {
	Columns []PdColumn
}

type PdColumn struct {
	Title string
	Position int
	ExcelPosition int
	ExcelColumn string
	Samples []string
}

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

		// Set up preview struct

		// read file data
		csv_data := csv.New(active_file.Name())
		read_record_count := 0

		for csv_data.LoadData(); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {
			// record first row as "titles" and title "position"
			if read_record_count == 0 {
				position_count := 0
				for _, title := range csv_data.ActiveRecord {
					
					fmt.Println(title)
					position_count++
				}
			}

			// Exit loop if row limit is reached.
			if read_record_count > MAX_READ_ROWS {
				break
			}
			read_record_count++
		}
			
		// set preview file name
		fmt.Println(strings.Replace(active_file.Name(), ".csv", "_preview.csv", -1))

	}
}