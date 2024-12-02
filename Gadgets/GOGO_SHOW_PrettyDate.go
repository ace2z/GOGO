package CUSTOM_GO_MODULE

import (
	"strconv"
	"strings"
	"time"
)

var LOCAL_Location_OBJ, _ = time.LoadLocation("Local")
var EST_Location_OBJ, _ = time.LoadLocation("America/New_York")
var CST_Location_OBJ, _ = time.LoadLocation("America/Chicago")     // aka CST	}
var MST_Location_OBJ, _ = time.LoadLocation("America/Denver")      // MDT / Mountain Standard
var PST_Location_OBJ, _ = time.LoadLocation("America/Los_Angeles") // aka PST
var UTC_Location_OBJ, _ = time.LoadLocation("UTC")

// Entire LIST is: available here:
// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

func Check_for_TZ_shortcut_or_IANA(TZ_to_use string) (bool, *time.Location) {
	var TIMEZONE_OBJ = UTC_Location_OBJ

	is_valid_input := false

	//2. Time Zone logic.. Only supports the majors... Defaults to UTC
	if TZ_to_use != "" {
		tmpLCASE := strings.ToLower(TZ_to_use)
		switch tmpLCASE {
		case "est", "edt":
			TIMEZONE_OBJ = EST_Location_OBJ
			is_valid_input = true
			break

		case "cst", "cdt":
			TIMEZONE_OBJ = CST_Location_OBJ
			is_valid_input = true
			break

		case "mst", "mdt":
			TIMEZONE_OBJ = MST_Location_OBJ
			is_valid_input = true
			break

		case "pst", "pdt":
			TIMEZONE_OBJ = PST_Location_OBJ
			is_valid_input = true
			break

		case "local":
			TIMEZONE_OBJ = LOCAL_Location_OBJ
			is_valid_input = true
			break
		}
	}

	// If input is still FALSE... lets try to see if it is a valid IANA timezone
	if is_valid_input == false {
		new_tz, err := time.LoadLocation(TZ_to_use)
		if err != nil {
			M.Println("Invalid Timezone specified: ", TZ_to_use)
			W.Println("Not valid shortcut or IANA timezone")
			return false, UTC_Location_OBJ
		} else {
			TIMEZONE_OBJ = new_tz
			is_valid_input = true
		}
	}

	return is_valid_input, TIMEZONE_OBJ
}

/*
	   ADD_LEADING_ZERO: This takes in a number and returns a string with a leading 0
	   If the number is already 10 or greater, it returns that same number as is

		SHOW_PRETTY_DATE is dependant on this
*/
func ADD_LEADING_ZERO(myNum int) string {

	RESULT := strconv.Itoa(myNum)

	if myNum <= 9 {
		RESULT = "0" + RESULT
	}

	return RESULT
}

/* SHOW_PRETTY_DATE Takes in a time.Time DATE_OBJ and returns a PRETTY formatted based on what you specify
   Acceot: basic, simple, full, justdate, justtime, timestamp, british, nano, weekday

   For FORMAT specifiy: basic, simple, full, nano, british, justtime, justdate, timestamp
   You can also modify format by adding:
   _noday   (ie, basic_noweek) - Prevents the weekday info from showing
   _nozone   		- prevents the timezone info from showing
   _reset_time 		- For situations where you want to omit the HH:MM cause you dont need it...resets time to 00:00

*/

