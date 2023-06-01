package GOGO_MDC

import (
		//"strings"
		"math"
		//"strconv"

		. "github.com/ace2z/GOGO/Gadgets"
)




// Returns Percentages INCREASE DECREASE for stocks etc... Takes in floats or INTs
/*
	As of MAY 2023.. This is the ULTIMATE GET PERCENTAGE function (replaces the previous GET_PRECENT)
	Also returns the DIFF between two numbers and if it was an INCREASE or DECERASE based on which number was passed FIRST

*/
func GET_INCDEC_PERCENT(ALL_PARAMS ...interface{}) (string, float64, float64) {
	
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

				       if n == 0 { firstNUM = float64(intval) 
                } else if n == 1 { secNUM   = float64(intval) }

			} else if IS_FLOAT {
				       if n == 0 { firstNUM = floatval
				} else if n == 1 { secNUM = floatval }

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

		return res_DESC, res_PERC, res_diff
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
		SHOW_BOX("GET_PERCENTAGE", "|cyan|" + firstSTRING + " --> " + secSTRING, "|yellow|" + mode + " by", "|green|" + percSTRING, "DIFF: ", "|yellow|" + diffstring)
	}

	// Finally return everything
	return res_DESC, res_PERC, res_diff
}


//GET_INC_DEC_PERCENT(ALL_PARAMS ...interface{})
// Alias to GET_INCDEC_PERCENT
func GET_PERCENT(ALL_PARAMS ...interface{}) (string, float64, float64) {
	return GET_INCDEC_PERCENT(ALL_PARAMS...)
}

// Alias Hypbrid... JUST returns JUST the DIFF of the two numbers
func GET_DIFF(ALL_PARAMS ...interface{}) float64 {
	_, _, diff := GET_INCDEC_PERCENT(ALL_PARAMS...)
	return diff
}

// alias Hybrid.. returns JUST the perc of two numbers
func GET_PERC(ALL_PARAMS ...interface{}) float64 {
	_, perc, _ := GET_INCDEC_PERCENT(ALL_PARAMS...)
	return perc
}



