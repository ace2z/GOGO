package CUSTOM_GOMOD

import (
	//. "local/CORE"

	//	. "github.com/ace2z/TWC"

	. "github.com/ace2z/GOGO/Gadgets"
	"go.mongodb.org/mongo-driver/bson"
)

// All generic MONGO RECORDS have ID field. this is a unique field that you generate based on
// your record criteria
type GENERIC_MONGO_RECORD struct {
	ID        string
	SAVE_DATA interface{}
}

var DBPREFIX = "==| "

func GEN_MONGO_RECORD(ID string, data interface{}) GENERIC_MONGO_RECORD {
	var rec GENERIC_MONGO_RECORD
	rec.ID = ID
	rec.SAVE_DATA = data

	return rec
}

// Alias
func GMR(ID string, data interface{}) GENERIC_MONGO_RECORD {
	return GEN_MONGO_RECORD(ID, data)
}

func DB_not_EXIST(tmpid string, LIST *[]GENERIC_MONGO_RECORD) bool {
	for _, x := range *LIST {
		if x.ID == tmpid {
			return false
		}
	}

	return true
}

// This will save all records using a "cache" of max_cache_size records at a time
// Also accepts -verbose and a number in secs representing Sleep between cache saves
func CACHED_Mongo_SAVE(collname string, RECS []GENERIC_MONGO_RECORD, max_cache_size int, PARAMS ...interface{}) {

	var sleep_between = 0
	var verbose = false
	for _, field := range PARAMS {
		val_int, IS_INT := field.(int)
		//val_float, IS_FLOAT := field.(float64)
		val_string, IS_STRING := field.(string)
		//val_bool, IS_BOOL := field.(bool)

		if IS_STRING {
			if val_string == "-verbose" {
				verbose = true
			}
			continue
		}

		// If INT is passed, means we are to sleep between cache saves
		if IS_INT {
			sleep_between = val_int
		}

	}

	// Now the cached save routine
	var CACHE []GENERIC_MONGO_RECORD
	for _, x := range RECS {
		CACHE = append(CACHE, x)
		if len(CACHE) >= max_cache_size {
			if verbose {
				Y.Print(DBPREFIX, "Cache FULL, now ")
				G.Println("Saving..")
			}
			BULK_Mongo_SAVE(collname, CACHE, PARAMS...)
			// always purge the cache
			CACHE = []GENERIC_MONGO_RECORD{}

			if sleep_between > 0 {
				if verbose {
					C.Print(DBPREFIX, "...Sleeping ")
					W.Print(sleep_between, " secs ")
					C.Println("between Cache Saves...")
					Sleep(sleep_between, false)
				}
			}

		}
	}

	// Now at the end.. lets do a final safety purge of whatever may still be in cache
	if len(CACHE) > 0 {
		BULK_Mongo_SAVE(collname, CACHE, PARAMS)
	}
}

// Most recent bulk save
func BULK_Mongo_SAVE(COLL_NAME string, RECS []GENERIC_MONGO_RECORD, PARAMS ...interface{}) {
	var verbose = false
	var purge_first = false

	for _, field := range PARAMS {
		//val_int, IS_INT := field.(int)
		//val_float, IS_FLOAT := field.(float64)
		val_string, IS_STRING := field.(string)
		//val_bool, IS_BOOL := field.(bool)

		if IS_STRING {
			if val_string == "-verbose" {
				verbose = true
				continue
			}
			if val_string == "-purge" {
				purge_first = true
				continue
			}
		}
	}

	if len(RECS) <= 0 {
		if verbose {
			Y.Println(DBPREFIX, "Incoming Record set is empty!")
		}
		return
	}

	//Create a list of all the IDs in the payload
	var LIST []string
	for _, x := range RECS {
		LIST = append(LIST, x.ID)
	}

	//2. Now, Create a search filter for this LIST using $in operator
	var filter bson.D
	filter = append(filter, bson.E{
		"ID", bson.D{
			{"$in", LIST},
		},
	})

	/*
	  If Purge First was set, we delete all the items found before re-inserting
	  This ensures whatever is passed always overwrites anything existing
	*/
	if purge_first {
		if verbose {
			C.Print(DBPREFIX)
			BW.Print(" PURGE existing BEFORE Save ")
			Y.Println("")
		}

		DO_BULK_DELETE(DBNAME, COLL_NAME, filter)
		Sleep(1, false) // always sleep 1 second after a delete..
	}

	//3. WE ALWAYS NEED TO CREATE THE FINAL_DELTAS Interface...
	// Even if we ran bulk delete. This is what ultimatly gets saved to Mongo
	// So first search for any existing (if we just ran delete, nothing will come back)
	err, RES_CURSOR := MON_SEARCH(DBNAME, COLL_NAME, filter)
	if err != nil {
		M.Print(DBPREFIX, "MONSEARCH failed for Some reason!! on ")
		Y.Println(COLL_NAME)
		W.Println(DBPREFIX, "This should, normally, never happen!")
		DO_EXIT()
	}
	var RES []GENERIC_MONGO_RECORD
	RES_CURSOR.All(MONGO_CONTEXT, &RES)

	//4. Now CHECK the results. save anything that DOESNT already exist (from the list) in the FINAL_DELTAS
	var FINAL_DELTAS []interface{}
	for _, newOBJ := range RECS {
		var SKIP = false
		for _, res := range RES {
			if res.ID == newOBJ.ID {
				SKIP = true
				break
			}
		}
		if SKIP == false {
			FINAL_DELTAS = append(FINAL_DELTAS, newOBJ.SAVE_DATA)
		}
	}

	//5. Finally Insert ONLY the deltas (new items)
	if len(FINAL_DELTAS) > 0 {
		DO_BULK_INSERT(DBNAME, COLL_NAME, FINAL_DELTAS)
		if verbose {
			G.Println(DBPREFIX, "Save Complete")
		}
	}
}
