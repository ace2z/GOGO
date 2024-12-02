/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (

	// = = = = = Native Libraries
	"strconv"
	//"math/rand"
	// = = = = = CUSTOM Libraries
	//. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/StringOPS"
	// = = = = = 3rd Party Libraries
)

func CHECK_for_SUPPORTED_DATE_INPUT(inputDate string) (bool, bool, map[string]interface{}) {

	var res_MAP map[string]interface{}

	var includes_TIME = false

	//1. Otherwise.. lets determine which input format this is
	// add new formats here as needed
	var final_map map[string]interface{}

	// add new formats here as needed
	//is_EPOCH, emap := have_EPOCH_FORMAT(inputDate)
	is_valid_FORMAT, timefound, dmap := have_SUPPORTED_DEFAULT_FORMAT(inputDate)

	if is_valid_FORMAT {
		final_map = dmap

		// Otherwise, we dont have a supported format
	} else {
		return false, false, res_MAP
	}
	includes_TIME = timefound

	//5. First cast to a string from the map
	var sMON = final_map["month"].(string)
	var sDAY = final_map["day"].(string)
	var sYEAR = final_map["year"].(string)
	var sHOUR = final_map["hour"].(string)
	var sMIN = final_map["min"].(string)
	var sSEC = final_map["sec"].(string)

	//3. Then cast to INT so we can give it to time.Date
	var num_Mon, _ = strconv.Atoi(sMON)
	var num_Day, _ = strconv.Atoi(sDAY)
	var num_Year, _ = strconv.Atoi(sYEAR)
	var num_Hour, _ = strconv.Atoi(sHOUR)
	var num_Min, _ = strconv.Atoi(sMIN)
	var num_Sec, _ = strconv.Atoi(sSEC)

	res_MAP = map[string]interface{}{
		"month": num_Mon,
		"day":   num_Day,
		"year":  num_Year,
		"hour":  num_Hour,
		"min":   num_Min,
		"sec":   num_Sec,
	}
	return true, includes_TIME, res_MAP
}
