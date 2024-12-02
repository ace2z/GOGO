package CUSTOM_GO_MODULE

import (
	//"strings"
	"math"
	//"strconv"

	. "github.com/ace2z/GOGO/Gadgets"
)

// Returns Percentages INCREASE DECREASE for stocks etc... Takes in floats or INTs
/*
	As of MAY 2023.. This is the ULTIMATE GET PERCENTAGE INC/DEC function
	Also returns the DIFF between two numbers and if it was an INCREASE or DECERASE based on which number was passed FIRST

	RETURNS:
	res_PERC, changeTYPE, res_DESC, res_diff
*/
func GET_INCDEC_PERCENT(ALL_PARAMS ...interface{}) (float64, string, string, float64) {

	var firstNUM = 0.0
	var secNUM = 0.0
	var SHOW_OUTPUT = false

	var res_diff = 0.0
	var res_DESC = "0.0% No Change"
	var res_PERC = 0.0

	var precision = 1

	// If true is passed, we show the output of this function
	for n, param := range ALL_PARAMS {
		intval, IS_INT := param.(int)
		floatval, IS_FLOAT := param.(float64)
		boolval, IS_BOOL := param.(bool)

		// First paramn is always FIRSTNUM
		if n == 0 || n == 1 {

			if IS_INT {

				if n == 0 {
					firstNUM = float64(intval)
				} else if n == 1 {
					secNUM = float64(intval)
				}

			} else if IS_FLOAT {
				if n == 0 {
					firstNUM = floatval
				} else if n == 1 {
					secNUM = floatval
				}

				// Hard exit for error handling
			} else {
				M.Println(" ERROR: Invalid values setn to GET_DIFF")
				Y.Println(" Must be INT or FLOAT")
				DOEXIT()
			}
			continue
		}

		// If an integer was passed (that isnt first and secnum) assume this is precision
		if IS_INT {

			precision = intval
			continue
		}

		// if bool was passed.. assume we are SHOWING the output
		if IS_BOOL {
			SHOW_OUTPUT = boolval
			continue
		}
	}

	// ERROR HANLDING
	//1. First get the diff... if they are equal, we return
	if firstNUM == secNUM {

		return res_PERC, "no_change", res_DESC, res_diff
	}
	//2. This is for convenience and makes it easier to remember what number is what (large or small)
	var smallNUM = firstNUM
	var largeNUM = secNUM
	var mode = "INCREASE"

	if firstNUM > secNUM {
		mode = "DECREASE"
	}

	var ERROR_HANDLE_OVERRIDE_FLAG = false
	// ERROR HANLDING for 0.0
	//3. Also if one number is 0.0 .. we return.. obviously this is 100%
	res_PERC = 100.0
	res_DESC = "100% ( " + mode + " from zero ) "

	if smallNUM == 0.0 {
		res_diff = largeNUM
		ERROR_HANDLE_OVERRIDE_FLAG = true
		//return res_DESC, res_PERC, res_diff

	} else if largeNUM == 0.0 {
		res_diff = smallNUM
		ERROR_HANDLE_OVERRIDE_FLAG = true
		//return res_DESC, res_PERC, res_diff
	}

	if ERROR_HANDLE_OVERRIDE_FLAG == false {
		// ERROR HANDLING for negative numbers
		// If both numbers are NEGATIVE.. flip them to positive
		if smallNUM < 0.0 && largeNUM < 0.0 {
			smallNUM = math.Abs(smallNUM)
			largeNUM = math.Abs(largeNUM)

		}

		//4. Now that we have the numbers in the correct place, lets do the math
		// CORRECT PERCENTAGE CALCULATION: as of May 2023.. checked on percentage calculator as well
		tdiff := largeNUM - smallNUM
		div_result := tdiff / smallNUM
		var mperc = div_result * 100

		//4b If negative number... we flip it to positive..
		if mperc < 0.0 {
			mperc = math.Abs(mperc)
		}

		//5. Convert the percentages into readable objects
		res_PERC = FIX_FLOAT_PRECISION(mperc, precision)

		// Now, lets adjust percSTRING so we have a leading % char
		// (must be done AFTER the call to parseFLoat)
		ptemp := FLOAT_to_STRING(res_PERC)
		res_DESC = mode + " by " + ptemp

		//6. Return the DIFFERENCE between the two numbers
		res_diff = FIX_FLOAT_PRECISION(tdiff, precision)

	}

	//6b fix for res_diff
	if res_diff < 0.0 {
		res_diff = math.Abs(res_diff)
	}
	//7. If we are showing the output
	if SHOW_OUTPUT {
		firstSTRING := FLOAT_to_STRING(firstNUM)
		secSTRING := FLOAT_to_STRING(secNUM)
		diffstring := FLOAT_to_STRING(res_diff)
		percSTRING := FLOAT_to_STRING(res_PERC) + "%"
		SHOW_BOX("GET_PERCENTAGE", "|cyan|"+firstSTRING+" --> "+secSTRING, "|yellow|"+mode+" by", "|green|"+percSTRING, "DIFF: ", "|yellow|"+diffstring)
	}

	var changeTYPE = mode

	// Finally return everything
	return res_PERC, changeTYPE, res_DESC, res_diff
}

