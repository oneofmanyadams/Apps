package main

import (
	"os"
	"time"

	"data/csv"
	"derp/ManufactureOrder"
)

type DerpEnv struct {
	MosFolder string
	TrackingFolder string
	ClassificationsFolder string

	MosCsvFile string
	MosJsonFile string
	MoTrackingFile string
	MoClassificationCsvFile string

	DerpStartTime time.Time
}

func (de DerpEnv) Setup() {
	if _, err :=os.Stat(de.MosFolder); os.IsNotExist(err) {
		os.Mkdir(de.MosFolder, os.ModePerm)
	}
	if _, err :=os.Stat(de.TrackingFolder); os.IsNotExist(err) {
		os.Mkdir(de.TrackingFolder, os.ModePerm)
	}
	if _, err :=os.Stat(de.ClassificationsFolder); os.IsNotExist(err) {
		os.Mkdir(de.ClassificationsFolder, os.ModePerm)
	}

}

func main() {
	var env DerpEnv

	env.MosFolder = "MOs"
	env.TrackingFolder = "Tracking"
	env.ClassificationsFolder = "MO_Classifications"
	env.MosCsvFile = env.MosFolder+"/mos.csv"
	env.MosJsonFile = env.MosFolder+"/mos.json"
	env.MoClassificationCsvFile = env.MosFolder+"/mos_classified.csv"
	env.MoTrackingFile = env.TrackingFolder+"/mo_tracking.json"
	
	env.DerpStartTime = time.Now()
	
	env.Setup()
	
	var existing_mos ManufactureOrder.Collection
	var new_mos ManufactureOrder.Collection

	// Load older MO data from JSON file 
	existing_mos.LoadFromJsonFile(env.MosJsonFile)

	// Load new MO data from csv file
	new_mos.LoadFromCsvFile(env.MosCsvFile)

	// Load tracking data from JSON file
	mo_tracker := ManufactureOrder.NewMoTrackerFromJsonFile(env.MoTrackingFile)

	// Load Classifications
	mo_classifier := ManufactureOrder.NewClassifier()
	mo_classifier.LoadClassificationsFromFolder(env.ClassificationsFolder)

	// Create template classification file if none exist
	if len(mo_classifier.Classifications) < 1 {
		ManufactureOrder.GenerateClassificationTemplateInFolder(env.ClassificationsFolder)
	}

	// Loop through new MOs
	for _, new_mo := range new_mos.Mos {
		existing_key := existing_mos.ExistingMoKey(new_mo.Number) // this return -1 if Mo Number does not exist in existing_mos

		// If MO in new data does not exist in previous data, create records and import.
		if existing_key == -1 {

			// classify MO
			mo_c := mo_classifier.ClassifyMo(new_mo)
			if mo_c != "" {
				new_mo.Classification = mo_c
			} else {
				new_mo.Classification = "UNKOWN"
			}
			existing_mos.Mos = append(existing_mos.Mos, new_mo)

			var new_mo_change ManufactureOrder.MoChange
			new_mo_change.ChangeType = "New MO"
			new_mo_change.ChangeTime = env.DerpStartTime
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
				var new_mo_change ManufactureOrder.MoChange

				if !mo_comparison.SameOrderQty {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" OrderQty"
				}
				if !mo_comparison.SameReceivedQty {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" ReceivedQty"
				}
				if !mo_comparison.SamePriority {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" Priority"
				}
				if !mo_comparison.SameMoStatus {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" MoStatus"
				}
				if !mo_comparison.SameMatStatus {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" MatStatus"
				}
				if !mo_comparison.SameStartDate {
					new_mo_change.ChangeType = new_mo_change.ChangeType+" StartDate"
				}

				new_mo_change.ChangeTime = env.DerpStartTime

				new_mo_change.CurrentOrderQty = new_mo.OrderQty
				new_mo_change.CurrentReceivedQty = new_mo.ReceivedQty
				new_mo_change.CurrentPriority = new_mo.Priority
				new_mo_change.CurrentMoStatus = new_mo.Status
				new_mo_change.CurrentMatStatus = new_mo.MaterialStatus
				new_mo_change.CurrentStartDate = new_mo.PlannedStartDate
	
				mo_tracker.RecordChange(new_mo.Number, new_mo_change)

				// Update new MO info to existing_mos data
				existing_mos.Mos[existing_key] = new_mo
			} else {
				if !mo_tracker.MoIsTracked(new_mo.Number) {
					var new_mo_change ManufactureOrder.MoChange
					new_mo_change.ChangeType = "New MO"
					new_mo_change.ChangeTime = env.DerpStartTime
					new_mo_change.CurrentOrderQty = new_mo.OrderQty
					new_mo_change.CurrentReceivedQty = new_mo.ReceivedQty
					new_mo_change.CurrentPriority = new_mo.Priority
					new_mo_change.CurrentMoStatus = new_mo.Status
					new_mo_change.CurrentMatStatus = new_mo.MaterialStatus
					new_mo_change.CurrentStartDate = new_mo.PlannedStartDate
		
					mo_tracker.RecordChange(new_mo.Number, new_mo_change)
				}
			}

		}


	}

	// Save updated existing MOs
	existing_mos.SaveToJsonFile(env.MosJsonFile)

	// Save updated MO tracking data
	mo_tracker.SaveDataToJsonFile(env.MoTrackingFile)

	// Write MO classification data to csv file
	classified_file := csv.WriteNewData(env.MoClassificationCsvFile)
	classified_file.WriteRecord([]string{"SKU","MO NUMBER", "CLASSIFICATION"})
	for _, mo := range existing_mos.Mos {
		classified_file.WriteRecord([]string{mo.Sku, mo.Number, mo.Classification})
	}
	classified_file.WriteRecordsToFile()
	

}