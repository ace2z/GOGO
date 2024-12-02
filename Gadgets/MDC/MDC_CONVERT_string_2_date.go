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
func CONVERT_DATE_STRING(ALL_PARAMS ...interface{}) (time.Time, string) {

	STRING_input := ""
	output_format_2use := "basic"
	requested_TZ := ""

	TZ_2use := LOCAL_Location_OBJ

	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		// time_val, IS_TIME := param.(time.Time)
		// int_val, IS_INT := param.(int)
		// int64_val, IS_INT64 := param.(int64)

		// First param is ALWAYWs the string input
		if n == 0 && IS_STRING {
			STRING_input = string_val
			continue
		}

		// 2nd param is the output format
		if n == 1 && IS_STRING {
			if string_val != "" {
				output_format_2use = string_val
			}
			continue
		}
		// IF TZ is sent, use it
		if n == 3 && IS_STRING {
			if string_val != "" {
				requested_TZ = string_val
			}

			continue
		}
	}

	//1. Make sure they sent int a supported String format for the date
	isVALID, includes_TIME, pmap := CHECK_for_SUPPORTED_DATE_INPUT(STRING_input)

	if includes_TIME == false {
		output_format_2use = output_format_2use + "reset_time"
	}

	// errro handling
	if isVALID == false {
		M.Print("*** INVALID String Date Format sent to: ")
		Y.Println("CONVERT_DATE")
		Y.Print("Input was: ")
		W.Println(STRING_input)

		DO_EXIT()
		//return "", "", time.Time{}
	}

	// determine the time zone to use (if passed) .. otherwise uses local TZ
	if requested_TZ != "" {
		TZ_2use = GET_PROPER_TZONE_Logic(requested_TZ)
	}

	var num_Mon = pmap["month"].(int)
	var num_Day = pmap["day"].(int)
	var num_Year = pmap["year"].(int)
	var num_Hour = pmap["hour"].(int)
	var num_Min = pmap["min"].(int)
	var num_Sec = pmap["sec"].(int)
	monthObj := time.Month(num_Mon)

	date_OBJ := time.Date(num_Year, monthObj, num_Day, num_Hour, num_Min, num_Sec, 0, TZ_2use)

	// Now... we'll make a pretty date. Based on output format (if sent)
	pretty := OUTPUT_Format_Pretty_Logic(output_format_2use, date_OBJ)

	return date_OBJ, pretty

} //end of func

// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects

} // end of main
