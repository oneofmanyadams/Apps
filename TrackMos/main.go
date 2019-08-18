package main

import (
	"time"

	"derp/ManufactureOrder"
)

func main() {
	mos_csv_file_location := "..\\MOs\\mos.csv"
	mos_json_file_location := "..\\MOs\\mos.json"
	tracking_file := "..\\Tracking\\mo_tracking.json"

	var existing_mos ManufactureOrder.Collection
	var new_mos ManufactureOrder.Collection
	
	// Load older MO data from JSON file 
	existing_mos.LoadFromJsonFile(mos_json_file_location)

	// Load new MO data from csv file
	new_mos.LoadFromCsvFile(mos_csv_file_location)

	// Load tracking data from JSON file
	mo_tracker := ManufactureOrder.NewMoTrackerFromJsonFile(tracking_file)

	// Loop through new MOs
	for _, new_mo := range new_mos.Mos {
		existing_key := existing_mos.ExistingMoKey(new_mo.Number) // this return -1 if Mo Number does not exist in existing_mos

		// If MO in new data does not exist in previous data, create records and import.
		if existing_key == -1 {
			existing_mos.Mos = append(existing_mos.Mos, new_mo)

			var new_mo_change ManufactureOrder.MoChange
			new_mo_change.ChangeType = "New MO"
			new_mo_change.ChangeTime = time.Now()
			new_mo_change.CurrentOrderQty = new_mo.OrderQty
			new_mo_change.CurrentReceivedQty = new_mo.ReceivedQty
			new_mo_change.CurrentPriority = new_mo.Priority
			new_mo_change.CurrentMoStatus = new_mo.Status
			new_mo_change.CurrentMatStatus = new_mo.MaterialStatus
			new_mo_change.CurrentStartDate = new_mo.PlannedStartDate

			mo_tracker.RecordChange(new_mo.Number, new_mo_change)
		} else {
			// If MO does exist in previous data, but something has changed, track change and update record.
			existing_mo := existing_mos.Mos[existing_key]
			mo_comparison := existing_mo.CompareTo(new_mo)
			if mo_comparison.Result() == false {

			}
	
			// Update new MO info to existing_mos data here 
			
			// If MO exists and nothing has changed, do nothing.


		}


	}


}