// func SHOW_PRETTY_DATE(input_DATE time.Time, EXTRA_ARGS...string) string {
func SHOW_PRETTY_DATE(ALL_PARAMS ...interface{}) string {
	var output_FORMAT = "basic"

	convert_TIME := false
	convert_EPOCH := false
	convert_64EPOCH := false

	input_DATE := time.Now()
	EPOCH_input := 0
	var EPOCH64_input int64 = 0

	TZ_2use := LOCAL_Location_OBJ
	requested_TZ := ""

	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		time_val, IS_TIME := param.(time.Time)
		int_val, IS_INT := param.(int)
		int64_val, IS_INT64 := param.(int64)

		// First parama is always a time.Time DATE_OBJ or an int epoch
		// Figure out if we are converting a STRING to a DATE_OBJ... or a DATE_OBJ to a string
		if n == 0 {
			if IS_TIME {
				input_DATE = time_val
				convert_TIME = true

			} else if IS_INT {
				EPOCH_input = int_val
				convert_EPOCH = true

			} else if IS_INT64 {
				EPOCH64_input = int64_val
				convert_64EPOCH = true

			} else {
				M.Print(" Invalid input sent for DATE!!!")
				Y.Println(" Need either time.TIME DateOBJ ... or EPOCH as int or int64")
				DO_EXIT()
			}
			continue
		}

		//2nd param is always going to be what we need for the output format
		if n == 1 {
			if IS_STRING {
				if string_val != "" {
					output_FORMAT = string_val
				}
			} else {
				M.Print(" Invalid input for OUTPUT_FORMAT")
				Y.Println(" Need a STRING seperated by underscores")
				DO_EXIT()
			}
			continue
		}

		//3rd param is always going to be the timezone ... it accepts the basic ones like EST, CST, MST, PST
		// also accepts anything valid from the IANA TZ database
		if n == 2 {
			if IS_STRING {
				requested_TZ = string_val

			} else {
				M.Println("Invalid input TZ to use! Try est, cst, pst, UTC, GMT... or anything from the IANA Database")
				W.Println("https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
				DO_EXIT()
			}
			continue
		}

	} // end of input forloop

	// First determine what input type we are using.. If its a time.Time or an EPOCH
	// We Default to converting to the time.. So if its EPOCH.. this is what we do
	if convert_TIME == false {
		if convert_EPOCH {
			s64 := int64(EPOCH_input)
			input_DATE = time.Unix(s64, 0)

		} else if convert_64EPOCH {
			input_DATE = time.Unix(EPOCH64_input, 0)
		}
	}

	// Now.. determine if we need to convert this timeObject to a NEW time zone
	if requested_TZ != "" {
		valid_shortcut, NEW_TZ_OBJ := Check_for_TZ_shortcut_or_IANA(requested_TZ)
		if valid_shortcut == false {
			DO_EXIT()
		} else {
			TZ_2use = NEW_TZ_OBJ
		}

	}
	// finally.. convert the time input to the NEW timeon object
	tmptime := input_DATE.In(TZ_2use)
	input_DATE = tmptime

	// //1. Parse out EXTRA_ARGS
	// for _, VAL := range EXTRA_ARGS {

	// 	//1e. only parameter this takes is the output format we want
	// 	// If full is passed, we show this format: Wednesday, 11/20/2020 @ 13:56
	// 	// if british or iso is passed, we show: 2015-05-30
	// 	if VAL != "" {
	// 		output_FORMAT = VAL
	// 		continue
	// 	}

	// } // end of for

	//1c. Here is the default DELIMITER we use (can be overridden by hyphen or underscore)
	delim := "/"

	if strings.Contains(output_FORMAT, "hyphen") {
		delim = "-"
	} else if strings.Contains(output_FORMAT, "underscore") {
		delim = "_"
	}

	//2. From this object, extract the M/D/Y HH:MM
	montemp := int(input_DATE.Month())
	daytemp := input_DATE.Day()

	hourtemp := input_DATE.Hour()
	mintemp := input_DATE.Minute()

	//3. Then, we add leading 0's as needed
	cMon := ADD_LEADING_ZERO(montemp)
	cDay := ADD_LEADING_ZERO(daytemp)
	cHour := ADD_LEADING_ZERO(hourtemp)
	cMin := ADD_LEADING_ZERO(mintemp)

	sectemp := input_DATE.Second()
	cSec := ADD_LEADING_ZERO(sectemp)
	nanotemp := input_DATE.Nanosecond()

	cNanoSecond := strconv.Itoa(nanotemp)

	//4. Thankfully we dont have to worry about this fuckery with the year!
	cYear := strconv.Itoa(input_DATE.Year())
	weekd := input_DATE.Weekday().String()

	//5. Update ZONE info
	tmp_zone, tmp_offset := input_DATE.Zone()

	tmp_off_string := strconv.Itoa(tmp_offset)

	TMP_ZONE_FULL := "(" + tmp_zone + " " + tmp_off_string + ")"

	/* 7. Here is the DEFAULT Pretty format that is returned

	09/26/1978 @ 13:58

		or (if SHOW_SECONDS is passed)

	09/26/1978 @ 13:58:05
	*/
	//result_TEXT := cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin
	result_TEXT := ""

	ADD_ZONE := true
	ADD_WEEKDAY := true
	need_just_weekday := false

	if strings.Contains(output_FORMAT, "nozone") {
		ADD_ZONE = false
	}

	if strings.Contains(output_FORMAT, "noweekday") || strings.Contains(output_FORMAT, "noweek") || strings.Contains(output_FORMAT, "noday") || strings.Contains(output_FORMAT, "nodow") {
		ADD_WEEKDAY = false
	}

	if strings.Contains(output_FORMAT, "reset_time") {
		cHour = "00"
		cMin = "00"
		cSec = "00"
	}

	if strings.Contains(output_FORMAT, "shortweek") || strings.Contains(output_FORMAT, "shortdow") {
		weekd = weekd[0:3]
	}

	//if we want a two digit year
	if strings.Contains(output_FORMAT, "twoyear") || strings.Contains(output_FORMAT, "shortyear") {
		cYear = cYear[len(cYear)-2:]
	}

	//8. BASIC Format just returns the following.. no weekday or Tzone
	if strings.Contains(output_FORMAT, "basic") || strings.Contains(output_FORMAT, "simple") {

		result_TEXT = cMon + delim + cDay + delim + cYear + " @ " + cHour + ":" + cMin

		// Basic format is just that.. no zone no weekday
		ADD_WEEKDAY = false
		//		ADD_ZONE = false

		//9. FULL Format: //Wednesday, 11/20/2020 @ 13:56 EST (-5 Hours)
	} else if strings.Contains(output_FORMAT, "full") {

		result_TEXT = cMon + delim + cDay + delim + cYear + " @ " + cHour + ":" + cMin + ":" + cSec

	} else if strings.Contains(output_FORMAT, "nano") {
		result_TEXT = cMon + delim + cDay + delim + cYear + " @ " + cHour + ":" + cMin + ":" + cSec + ":n:" + cNanoSecond

		//10. This is the british/iso format: 2020-09-26
	} else if strings.Contains(output_FORMAT, "british") || strings.Contains(output_FORMAT, "iso") {

		result_TEXT = cYear + delim + cMon + delim + cDay

		//11. This is JUSTDATE:  09/26/1988
	} else if strings.Contains(output_FORMAT, "justdate") {

		result_TEXT = cMon + "/" + cDay + "/" + cYear

	} else if strings.Contains(output_FORMAT, "justtime") {

		result_TEXT = cHour + ":" + cMin

		//12. For use as a simple timestamp for a file suffix using _ underscores
	} else if strings.Contains(output_FORMAT, "timestamp") {

		result_TEXT = cMon + "_" + cDay + "_" + cYear + "_" + cHour + "_" + cMin

		//13. If we JUST want the weekday
	} else if strings.Contains(output_FORMAT, "weekday") || strings.Contains(output_FORMAT, "dow") || strings.Contains(output_FORMAT, "justweek") {
		result_TEXT = weekd
		need_just_weekday = true
	}

	if ADD_WEEKDAY && need_just_weekday == false {
		result_TEXT = weekd + ", " + result_TEXT
	}

	// default to adding the WEEKDAY and ZONE suffix... if _nozone or _noday was specified, this gets ignored
	if ADD_ZONE {
		result_TEXT = result_TEXT + " " + TMP_ZONE_FULL
	}

	// error handling
	if result_TEXT == "" {

		R.Println(" ERROR in SHOW_PRETTY_DATE: .. invalid output_FORMAT sent!!!")
		DO_EXIT()
	}

	//12. As a bonus, we always return the weekday as a second variable
	return result_TEXT
} //end of func

// func SHOW_PRETTY_DATE(ALL_PARAMS ...interface{}) string {
// alias for SHOW_PRETTY_DATE
func PRETTY_DATE(ALL_PARAMS ...interface{}) string {
	return SHOW_PRETTY_DATE(ALL_PARAMS...)
}
func PRETTY_TIME(ALL_PARAMS ...interface{}) string {
	return SHOW_PRETTY_DATE(ALL_PARAMS...)
}
