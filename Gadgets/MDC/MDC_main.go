/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)

	
	Aug 28, 2021    v1.23   - Initial Rollout

*/

package GOGO_MDC


import (

	// = = = = = Native Libraries

		"time"
		"strconv"
		"math"
		//"math/rand"

	// = = = = = CUSTOM Libraries

	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/StringOPS"



	// = = = = = 3rd Party Libraries

)


// Makes a floating point number rounded up and returns integer
func MakeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func FIX_FLOAT_PRECISION(num float64, precision int) float64 {

	// Old way.. seems inconsistant sometimes?? 
	//output := math.Pow(10, float64(precision))
	// return float64(MakeRound(num*output)) / output

	percSTRING := strconv.FormatFloat(num, 'f', precision, 64)
	fixed_percNUM, _ := strconv.ParseFloat(percSTRING, 64) 	// this reformats the percentage to have X decimals (based on precisions

	return fixed_percNUM
}


// Easy way to get the UTC form of DATE_OBJ so there is no confusion.. Returns Time, String(pretty date) and Weekday all converted from the orig time
func GET_DB_DATE_UTC(input_DATE_OBJ time.Time) (time.Time, string, string, string) {

	result_DATE_OBJ := input_DATE_OBJ.In(UTC_Location_OBJ)

	pretty, weekday := SHOW_PRETTY_DATE(result_DATE_OBJ)

	result_as_STRING := result_DATE_OBJ.String()

	return result_DATE_OBJ, result_as_STRING, pretty, weekday

}


func CALC_RATIO(smallNUM float64, bigNUM float64, FORCE_DETECT_LARGEST_NUM ...bool) float64 {

	divid := smallNUM / bigNUM
	if FORCE_DETECT_LARGEST_NUM[0] == true {
		if smallNUM > bigNUM {
			divid = bigNUM / smallNUM
		}
	}

	perc := divid * 100.0
	fixed_PERC := FIX_FLOAT_PRECISION(perc, 2)

	return fixed_PERC
} //end of func
	
func GET_RATIO(smallNUM float64, bigNUM float64, FORCE_DETECT_LARGEST_NUM ...bool) float64 {
	return CALC_RATIO(smallNUM, bigNUM, FORCE_DETECT_LARGEST_NUM...)
}




// Takes in Two Time periods.. and returns the duration in DAYS, Hours and Minutes (and comprable strings)
// Returns MINS, HOURS, DAYS (in float first, then strings)
func GET_DURATION(startTIME time.Time, endTIME time.Time, EXTRA_ARGS ...string) (float64, string, string) {
	var precision = 1

	var interval = ""

	//1. First parameter is always the interval. We use this to "force" the value returned
	for n, VAL := range EXTRA_ARGS {

		//1b. If short or full was passed, we format the output date that way
		if n == 0 && VAL != "" {
			interval = VAL
			continue
		}		

		if n == 1 && VAL != "" {			
			precision, _ = strconv.Atoi(VAL)
			continue
		}				
	} //end of for

	temp_secs := endTIME.Sub(startTIME).Seconds()
	temp_mins := endTIME.Sub(startTIME).Minutes()
	temp_hours := endTIME.Sub(startTIME).Hours()	
	temp_Days := temp_hours / 24

	DIFF_SECS := FIX_FLOAT_PRECISION(temp_secs, 1)
	DIFF_MINS := FIX_FLOAT_PRECISION(temp_mins, 1)
	DIFF_Hours := FIX_FLOAT_PRECISION(temp_hours, 1)
	DIFF_Days := FIX_FLOAT_PRECISION(temp_Days, 1)
	
	// TEXT versions:
	sec_text := strconv.FormatFloat(DIFF_SECS, 'f', precision, 64)
	min_text := strconv.FormatFloat(DIFF_MINS, 'f', precision, 64)
	hour_text := strconv.FormatFloat(DIFF_Hours, 'f', precision, 64)
	day_text := strconv.FormatFloat(DIFF_Days, 'f', precision, 64)

	var num_val = 0.0	
	var text_value = ""
	var pretty = ""

	if interval == "hour" || interval == "hours" {
		num_val = DIFF_Hours
		text_value = hour_text
		pretty = hour_text + " Hours"

	} else if interval == "min" || interval == "mins" {
		num_val = DIFF_MINS
		text_value = min_text
		pretty = min_text + " Mins"
	

	} else if interval == "day" || interval == "days" {
		num_val = DIFF_Days
		text_value = day_text
		pretty = day_text + " Days"

	
	} else if interval == "sec" || interval == "secs" {
		num_val = DIFF_SECS
		text_value = sec_text
		pretty = sec_text + " Seconds"

	}
	

	return num_val, text_value, pretty
}


// Tells you if an INT is even 
func IS_EVEN(input_NUM int) bool {

	if input_NUM%2 == 0 {
		return true

	}

	return false
}
// Tells you if an INT is ODD
func IS_ODD(input_NUM int) bool {

	if input_NUM%2 == 0 {

	} else {

		return true
	}

	return false
}





// Kept here as filler/example.. anything you put in this function will start when the module is imported
func init() {

	//1. Startup Stuff (init the command line params etc) . We need these Time ZONE Objects




} // end of main
