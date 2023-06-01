package CUSTOM_GO_MODULE

import (
	"time"
	"strconv"
	"strings"
	"os"

)


/* SHOW_PRETTY_DATE Takes in a time.Time DATE_OBJ and returns a PRETTY formatted based on what you specify

   For FORMAT specifiy: basic, simple, full, nano, british, justtime, justdate, timestamp
   You can also modify format by adding: 
   _noday   (ie, basic_noweek) - Prevents the weekday info from showing
   _nozone   		- prevents the timezone info from showing
   _reset_time 		- For situations where you want to omit the HH:MM cause you dont need it...resets time to 00:00

*/
func SHOW_PRETTY_DATE(input_DATE time.Time, EXTRA_ARGS...string) (string, string) {
	var output_FORMAT = "basic"

	//1. Parse out EXTRA_ARGS
	for _, VAL := range EXTRA_ARGS {


		//1e. only parameter this takes is the output format we want
		// If full is passed, we show this format: Wednesday, 11/20/2020 @ 13:56
		// if british or iso is passed, we show: 2015-05-30
		if VAL != "" {
			output_FORMAT = VAL
			continue
		}

	} // end of for

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
	result_TEXT := "nullDATE_specified"
	

	ADD_ZONE := true
	ADD_WEEKDAY := true
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


	//8. BASIC Format just returns the following.. no weekday or Tzone
	if strings.Contains(output_FORMAT, "basic") || strings.Contains(output_FORMAT, "simple") {	

		result_TEXT = cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin

		// Basic format is just that.. no zone no weekday
		if strings.Contains(output_FORMAT, "basic") {
			ADD_WEEKDAY = false
			ADD_ZONE = false
		}

	//9. FULL Format: //Wednesday, 11/20/2020 @ 13:56 EST (-5 Hours)
	} else if strings.Contains(output_FORMAT, "full") {
		
		result_TEXT = cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin + ":" + cSec
	
	} else if strings.Contains(output_FORMAT, "nano") {
		result_TEXT = cMon + "/" + cDay + "/" + cYear + " @ " + cHour + ":" + cMin + ":" + cSec + ":n:" + cNanoSecond
	
	//10. This is the british/iso format: 2020-09-26
	} else if strings.Contains(output_FORMAT, "british") {

		result_TEXT = cYear + "-" + cMon + "-" + cDay

	//11. This is JUSTDATE:  09/26/1988
	} else if strings.Contains(output_FORMAT, "justdate") {

		result_TEXT = cMon + "/" + cDay + "/" + cYear

	} else if strings.Contains(output_FORMAT, "justtime") {

		result_TEXT = cHour + ":" + cMin
	
	//12. For use as a simple timestamp for a file suffix using _ underscores
	} else if strings.Contains(output_FORMAT, "timestamp"){
	
		result_TEXT = cMon + "_" + cDay + "_" + cYear + "_" + cHour + "_" + cMin
	

	} else {

		R.Println(" ERROR in SHOW_PRETTY_DATE: .. invalid output_FORMAT sent!!!")
		os.Exit(-9)
	}


	// default to adding the WEEKDAY and ZONE suffix... if _nozone or _noday was specified, this gets ignored
	if ADD_ZONE {
		result_TEXT = result_TEXT + " " + TMP_ZONE_FULL
	}
	if ADD_WEEKDAY {

		if strings.Contains(output_FORMAT, "shortweek") || strings.Contains(output_FORMAT, "shortdow") {
			weekd = weekd[0:3]
		}

		result_TEXT = weekd + ", " + result_TEXT
	}



	//12. As a bonus, we always return the weekday as a second variable

	return result_TEXT, weekd
} //end of func

/*
   ADD_LEADING_ZERO: This takes in a number and returns a string with a leading 0
   If the number is already 10 or greater, it returns that same number as is
 
	SHOW_PRETTY_DATE is dependant on this
*/
func ADD_LEADING_ZERO( myNum int) string {

	RESULT := strconv.Itoa(myNum)

	if myNum <= 9 {
		RESULT = "0" + RESULT
	}

	return RESULT
}