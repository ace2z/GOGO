package CUSTOM_GOMOD

import (

//	"io/ioutil"
//	"context"
//	"time"


	. "github.com/ace2z/GOGO/Gadgets"
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