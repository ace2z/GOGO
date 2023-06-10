/*
Mongolian - Cool Mongo Wrapper using the OFFICIAL Mongo Driver for golang
---------------------------------------------------------------------------------------
This depreciates the old tried and true mgo/mango driver
	v2.00	- 	Jan 11, 2023	- Finally Finished and released!
	v1.23	- 	Nov 05, 2016	-	Initial Release

*/

package CUSTOM_GOMOD

import (
    "context"
    "time"
    "strings"

   . "github.com/ace2z/GOGO/Gadgets"


    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
    "go.mongodb.org/mongo-driver/bson"    

)

/*  REMEMBER!! for GO structs..if you have nested Structs in your struct.. you HAVE TO include the omitempty tag
type GAME_OBJ struct {	
	SUMMARY 		string	`bson:"SUMMARY"`
	DATE 			string	`bson:"DATE"`

	HOME			TEAM_OBJ	`bson:"HOME,omitempty"`
	AWAY			TEAM_OBJ	`bson:"AWAY,omitempty"`

	HOME_BOX_INFO	[]BOX_INFO_OBJ	`bson:"HOME_BOX_INFO,omitempty"`
	AWAY_BOX_INFO	[]BOX_INFO_OBJ	`bson:"AWAY_BOX_INFO,omitempty"`
		
	DATE_OBJ 		time.Time	`bson:"DATE_OBJ"`
}
*/


var MONGO_PORT = "27017"            // defualt mongo port.. can be overrided if you ever need to for some reason

var MONGO_CLIENT *mongo.Client
var MONGO_CLIENT_OPTIONS *options.ClientOptions
var MONGO_CONTEXT context.Context
var MONGO_CANCEL *mongo.Collection

var MONGO_TIMEOUT_SECS = 20
var MON_PREFIX = " --| "
var err error

func TestFUNC() {
    C.Println("MOngolian TESTFUNC")
}

func MONGO_SHOW_DATABASES() {

    //2. Show Databases
    Y.Println(MON_PREFIX, "Listing Databases: ")
    databases, err := MONGO_CLIENT.ListDatabaseNames(MONGO_CONTEXT, bson.M{})
    if err != nil {
        M.Println(err.Error())
        return
    }
    SHOW_STRUCT(databases)

}
func MONGO_INIT(mongo_host string) {

    var mongo_url = "mongodb://" + mongo_host + ":" + MONGO_PORT
    
    Y.Print(MON_PREFIX, "Connecting To MONGO: ")
    W.Print(mongo_url)

    var MONGO_CONTEXT, MONGO_CANCEL = context.WithTimeout(context.Background(), time.Duration(MONGO_TIMEOUT_SECS) * time.Second )
    MONGO_CLIENT_OPTIONS = options.Client().ApplyURI(mongo_url)
    defer MONGO_CANCEL()

    MONGO_CLIENT, err = mongo.Connect(MONGO_CONTEXT, options.Client().ApplyURI(mongo_url))

    if err != nil {
        M.Println(err.Error())
        return
    }
    G.Println(" Success!")

    Y.Print(MON_PREFIX, "Test Ping of Connection:")
    err = MONGO_CLIENT.Ping(MONGO_CONTEXT, readpref.Primary())
    if err != nil {
        M.Println(err.Error())      
        return
    }
    C.Println(" Ping Works!")

//    MONGO_SHOW_DATABASES()
}

/* BROKEN not yet working
var TMP_COMPOUND_INDEXES []string
func NEW_COMPOUND_INDEX(indexname string, itype string, unique bool, PARAMS ...string ) {

    var result = ""
    if itype == "asc" {
        result = indexname + ",asc"

    } else if itype == "desc" {
        result = indexname + ",desc"

    }

    if unique {
        result += ",yes"
    }

    TMP_COMPOUND_INDEXES = append(TMP_COMPOUND_INDEXES, result)
}
*/



var TEMP_INDEXES []string

// Pass INdex Name, the type (asc,desc,text).. and true/false (for unique or not)
func NEW_INDEX(indexname string, itype string, unique bool) {

    var result = ""
    if itype == "asc" {
        result = indexname + "|asc"

    } else if itype == "desc" {
        result = indexname + "|desc"

    } else if itype == "text" {
        result = indexname + "|text"
    }

    if unique {
        result += "|yes"
    }

    TEMP_INDEXES = append(TEMP_INDEXES, result)
}


func CREATE_DATABASE(dbname string, collname string) {

    Y.Println("")
    Y.Print(MON_PREFIX, "Creating DATABASE: ")
    W.Println(dbname)

    //3. Create a database:
    var MONGO_DB_OBJ = MONGO_CLIENT.Database(dbname)
    var MONGO_COLL_OBJ = MONGO_DB_OBJ.Collection(collname)

     // Now create the indexes (the ones that are SINGLE indexes)
    for _, x := range TEMP_INDEXES {

        msplit := strings.Split(x, "|") 

        index_field_name := msplit[0]
        field_type := msplit[1]
        

        uni_flag := false 
        if len(msplit) == 3 {
            unique := msplit[2]
            if unique == "yes" || unique == "true" || strings.Contains(unique, "unique") {
                uni_flag = true
            }
        }

        var indexName = ""
        var err error

        if field_type == "asc" {
            indexName, err = MONGO_COLL_OBJ.Indexes().CreateOne(
                context.Background(),
                mongo.IndexModel{
                        Keys: bson.M{
                            index_field_name: 1,
                        },
                        Options: options.Index().SetUnique(uni_flag),
                },
            )

        } else if field_type == "desc" {
            indexName, err = MONGO_COLL_OBJ.Indexes().CreateOne(
                context.Background(),
                mongo.IndexModel{
                        Keys: bson.M{
                            index_field_name: -1,
                        },
                        Options: options.Index().SetUnique(uni_flag),
                },
            )
        } else if field_type == "text" {
            indexName, err = MONGO_COLL_OBJ.Indexes().CreateOne(
                context.Background(),
                mongo.IndexModel{
                        Keys: bson.M{
                            index_field_name: "text",
                        },
                        Options: options.Index().SetUnique(uni_flag),
                },
            )
        }
    
        if err != nil {
            M.Println(" Index Creation ERROR: ", err.Error())
            return
        }
    
        Y.Print(MON_PREFIX, "Created INDEX: ")
        G.Print(indexName)
        C.Println("", field_type)

    } //end of for

    // Purge the index array  .. ALWAYS DO THIS!!!
    TEMP_INDEXES = []string{}
    
    
}
