package CUSTOM_GOMOD

import (

	// For MOnGO using Latest Driver

	. "github.com/ace2z/GOGO/Gadgets"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Inserts a record if it does NOT exist... otherwise, Updates an EXISTING record based on filter..
func DO_UPSERT(dbname string, collname string, filter interface{}, RECORD interface{}) {
	C.Println("")
	C.Println(" = =| Performing UPSERT on single record")

	var MONGO_DB_OBJ = MONGO_CLIENT.Database(dbname)
	var MONGO_COLL_OBJ = MONGO_DB_OBJ.Collection(collname)
	var opts = options.Update().SetUpsert(true)

	// Updates tthe whole record if it doesnt exist
	updateDATA := bson.D{{"$set",
		RECORD,
		// bson.D{
		//     {"", STO},
		// },
	}}

	_, err2 := MONGO_COLL_OBJ.UpdateOne(MONGO_CONTEXT, filter, updateDATA, opts)

	if err2 != nil {
		M.Println(" ERROR in UPSERT: ")
		Y.Println(err2.Error())

	} else {
		G.Println(" = =| Success!")
		G.Println("")
	}
}
