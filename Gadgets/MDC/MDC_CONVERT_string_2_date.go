/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (

	// = = = = = Native Libraries

	"time"

	//"math/rand"

	// = = = = = CUSTOM Libraries

	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/StringOPS"
	// = = = = = 3rd Party Libraries
)

/*
Converts the following
  - a STRING date (in the proper format) to a time.Time DATE_OBJ

STRING format for the Date must be in one of the following or you will error:

  - MM-DD-YYYY

  - 2024-03-15 18:30:00		( ISO Full British Format )

  - YYYY-MM-DD		(ISO / British format)

  - MM/DD/YYYY

  - YYYY/MM/DD

    Also accepts TIME.. Which must be apppended as:  (24 hour format only supported)

  - XXXXX_18:05

  - XXXXX@18:05

    // USES THE SAME OUTPUT FORMATS and ZONE modifiers as the SHOW_PRETTY_DATE function
*/
func CONVERT_DATE_STRING(ALL_PARAMS ...interface{}) time.Time {

	STRING_input := ""
	OUTPUT_format := "basic"		

	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		time_val, IS_TIME := param.(time.Time)
		int_val, IS_INT := param.(int)
		int64_val, IS_INT64 := param.(int64)

		// First param is ALWAYWs the string input
		if n == 0 && IS_STRING {
			STRING_input = string_val
			continue
		}

		// 2nd param is the output format
		if n == 1 && IS_STRING {



		
	}

	//1. Remove all spaces in this string just in case
	isVALID, pmap := CHECK_for_SUPPORTED_DATE_INPUT(STRING_input)
	// errro handling
	if isVALID == false {
		M.Print("*** INVALID Date Format sent to: ")
		Y.Println("CONVERT_DATE")

		Y.Print("Input was: ")
		W.Println(STRING_input)

		DO_EXIT()
		//return "", "", time.Time{}
	}

	var num_Mon = pmap["month"].(int)
	var num_Day = pmap["day"].(int)
	var num_Year = pmap["year"].(int)
	var num_Hour = pmap["hour"].(int)
	var num_Min = pmap["min"].(int)
	var num_Sec = pmap["sec"].(int)
	monthObj := time.Month(num_Mon)

	var LOCAL_Location_OBJ, _ = time.LoadLocation("Local")

	date_OBJ := time.Date(num_Year, monthObj, num_Day, num_Hour, num_Min, num_Sec, 0, LOCAL_Location_OBJ)

	return date_OBJ

} //end of func

// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects

} // end of main
