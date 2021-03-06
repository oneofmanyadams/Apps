// Read all csv files in a directory and saves priviews to file with "_preview" suffix
package main

import (	
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"
	"path/filepath"

	"data/csv"
)

const (
	MAX_READ_ROWS = 3
	SUFFIX = "_preview"
	POOL = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func main() {
	// Create character pool used for making excel columns
	cp := NewPool(POOL)

	// Determine all files in active directory
	all_files, read_dir_err := ioutil.ReadDir("./")
	if read_dir_err != nil {
		fmt.Println(read_dir_err)
	}

	// Loop through files in active directory
	for _, active_file := range all_files {
		// Get file name
		file_name := active_file.Name()

		// Get file extension of file
		file_extension := filepath.Ext(file_name)

		// skip to next file if current file is not .csv
		if file_extension != ".csv" {
			continue
		}

		// Skip to the next file if current file ends in "_preview.csv".
		// Prevents from creating preview file of a preview file.
		if len(file_name) > len(SUFFIX+".csv")+1 && (file_name[len(file_name)-len(SUFFIX+".csv"):] == SUFFIX+".csv") {
			continue
		}

		// Set up preview structs (from csv_preview.go file)
		var preview_data PreviewData 
		
		// read file data
		read_record_count := 0
		csv_data := csv.New(file_name)
		// Loop through each row of csv file
		for csv_data.LoadData(); csv_data.HasMoreRecords; csv_data.LoadNextRecord() {
			
			if read_record_count == 0 {
			// record first row as "titles" and title "position"
				position_count := 0

				// Loop through each record in csv row.
				for _, title := range csv_data.ActiveRecord {
					// create cata column record, and add all data
					// 	(except sample data, which is added in next iterations of this loop)	
					var data_column PdColumn // (from csv_preview.go file)
					data_column.Title = title
					data_column.Position = position_count 
					data_column.ExcelPosition = position_count + 1 
					data_column.ExcelColumn = cp.CharacterFromNumber(position_count) // (from position_calculator.go)
					
					// Add column record to preview data
					preview_data.Columns = append(preview_data.Columns, data_column)

					position_count++
				}
			} else {
			// Sample data is recorded in this block
				position_count := 0
				for _, sample_data := range csv_data.ActiveRecord {
					preview_data.Columns[position_count].Samples = append(preview_data.Columns[position_count].Samples, sample_data)
					position_count++
				}
			
			}

			// Exit loop if row limit is reached.
			if read_record_count >= MAX_READ_ROWS {
				break
			}
			read_record_count++
		}

		// set preview file name
		preview_file_name := strings.Replace(file_name, ".csv", SUFFIX+".csv", -1)

		// Create preview csv file
		preview_file := csv.WriteNewData(preview_file_name)

		// Create title row
		titles := []string{"Title", "Position", "ExcelPosition", "ExcelColumn"}
		// Dynamically determine # of samples based on how many iterations of the
		// number of rows looped through previously
		for sample_count := 0; sample_count < read_record_count; sample_count++ {
			titles = append(titles, "Sample "+strconv.Itoa(sample_count+1))
		}

		// "Add title row to csv preview data"
		preview_file.WriteRecord(titles)

		for _, pd_column := range preview_data.Columns {
			// Add standard data to preview csv row
			column_list := []string{pd_column.Title,
				strconv.Itoa(pd_column.Position),
				strconv.Itoa(pd_column.ExcelPosition),
				pd_column.ExcelColumn}
			
			// Add sample data to csv row
			for _, pd_column_sample := range pd_column.Samples {
				column_list = append(column_list, pd_column_sample)
			}

			// Add row to row to preview data
			preview_file.WriteRecord(column_list)
		}

		// Write all preview data to file
		preview_file.WriteRecordsToFile()
			
	}

}