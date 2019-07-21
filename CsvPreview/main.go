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
		// Set up preview structs
		var preview_data PreviewData 
		
		// read file data
		read_record_count := 0
		csv_data := csv.New(active_file.Name())
		for csv_data.LoadData(); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {
			// record first row as "titles" and title "position"
			if read_record_count == 0 {
				position_count := 0
				for _, title := range csv_data.ActiveRecord {
					var data_column PdColumn
					data_column.Title = title
					data_column.Position = position_count 
					data_column.ExcelPosition = position_count + 1 
					position_count++
					preview_data.Columns = append(preview_data.Columns, data_column)
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
		fmt.Println(preview_data)
		fmt.Println("")
	}
}