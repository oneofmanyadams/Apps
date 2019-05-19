// Attempts to read all .json files in active directory as "CommonQuery".
// Displays the Query:
//	-Name, -FullPathLength
//	-FullPath
package main

import (	
	"fmt"

	"io/ioutil"
	"path/filepath"

	"data/commonquery"
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
		
		// We only care about json files
		if file_extension == ".json" {
			queries := commonquery.NewMultiQuery()
			queries.LoadFrom(a_file.Name())
			fmt.Println(a_file.Name())
			for query_name, query := range queries.Queries {
				fmt.Println(query_name, query.FullPathLength)
				fmt.Println(query.FullPath)
				fmt.Println("")
			}
			fmt.Println("")
		}
	}


}