/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package GOGO_MDC

import (

	// = = = = = Native Libraries
	"strings"
	"time"

	//"math/rand"

	// = = = = = CUSTOM Libraries
	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/StringOPS"
	// = = = = = 3rd Party Libraries
)

// Use ful Time zone locations
var LOCAL_Location_OBJ, _ = time.LoadLocation("Local")
var EST_Location_OBJ, _ = time.LoadLocation("America/New_York")
var CST_Location_OBJ, _ = time.LoadLocation("America/Chicago")     // aka CST	}
var MST_Location_OBJ, _ = time.LoadLocation("America/Denver")      // MDT / Mountain Standard
var PST_Location_OBJ, _ = time.LoadLocation("America/Los_Angeles") // aka PST
var UTC_Location_OBJ, _ = time.LoadLocation("UTC")

/*
Converts the following
  - a STRING date (in the proper format) to a time.Time DATE_OBJ
  - a DATE_OBJ into a 'pretty date'
  - an EPOCH ..  into a pretty date.. and time.Time Date_OBJ
    (specify this with int(EPOCH_wasint64)  .. this is because Go doesnt recognize int64 as a param explicitly

TZONE: Specify cst, est, mdt or pst if you need to override the timezone format returned

STRING format for the Date must be in one of the following or you will error:

  - MM-DD-YYYY

  - 2024-03-15 18:30:00		( ISO Full British Format )

  - YYYY-MM-DD		(ISO / British format)

  - MM/DD/YYYY

  - YYYY/MM/DD

    Also accepts TIME.. Which must be apppended as:  (24 hour format only supported)

  - XXXXX_18:05

  - XXXXX@18:05

Final param is for FORMAT specifiy: basic, simple, full, nano, british, justtime, justdate, timestamp
(this gets passed to SHOW_PRETTY_DATE )
You can also add any of the following to modify output_format:

_noday   (ie, basic_noweek) - Prevents the weekday info from showing
_nozone   		- prevents the timezone info from showing
_reset_time 	- For situations where you want to omit the HH:MM cause you dont need it...resets time to 00:00

	Major BUGFIX done on 4/14/2024
*/
func CONVERT_DATE(ALL_PARAMS ...interface{}) (string, string, time.Time) {

	var STRING_input = ""
	var DATE_input time.Time
	var EPOCH_input int

	var output_FORMAT = ""
	var TZ_2_USE = UTC_Location_OBJ

	var need_DATE_convert = false
	var need_STRING_convert = false
	var need_EPOCH_convert = false

	//1. First parameter is always the input date in the proper format
	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		time_val, IS_TIME := param.(time.Time)
		int_val, IS_INT := param.(int)

		// Figure out if we are converting a STRING to a DATE_OBJ... or a DATE_OBJ to a string
		if n == 0 {
			if IS_STRING {
				STRING_input = string_val
				need_STRING_convert = true

			} else if IS_TIME {
				DATE_input = time_val
				need_DATE_convert = true

			} else if IS_INT {
				EPOCH_input = int_val
				need_EPOCH_convert = true
			} else {
				M.Print(" Invalid input sent to CONVERT_DATE!!!")
				Y.Println(" Need STRING, time.Time or EPOCH/INT64")
				Y.Print("Input was: ")
				W.Println(param)
				DO_EXIT()
			}
			continue
		}

		//2. Next param is the output format to use ... ie (short, basic, simple, nano etc)
		// You can also tack on _est, _mst, _edt to return the date in THAT TZ
		if n == 1 {
			if IS_STRING {
				output_FORMAT = string_val
				continue
			}
			continue
		}

		//3. Third param can be TZ time zone.. so you get the date back in est/cst/mdt
		if n == 2 {
			if IS_STRING {
				is_supported, tmp_obj := GET_TZ_conv_OBJECT(string_val)
				if is_supported {
					TZ_2_USE = tmp_obj

				} else {
					M.Println("Invalid TZ Sent to CONVERT_DATE")
					Y.Println(string_val)
					DO_EXIT()
				}
			}
			continue
		}

	} //end of for

	if need_STRING_convert {
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

		// if reset time was passed in output format
		if strings.Contains(output_FORMAT, "reset_time") {
			num_Hour = 0
			num_Min = 0
			num_Sec = 0
		}

		date_OBJ := time.Date(num_Year, monthObj, num_Day, num_Hour, num_Min, num_Sec, 0, TZ_2_USE)

		//12. Now pass to show_Pretty_Date with the output format if specified
		OUTPUT, weekday := SHOW_PRETTY_DATE(date_OBJ, output_FORMAT)

		return OUTPUT, weekday, date_OBJ

		//5. If were converting a date_OBJ to a string
	} else if need_DATE_convert {

		// if reset time was passed in output format
		if strings.Contains(output_FORMAT, "reset_time") {

			num_MON_OBJ := DATE_input.Month()
			num_DAY := DATE_input.Day()
			num_YEAR := DATE_input.Year()
			num_Hour := 0
			num_Min := 0
			num_Sec := 0

			DATE_input = time.Date(num_YEAR, num_MON_OBJ, num_DAY, num_Hour, num_Min, num_Sec, 0, TZ_2_USE)
		}

		OUTPUT, weekday := SHOW_PRETTY_DATE(DATE_input, output_FORMAT)
		return OUTPUT, weekday, DATE_input

		//6. If they passed an EPOCH Int64...
	} else if need_EPOCH_convert {

		s64 := int64(EPOCH_input)
		date_OBJ := time.Unix(s64, 0)
		date_OBJ = date_OBJ.In(TZ_2_USE)

		OUTPUT, weekday := SHOW_PRETTY_DATE(date_OBJ, output_FORMAT)
		return OUTPUT, weekday, date_OBJ

	} else {
		M.Print("*** INVALID Date Format Specified!!! ")
		Y.Println(" to CONVERT_DATE")
		DO_EXIT()
	}

	// Default return... for errors or something we dont support
	return "", "", time.Time{}
} //end of func

// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects

} // end of main
