// Attempts to read all .json files in active directory as "CommonQuery".
// Runs the Query, saves the file based on the Query name
package main

import (	
	"fmt"

	"io/ioutil"
	"net/http"
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
			for query_name, query := range queries.Queries {
				ReqString(query.FullPath, query_name+".rdl")
			}
		}
	}


}

func ReqString(request_string string, save_location string) {

	resp, http_err := http.Get(request_string)
	if http_err != nil {
		fmt.Println(http_err)
		return
	}
	defer resp.Body.Close()
	
	data, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		fmt.Println(read_err)
		return
	}
	
	write_data := string(data)

	fmt.Println(write_data)
	fmt.Println(save_location)
	return
}
