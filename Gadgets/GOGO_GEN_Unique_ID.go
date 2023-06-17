/*   
	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	"regexp"
	"strings"
	"strconv"
)


var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)
func clearString(str string) string {
    return nonAlphanumericRegex.ReplaceAllString(str, "")
}
// gens a 'unique' id from the values provided. Use this to insert into the ID field for mongo
func GEN_UNIQUE_ID(GENFROM ...interface{} ) (string, string) {
	var result = ""

	var USE_PREFIX = ""
	var USE_PIPE = false

	for n, field := range GENFROM {
		val_int, IS_INT := field.(int)
		val_float, IS_FLOAT := field.(float64) 
		val_string, IS_STRING := field.(string) 
		val_bool, IS_BOOL := field.(bool) 		

		if IS_INT {
			result = result + INT_to_STRING(val_int) + "_"
			continue
		}

		if IS_FLOAT {
			result = result + FLOAT_to_STRING(val_float) + "_"
			continue
		}

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
			} else if val_string == "-pipe" {
				USE_PIPE = true
				continue
			}

			val_string = clearString(val_string)
			val_string = strings.Replace(val_string, " ", "", -1)
			result = result + val_string + "_"
			continue
		}

		if IS_BOOL {
			result = result + strconv.FormatBool(val_bool) + "_"
			continue
		}
	}

	//2. Our result
	result = strings.TrimSuffix(result, "_")
	result = strings.ToLower(result)

	//3. Also..generate the MD5 of this 'unique' id.. 
	md5string := GET_MD5(result)	

	//4. If we default to _ .. if we want we can use | as the seperator for the uniq id's
	delim := "_"
	if USE_PIPE {
		delim = "|"
	}
	if delim != "" {
		for i := 5; i < len(md5string); i += 6 {
			md5string = md5string[:i] + delim + md5string[i:]
		}		
	}

	//5. And if a prefix was specified
	if USE_PREFIX != "" {
		result = USE_PREFIX + result
		md5string = USE_PREFIX + md5string
	}
	

	return result, md5string
}
