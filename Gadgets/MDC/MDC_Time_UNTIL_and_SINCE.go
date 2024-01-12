/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package GOGO_MDC

import (
	"time"
)

// Pass the Future time first and curr Time Second
func helper_GET_SINCE_UNTIL(future time.Time, curr_time time.Time, format string) float64 {

	// Defaults to untiol
	duration := future.Sub(curr_time)

	// Doing a simple loop. We want to make sure result_until is NOT a negative value
	var result_until = duration.Hours()
	for x := 0; x < 5; x++ {
		// Ok, default the result to hours
		result_until = duration.Hours()

		// Days
		if format == "day" || format == "days" {
			result_until = (duration.Hours() / 24)

		} else if format == "min" || format == "minutes" || format == "mins" {
			result_until = duration.Minutes()

		} else if format == "sec" || format == "seconds" || format == "secs" {
			result_until = duration.Seconds()
		}

		result_until = FIX_FLOAT_PRECISION(result_until, 2)

		// If result_until is less than 0.. lets flip the duration comparison
		if result_until < 0.0 {
			duration = curr_time.Sub(future)
		} else {
			break
		}
	}

	return result_until
}

// Pass PREVIOUS time.. then current or alt time) .. also send format: days or mins (defaults to hours)
func GET_Time_SINCE(past time.Time, curr_time time.Time, format string) float64 {
	return helper_GET_SINCE_UNTIL(past, curr_time, format)
}

/*

// Pass PREVIOUS time.. then current or alt time)
func GET_DAYS_SINCE(past time.Time, curr_time time.Time) float64 {
	return helper_GET_SINCE_UNTIL(past, curr_time, "since", "days")
}

// pass future time, then current or alt time
func GET_DAYS_UNTIL(future time.Time, curr_time time.Time) float64 {
	return helper_GET_SINCE_UNTIL(future, curr_time, "until", "days")
}

// pass future time, then current or alt time .. also send format: days or mins (defaults to hours)
func GET_Time_UNTIL(future time.Time, curr_time time.Time, format string) float64 {
	return helper_GET_SINCE_UNTIL(future, curr_time, "until", format)
}
*/
