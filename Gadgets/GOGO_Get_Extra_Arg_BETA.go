/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	//"crypto/md5"
	//"encoding/hex"		

	
	//"github.com/dustin/go-humanize"
)




/*
	FOr Dynamic Parameter support... DONT FORGET THE ... thre dots!!!
	Examples:
		1. for function call:
			WRITE_EMBED("ca.pem.gz", "verbose", "quiet", 55, "www.podshop.com")
		2. Then write the function like this:
			func WRITE_EMBED(name string, EXTRA_ARGS ...interface{}) string {
				var verbose = GET_EXTRA_ARG("verbose",  EXTRA_ARGS...).(bool)		// specify verbose will set to TRUE... case verbase was passed as a parameter
				var bequiet = GET_EXTRA_ARG(1, EXTRA_ARGS...).(string)				// parameter at index 1 is quiet.. So bequiet is set to string "quiet"
				
				var antcount = GET_EXTRA_ARG(2, EXTRA_ARGS...).(int)				// parameter at index 2 is the number 55... so antcount is set to  specify 2 to get the INT value 55 that was specified
				var mydomain = GET_EXTRA_ARG(3, EXTRA_ARGS...).(string)			  // and.. specify 3 to get the param at index 3 (which is www.podshop.com)
				C.Println(" RESULT Verbose is: ", verbose)
				C.Println(" antcount is: ", antcount)
				C.Println(" mydomain is: ", mydomain)
			}
*/
func GET_EXTRA_ARG(key interface{}, EXTRA_ARGS ...interface{}) interface{} {

	var find_by_INDEX = false
	var find_by_STRING = false


	var IND_2_find = -9
	var KEY_VAL_2_find = ""
	

	var EMPTY_RES interface{}

	/*
	
		We'll ether have an INT (implying we GET the value of the parameter passed..by INDEX
		Or we will have an explicit string ("meaning if this value is found, return true or false")
		
		if by index, doesnt matter if it is bool, float, string or whatever.. its returned explicitly
	*/
	
	if IS_INT(key) {
		find_by_INDEX = true
		IND_2_find = key.(int)

	} else if IS_STRING(key) {
		find_by_STRING = true
		KEY_VAL_2_find = key.(string)
		
	} else {
		R.Println(" Invalid Key type! Must be either INT (for index) or string (for key val)")

		return EMPTY_RES
	}

	// Now iterate through the args!
	for ind, arg := range EXTRA_ARGS {

		// Now if an index was passed.. and we match it.. we return the value at that index verbatim
		if find_by_INDEX {
			if ind == IND_2_find {
				return arg
			}
		} 		

		// If a KEy was passed (a  string) we just return true or false
		// example.. if someone passes verbose, we return true... which will turn on verbose mode
		// even tho the arg value is returned, we dont use it cause its identical to KEY_VAL_2_find
		if find_by_STRING {
			PARAM_string, FOUND_string := arg.(string)
			if FOUND_string {
				if PARAM_string == KEY_VAL_2_find {
					return arg
				}
			}
		}
	}


	// default return value if we ever get this far
	return EMPTY_RES

} //end of