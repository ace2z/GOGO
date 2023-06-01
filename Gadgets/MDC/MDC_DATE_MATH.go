package GOGO_MDC

import (
		"time"
		"strings"

		. "github.com/ace2z/GOGO/Gadgets"
)




// Easy way to find if FIRST date is AFTER the PREV DATE
func DATE_IS_AFTER(first, prev time.Time) bool {
    return first.After(prev)
}

// Correspondingly a date that is BEFORE
func DATE_IS_BEFORE(first, prev time.Time) bool {
    return first.Before(prev)
}


// Gets the difference between two dates (by days, hour or minutes)
func GET_DATE_DIFF(mtype string, currDATE time.Time, prevDATE time.Time) int {

	if strings.Contains(mtype, "day") {

		days := currDATE.Sub(prevDATE).Hours() / 24	
		return int(days)
	
	} else if strings.Contains(mtype, "hour") {

		delta := currDATE.Sub(prevDATE)
		result := int(delta.Hours())

		return result
		
	} else if strings.Contains(mtype, "min") {
		delta := currDATE.Sub(prevDATE)
		result := int(delta.Minutes())

		return result
	}

	return 0
}



/* Takes in two date objects and returns the TIME DIFFERNCE between them in the 5m40s format
 */
 func DISPLAY_TIME_DIFF(startTime time.Time, endTime time.Time) string {
	
	diff := endTime.Sub(startTime)
	return diff.String()
}

// Alias for DISPLAY_TIME_DIFF (which lives in GO_GO_Gadgets)
func GET_TIME_DIFF(startTime time.Time, endTime time.Time) string {
	return DISPLAY_TIME_DIFF(startTime, endTime)
}



/* Takes in a date object and adds or subtracts
based on the number and whatever operation you specify
returns a date object
*/
func DateMath(dateObj time.Time, operation string, v_amount int, interval string) (string, time.Time) {

	//dateObj = dateObj.UTC()


	//1. If we are subtracting, we change amount to a negative number/// otherwise we default to adding
	if operation == "sub" || operation == "subtract" {

		v_amount = -v_amount

	}

	//2. Now we do the add or subtract operattion based on the time.Duration that is interval
	// Default is minute

	timeINT := time.Minute

	if interval == "hour" || interval == "hours" {

		timeINT = time.Hour

	} else if interval == "min" || interval == "mins" || interval == "minute" || interval == "minutes" {

		timeINT = time.Minute

	} else if interval == "sec" || interval == "secs" || interval == "second" || interval == "seconds" {

		timeINT = time.Second

	} else if interval == "day" || interval == "days" {

		timeINT = (time.Hour * 24)

	}


	//3. Finally do the "date math" on the incoming dateObj
	result_DATE_OBJ := dateObj.Add(time.Duration(v_amount) * timeINT)
	prettyDATE, _ := SHOW_PRETTY_DATE(result_DATE_OBJ)

	return prettyDATE, result_DATE_OBJ

} //end of dateMath
