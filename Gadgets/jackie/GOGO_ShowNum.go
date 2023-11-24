package CUSTOM_GO_MODULE

import (
	//"crypto/md5"
	//"encoding/hex"		

	
	"github.com/dustin/go-humanize"
)




// This utilizes the HUMANIZE library and shows a HUMAN readable number of the passed variable
func ShowNum(innum int) string {

	result := humanize.Comma(int64(innum))

	// result = strconv.Itoa(result)
	return result
}


// Shows a pretty Number based on passed FLOAT
func ShowNum_FLOAT(innum float64) string {

        result := humanize.Comma(int64(innum))

        // result = strconv.Itoa(result)
        return result
}

// 64bit version of this.. not sure why im using this yet
func ShowNum64(innum int64) string {

	result := humanize.Comma(innum)

	// result = strconv.Itoa(result)
	return result
}


/*
	= = = = = = 
	= = = = = =  Better versions of the IS_xxx functions
	= = = = = = 
*/
func IS_INT(param interface{}) bool {
	_, found_int := param.(int) 
	_, found_int32 := param.(int32) 
	_, found_int64 := param.(int64)
	
	if found_int || found_int32 || found_int64{
		return true
	}

	return false
}
func IS_STRING(param interface{}) bool {

	_, found_string := param.(string) 
	return found_string
}
func IS_FLOAT(param interface{}) bool {
	_, found_32 := param.(float32) 
	_, found_64 := param.(float64) 
	
	if found_32 || found_64 {
		return true
	}

	return false
}
func IS_BOOL(param interface{}) bool {

	_, found_bool := param.(bool) 
	return found_bool
}

