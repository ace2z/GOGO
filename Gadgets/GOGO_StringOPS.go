/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	"regexp"
	"strings"
	"unicode"
	"crypto/md5"
	"encoding/hex"		
	"strconv"
)



// Returns true if the string contains ONLY numbers
func HasOnlyNumbers(s string) bool {
    for _, r := range s {
        if (r < '0' || r > '9') {
            return false
        }
    }
    return true
} //end of func

// Make sthe first character of a string UPPER CASE
func UpperFirst(inString string) string {

	a := []rune(inString)
	a[0] = unicode.ToUpper(a[0])

	return string(a)
}


// This removes all spaces from a string via unicode
func UNICODE_REMOVE_ALL_SPACES(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}

		return r
	}, str)
}

//  This removes all extra spaces in a string 
func NO_EXTRA_Spaces(input string) string {

	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	final := re_leadclose_whtsp.ReplaceAllString(input, "")
	final = re_inside_whtsp.ReplaceAllString(final, " ")

	return final
}

// This takes in a string and removes all non alphanumeric chars from it.. and extra spaces
func JUST_ALPHA_STRING(input string) string {
	C.Println("\nCleansing string of no-alpha chars")
	justAlpha, _ := regexp.Compile("[^a-zA-Z0-9_ ]")
	killExtraSpace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	PASS_1 := justAlpha.ReplaceAllString(input, "")
	FINAL_PASS := killExtraSpace.ReplaceAllString(PASS_1, " ")

	return FINAL_PASS
} //end of func




/*
	This takes in a number (float) and shows it as pretty PERCENT string
	// Specify another INT to define PRECISION
*/
func SHOW_PERCENT(ALL_PARAMS ...interface{}) string {

	var inputNUM = 0.0
	var PRECISION = 1

	for n, param := range ALL_PARAMS {
		intval, IS_INT := param.(int)
		floatval, IS_FLOAT := param.(float64)
		// boolval, IS_BOOL := param.(bool)

		// First parmater is always the input
		if n == 0 {

			if IS_INT { 
				inputNUM = float64(intval)

			} else if IS_FLOAT {
				inputNUM = floatval
			}
			continue
		}

		if IS_INT { 
			PRECISION = param.(int)
		}	
	} //end of for	

	percSTRING := strconv.FormatFloat(inputNUM, 'f', PRECISION, 64)		// Make a string of num with specific # of PRECISION dec points
	percSTRING = percSTRING + "%"

	return percSTRING
}



// GEts the MD5 of a string
func GET_MD5(input string) string {

    hasher := md5.New()
    hasher.Write([]byte(input))

	//Get the 16 bytes hash
	hashInBytes := hasher.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String := hex.EncodeToString(hashInBytes)
	
	return returnMD5String
}


func STRING_to_INT(input string) int {
	result, _ := strconv.Atoi(input)

	return result
}
func STRING_to_INT64(input string) int64 {
//	result, _ := strconv.Atoi(input)
	result, _ := strconv.ParseInt(input, 10, 64)

	return result
}


func STRING_to_FLOAT32(input string) float32 {
	result, _ :=  strconv.ParseFloat(input, 32)
	return float32(result)
}

func STRING_to_FLOAT(input string) float64 {
	result, _ :=  strconv.ParseFloat(input, 64)
	return result
}

// Converts a float to a string.. If another number is specified, it is interpreted as decimal precision
func FLOAT_to_STRING(input float64,  ALL_PARAMS ...int) string {

	var prec = -1
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			prec = VAL
		}	
	}

	result := strconv.FormatFloat(input, 'f', prec, 64)
	
	return result
}
func FLOAT_to_STRING32(input float32,  ALL_PARAMS ...int) string {

	var prec = -1
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			prec = VAL
		}	
	}

	result := strconv.FormatFloat(float64(input), 'f', prec, 32)
	
	return result
}
func INT_to_STRING(input int) string {
	return strconv.Itoa(input)
}

func INT64_to_STRING(input int64) string {
	return strconv.Itoa(int(input))
}
