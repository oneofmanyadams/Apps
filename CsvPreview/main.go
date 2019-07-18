// Read all csv files in a directory and saves priviews to file with "preview" suffix

package main

import (	
	"fmt"

	"io/ioutil"
	"path/filepath"
)

func main() {
	// Determine all files in active directory
	all_files, read_dir_err := ioutil.ReadDir("./")
	if read_dir_err != nil {
		fmt.Println(read_dir_err)
	}

	// Loop through files in active directory
	for _, a_file := range all_files {
		
		// Get file extension of file
		file_extension := filepath.Ext(a_file.Name())
		
		// We only care about atomsvc files
		if file_extension == ".atomsvc" {
			fmt.Println(a_file.Name())
		}
	}
}