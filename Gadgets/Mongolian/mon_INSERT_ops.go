package CUSTOM_GOMOD

import (

    // For MOnGO using Latest Driver    
    "go.mongodb.org/mongo-driver/mongo/options"

    . "github.com/ace2z/GOGO/Gadgets"
)


// Inserts a SINGLE record
func NEW_INSERT(dbname string, collname string, data interface{} ) {
    C.Println("")
    C.Println(" = =| Performing SINGLE record INSERT")

    var MONGO_DB_OBJ = MONGO_CLIENT.Database(dbname)
    var MONGO_COLL_OBJ = MONGO_DB_OBJ.Collection(collname)
    var opts = options.InsertOne().SetBypassDocumentValidation(true)

    _, err2 := MONGO_COLL_OBJ.InsertOne(MONGO_CONTEXT, data, opts)    
    if err2 != nil {
        M.Println(" ERROR in NEW_INSERT: ")
        Y.Println(err2.Error())
     
    } else {
        G.Println(" = =| Success!")
        G.Println("")
    }
}


func SHOW_BULK_HEADER(EXTRA_ARGS...string) {
	Y.Print(" = =| BULK Save for: ")	

	for n, VAL := range EXTRA_ARGS {
		if n == 0 {
			W.Print(VAL + " ")
			continue
		}
		C.Print(VAL + " ")
	}
	W.Println("")
}
func DO_BULK_INSERT(dbname string, coll_name string, DATA []interface{}) {


    C.Println("")
    C.Print(" = =| Attempting BULK INSERT of ")
    G.Print(len(DATA))
    W.Print(" NEW ")
    C.Println("items..")

    if len(DATA) <= 0 {
        Y.Println(" = =| Nothing new to save!")
        return
    }    
	var MONGO_DB_OBJ = MONGO_CLIENT.Database(dbname)
	var MONGO_COLL_OBJ = MONGO_DB_OBJ.Collection(coll_name)
	var opts = options.InsertMany().SetBypassDocumentValidation(true)
	_, err2 := MONGO_COLL_OBJ.InsertMany(MONGO_CONTEXT, DATA, opts)

	if err2 != nil {
		M.Println("BULK INSERT Error: ", err2.Error())
		return
	} else {
        G.Println(" = =| Success!")
        G.Println("")
    }
}
