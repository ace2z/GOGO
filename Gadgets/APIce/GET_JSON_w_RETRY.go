package CUSTOM_GOMOD

import (

	//	"bufio"
	//	"io"
	"strings"
	//	"os"

	//	"encoding/json"

	. "github.com/ace2z/GOGO/Gadgets"
	//"github.com/buger/jsonparser"
)

var RETRY_MAX = 500

type TMP_ERR_OBJ struct {
	MSG       string
	HANDLE_BY string
}

var LIMIT_ERRORS = []string{
	"You have run out of API credits",
}

var ABORT_ERRORS = []string{
	"deadline exceeded",
	"**symbol** not found",
	"Bad link",
	"deadline exceeded",
}

// = = = =Mostly API secific checks
func OVER_THE_LIMIT(jstext string, GLOBAL_PREFIX string, API_SLEEP int) bool {

	for _, msg := range LIMIT_ERRORS {
		if strings.Contains(jstext, msg) {
			Y.Println(GLOBAL_PREFIX, " WARNING: Limit Error")
			W.Println(jstext)
			Sleep(API_SLEEP, true)
			return true
		}
	}

	return false
}

func HAVE_ABORT_ERROR(jstext string, GLOBAL_PREFIX string) bool {

	for _, msg := range ABORT_ERRORS {
		if strings.Contains(jstext, msg) {
			Y.Println(GLOBAL_PREFIX, " CRITICAL ABORT Error:")
			W.Println(jstext)
			return true
		}
	}
	return false
}

/*
Retrieves JSON from API's using a retry method
Dont forget to append strings to LIMIT_ERRORS or ABORT_ERRORS arrays if needed
Returns BOOL, Byte and JSON_TEXT
*/
func GET_JSON_w_RETRY(URL string, PARAMS ...interface{}) (bool, []byte, string) {
	var JSON_BYTE_OBJ []byte
	var byte_VALID = false
	var result_JSON_TEXT = ""
	var POST_PAYLOAD *strings.Reader

	var retry_sleep = 15

	var use_POST = false

	// Get params if any

	for n, field := range PARAMS {
		//val_int, IS_INT := field.(int)
		//val_float, IS_FLOAT := field.(float64)
		val_string, IS_STRING := field.(string)
		//val_bool, IS_BOOL := field.(bool)
		val_reader, IS_POST_READER := field.(*strings.Reader)

		if IS_POST_READER {
			POST_PAYLOAD = val_reader
			use_POST = true
			continue
		}
		if IS_STRING {
			not_enough_specified := false
			o := n + 1
			if o >= len(PARAMS) {
				not_enough_specified = true
			}

			if val_string == "-sleep" && not_enough_specified == false {
				extra_val, _ := PARAMS[o].(int)
				retry_sleep = extra_val
			}

			if not_enough_specified {
				M.Print("Error! invalid num params sent to GET_JSON_w_RETRY for ")
				W.Println(val_string)
				M.Println("Need an additional paremter value as INT")
				DO_EXIT()
			}

			continue

		}

		// If this is an

	} // end of params for

	//1. We run this in a RETRY loop.. so we are able to work through the Limits imposed by 12Data
	for r := 1; r < RETRY_MAX; r++ {
		if r > RETRY_MAX {
			M.Println(" WARNING: Reached RETRY Max! this normally should never happen!")
			DO_EXIT()
		}
		if r > 1 {
			Y.Print(GLOBAL_PREFIX, " RETRY Attempt: ")
			M.Println("(", r, " of ", RETRY_MAX, ")")
		}
		var valid = false
		var JSON_TEXT = ""
		var err error

		// Default is to use a GET request. If however we want to use POST.. check for that here
		if use_POST == false {
			valid, JSON_BYTE_OBJ, JSON_TEXT, err = JSON_DOWNLOAD(URL)
			// If using POST
		} else {
			valid, JSON_BYTE_OBJ, JSON_TEXT, err = JSON_DOWNLOAD(URL, POST_PAYLOAD)
		}

		if valid == false {
			continue
		}
		if err != nil {
			M.Println(" JSON_DOWNLOAD Error: ")
			Y.Println(err)
			continue
		}

		// Retry/ SLEEP conditions
		if OVER_THE_LIMIT(JSON_TEXT, GLOBAL_PREFIX, retry_sleep) {
			continue
		}

		// Abort Conditions
		if HAVE_ABORT_ERROR(JSON_TEXT, GLOBAL_PREFIX) {
			break
		}

		result_JSON_TEXT = JSON_TEXT
		byte_VALID = true
		break
	} //end of RERY

	return byte_VALID, JSON_BYTE_OBJ, result_JSON_TEXT
}
