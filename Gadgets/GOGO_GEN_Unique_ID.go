/*
 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	"regexp"
	"strconv"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

/*
Generates a UNIQUE ID based on passed critieria. Mostly for Mongo ID vals
Also accepts:
-prefix - for an value to go before the id
-raw    - if you do DO NOT want to get an md5 generated value
-delim  - if you want to use somethiing OTHER than | for the delimiter
*/
func GEN_UNIQUE_ID(GENFROM ...interface{}) string {
	var result = ""
	var use_raw_mode = false

	var DEFAULT_delim = "|"

	var USE_PREFIX = ""

	for n, field := range GENFROM {
		val_int, IS_INT := field.(int)
		val_float, IS_FLOAT := field.(float64)
		val_string, IS_STRING := field.(string)
		val_bool, IS_BOOL := field.(bool)

		if IS_INT {
			result = result + INT_to_STRING(val_int) + DEFAULT_delim
			continue
		}

		if IS_FLOAT {
			result = result + FLOAT_to_STRING(val_float) + DEFAULT_delim
			continue
		}

		// Any values passed goes towards making up the unique id. We always remove spaces
		// Unless its a parameter like -prefix
		if IS_STRING {

			if val_string == "-prefix" {
				o := n + 1
				if o < len(GENFROM) {

					tval_string, tfound := GENFROM[o].(string)
					if tfound {
						USE_PREFIX = tval_string
					}
				}
				continue
			}

			// RAW mean we use the actual values that are passed for the id.. instead of the generated MD5
			// Good for some cases like troubleshooting
			if val_string == "-raw" {
				use_raw_mode = true
				continue
			}

			val_string = clearString(val_string)
			val_string = strings.Replace(val_string, " ", "", -1)
			result = result + val_string + DEFAULT_delim
			continue
		}

		if IS_BOOL {
			result = result + strconv.FormatBool(val_bool) + DEFAULT_delim
			continue
		}
	}

	//2. Our result
	result = strings.TrimSuffix(result, DEFAULT_delim)
	result = strings.ToLower(result)

	//3. Finally generate the md5 of this result
	var final_result_id = GET_MD5(result)

	//4. However is "-raw" was specified.. we just return the result, not the md5
	if use_raw_mode {
		final_result_id = result
	}

	//5. And if a prefix was specified
	if USE_PREFIX != "" {
		final_result_id = USE_PREFIX + final_result_id
	}

	return final_result_id
}
