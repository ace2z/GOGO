package CUSTOM_GOMOD

import (

//	"io/ioutil"
//	"context"
//	"time"

	"encoding/json"
	
	. "github.com/ace2z/GOGO/Gadgets"

	"github.com/TylerBrock/colorjson"
)



var ALL_JSON_RESULTS = ""
var ALL_COUNT = 0

var JSON_TOP =`
{
`

var JSON_DATA_HEADER =`
  "data" : [
    `

var JSON_BOTTOM =`
  ]
}
`



// This takes any struct and returns "regular" json and pretty colorized JSON
func GEN_PRETTY_JSON(tmpOBJ interface{}) (string, string) {

	// tmp_JSON_OBJ, err := json.Marshal(tmpOBJ)  // Marshall takes a struct and makes it into JSON
	tmp_JSON_OBJ, err := json.MarshalIndent(tmpOBJ, "", "\t")  // Marshall takes a struct and makes it into JSON
	
	if err != nil {
		R.Println(" error in the GEN_PRETTY_JSON ")
		W.Println(err)
		return "", ""
	}
	
	var obj map[string]interface{}
	json.Unmarshal(tmp_JSON_OBJ, &obj)		// Unmarshall takes json and puts it in a struct // marshall does the opposite

	// Marshall the Colorized JSON, Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4
	colorTEMP, _ := f.Marshal(obj)
	pretty_color_JSON := string(colorTEMP)
	regular_JSON := string(tmp_JSON_OBJ)

	return regular_JSON, pretty_color_JSON

} //end of func


func GEN_JSON_HEADER(first_element string) {
	var tmp_header =`
	"` + first_element + `" : [
	  `
	JSON_DATA_HEADER = tmp_header
}





func GENERATE_JSON_Response(DATASET []interface{}) string {
	var TEMP_MIDDLE_JSON = ""
	var ALL_JSON_RESULTS = ""

	var comma_count = 0

	for _, to := range DATASET {
		comma_count++
		pretty_JSON, _ := GEN_PRETTY_JSON(to) 

		TEMP_MIDDLE_JSON = TEMP_MIDDLE_JSON + pretty_JSON
		if comma_count == len(DATASET) {
			break
		}
		// Alwayws add a comma..(unless its the last item)
		TEMP_MIDDLE_JSON = TEMP_MIDDLE_JSON +  `,
       `	
	} //end of for

	ALL_JSON_RESULTS = JSON_TOP + JSON_DATA_HEADER + TEMP_MIDDLE_JSON + JSON_BOTTOM	

	return ALL_JSON_RESULTS
}