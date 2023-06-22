package CUSTOM_GOMOD

import (

	// For MOnGO using Latest Driver
	//"go.mongodb.org/mongo-driver/mongo/options"

	. "github.com/ace2z/GOGO/Gadgets"
)


func DO_BULK_DELETE(dbname string, coll_name string, search_filter interface{} ) {


    C.Println("")
    C.Print(" = =| Attempting BULK DELETE/PURGE using: ")
    Y.Println(search_filter)
        
	var MONGO_DB_OBJ = MONGO_CLIENT.Database(dbname)
	var MONGO_COLL_OBJ = MONGO_DB_OBJ.Collection(coll_name)
	//var opts = options.DeleteOptions.SetBypassDocumentValidation(true)

	_, err2 := MONGO_COLL_OBJ.DeleteMany(MONGO_CONTEXT, search_filter)

	if err2 != nil {
		M.Println("BULK DELETE ERROR: ", err2.Error())
		return
	} else {
        G.Println(" = =| BULK DELETE Success!")
        G.Println("")
    }

    // Always pause for a second after deleting.
    //Sleep(1, false)
}