// Ultimate get diff of two dumbers
// Pass 1st num, sec num... and if desired, int for decimal precision of output
func GET_DIFF(ALL_PARAMS ...interface{}) float64 {
	var small float64
	var large float64

	var precision = 1

	for n, param := range ALL_PARAMS {
		int_val, IS_INT := param.(int)
		float_val, IS_FLOAT := param.(float64)
		string_val, IS_STRING := param.(string)

		// Must be first anad second num  (for n == 0 and n == 1)
		if n == 0 {
			if IS_INT {
				small = float64(int_val)
			} else if IS_FLOAT {
				small = float_val
			}
			continue
		}

		if n == 1 {
			if IS_INT {
				large = float64(int_val)
			} else if IS_FLOAT {
				large = float_val
			}
			continue
		}

		// Must be an int.. this is for preceission
		if n == 2 {

			precision = int_val
		}

		// PLACEHOLDER .. if a string is passed.. do something with it
		if IS_STRING && string_val != "" {

		}

	} //end of params

	// Error handling
	if small == large {
		return 0.0
	}

	diff := float64(large) - float64(small)
	if diff < 0.0 {
		diff = small - large
	}
	final_diff := FIX_FLOAT_PRECISION(diff, precision)

	return final_diff

}

// Ultimate percentage of two numbers
// pass small num first... THEN large num
func PERCENT_OF(ALL_PARAMS ...interface{}) float64 {
	var small float64
	var large float64

	var precision = 2
	for n, param := range ALL_PARAMS {
		int_val, IS_INT := param.(int)
		float_val, IS_FLOAT := param.(float64)
		string_val, IS_STRING := param.(string)

		// Must be first anad second num  (for n == 0 and n == 1)
		if n == 0 {
			if IS_INT {
				small = float64(int_val)
			} else if IS_FLOAT {
				small = float_val
			}
			continue
		}

		if n == 1 {
			if IS_INT {
				large = float64(int_val)
			} else if IS_FLOAT {
				large = float_val
			}
			continue
		}

		// Must be an int.. this is for preceission
		if n == 2 {

			precision = int_val
		}

		// PLACEHOLDER .. if a string is passed.. do something with it
		if IS_STRING && string_val != "" {

		}

	} //end of params

	// Error handling
	if small == 0.0 && large == 0.0 {
		return 0.0
	}
	if small == large {
		return 99.9
	}
	if small == 0.0 {
		return 0.0
	}

	// DONT FLIP THE NUMBERS automatically.. this returns an INACCURATE PERCENTAGE

	divis := small / large
	perc := divis * 100
	perc = FIX_FLOAT_PRECISION(perc, precision)

	return perc

}

// alias
func GET_PERCENT_OF(ALL_PARAMS ...interface{}) float64 {
	return PERCENT_OF(ALL_PARAMS...)
}

// Ultimate get AVG function.. Pass either an ARRAY OF FLOAT or ARRAY of INT (and a precision int if you need to for output)
func GET_AVG(ALL_PARAMS ...interface{}) float64 {

	var precision = 2
	var float_arr []float64
	var int_arr []int

	var use_float = false
	for n, param := range ALL_PARAMS {
		tmp_int, IS_INT := param.([]int)
		tmp_float, IS_FLOAT := param.([]float64)

		prec_int, IS_PREC := param.(int)

		if n == 0 {
			// If they passed an array of int
			if IS_INT {
				int_arr = tmp_int
			} else if IS_FLOAT {
				float_arr = tmp_float
				use_float = true
			}

			continue
		}

		//if they also passed an int for precisions
		if n == 1 && IS_PREC {
			precision = prec_int
		}
	}

	//2. Now get the sum
	var sum = 0.0
	var total_items = 0
	if use_float {
		for _, x := range float_arr {
			sum += x
		}
		total_items = len(float_arr)
	} else {
		for _, x := range int_arr {
			sum += float64(x)
		}
		total_items = len(int_arr)
	}

	//3. Now get average

	avg := sum / float64(total_items)

	final_avg := FIX_FLOAT_PRECISION(avg, precision)

	return final_avg

}
