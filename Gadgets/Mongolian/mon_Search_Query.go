package CUSTOM_GOMOD

import (    
    "strings"

    // For MOnGO using Latest Driver
    
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"

    . "github.com/acedev0/GOGO/Gadgets"
)


// Gets a COUNT of documents based on filter.. MAKE SURE AN INDEX IS SETUP ON TEH FIELD BEING QUERIED
func MON_COUNT(dbname string, coll_name string, search_filter interface{}) int{
 
	SHOW_BOX("Performing MONGO Count MONGO ")
	SHOW_STRUCT(search_filter)

	
	//2. Defualt options we alway ssetup
	

	var tmp_collection_obj = MONGO_CLIENT.Database(dbname).Collection(coll_name)
	var itemCOUNT, err = tmp_collection_obj.CountDocuments(MONGO_CONTEXT, search_filter)
	if err != nil {
		M.Println(" Error in MON_COUNT: ")
		Y.Println(err)
	}

	return int(itemCOUNT)
}

/*
1. This is a basic query function. You provide it a custom Mongo Filter like this:
	
		filter = bson.D{
			{
				"PRICE.Open", bson.D{
					{"$eq", 19.31},
				},
			},
		}	

	NOTE: if you want to query an array within a collection..use tTHIS format:

		filter = bson.D{
			{"instock", bson.D{
				{"$elemMatch", bson.D{
					{"qty", bson.D{
						{"$gt", 10},
						{"$lte", 20},
					}},
				}},
			}},
		})

	AND.. if you need to compound several query elements together..use this:

		var filter bson.D	
		filter = append(filter, bson.E {
			"ID", bson.D {
				{ "$in", LIST },
			},
		})

2. Invoke as follows (it returns a cursor with objects you can iterate over)

	var RES []FV_NEWS_OBJ		// object of whatever struct you're pulling back
	RES_CURSOR := MON_SEARCH(DBNAME, NEWS_COLL, filter)	
	RES_CURSOR.All(MONGO_CONTEXT, &RES)

	for _, res := range RES {
		C.Println(" HEadline is: ", res.Headline)
	}

 */
func MON_SEARCH(dbname string, coll_name string, search_filter interface{}, CLAUSES ...string) *mongo.Cursor {

	//Extract any clauses (like sort or limit)	
	var SORT_KEY = ""	
	var SORT_DIRECTION = -9
	var LIMIT = 0
	var BE_VERBOSE = false
	for _, clause := range CLAUSES {

		if strings.Contains(clause, "$sort_") {
			msplit := strings.Split(clause, ",")
			SORT_BY := msplit[0]
			SORT_KEY = msplit[1]

			if SORT_BY == "$sort_asc" {
				SORT_DIRECTION = 1
			} else if SORT_BY == "$sort_desc" {
				SORT_DIRECTION = -1
			}
			continue
		}

		if strings.Contains(clause, "$limit") {
			msplit := strings.Split(clause, ",")
			LIMIT = STRING_to_INT(msplit[1])
			continue
		}

		if clause == "$verbose" {
			BE_VERBOSE = true
		}
	}


	if BE_VERBOSE {
		SHOW_BOX("Querying MONGO ")
		C.Println(search_filter)
	}

	//2. Defualt options we alway ssetup
	var option_FILTER = bson.D{{"_id", 0}}
	var find_OPTIONS = options.Find()
	find_OPTIONS.SetProjection(option_FILTER)

	// If we are LIMITING
	if LIMIT > 0 {
		find_OPTIONS.SetLimit(int64(LIMIT))
	}
	// if we are SORTING by a key
	if SORT_KEY != "" {
		find_OPTIONS.SetSort(bson.D{{SORT_KEY, SORT_DIRECTION}})
	}

	var tmp_collection_obj = MONGO_CLIENT.Database(dbname).Collection(coll_name)
	var tmp_cursor, err = tmp_collection_obj.Find(MONGO_CONTEXT, search_filter, find_OPTIONS)
	if err != nil {
		M.Println(" Error in MON SEARCH: ")
		Y.Println(err)
	}

	return tmp_cursor
}

