package CUSTOM_GO_MODULE

import (
	"bufio"
	"encoding/json"
	"math/rand"
	"os"
	"strings"

	"github.com/atotto/clipboard"
)

// When called, copies a specified string to the users CLIPBOARD
func CLIPBOARD_COPY(instring string) {
	clipboard.WriteAll(instring)
}


// This takes IN a string and returns a shuffle of the characters contained in it
func SHUFFLE_STRING(input_STRING string) string {

	//1. Get the length of the string
	slen := len(input_STRING)

	stringRUNE := []rune(input_STRING)

	shuffledString_RESULT := make([]rune, slen)

	for i := range shuffledString_RESULT {
		shuffledString_RESULT[i] = stringRUNE[rand.Intn(slen)]
	}
	return string(shuffledString_RESULT)
} // end of genSESSION


func VERIFICATION_PROMPT(warning_TEXT string, required_input string, ALL_PARAMS ...interface{}) bool {
	var MAX_ATTEMPTS = 3
	var exit_on_fail = false

	for _, param := range ALL_PARAMS {
		string_val, is_string := param.(string)
		int_val, is_int := param.(int)

		// If parma is an int, means they are passing MAX_ATTEMPTS
		if is_int {
			MAX_ATTEMPTS = int_val
			continue
		}
		if is_string {
			if string_val == "-exit_on_fail" {
				exit_on_fail = true
				continue	
			}
 		}
	} //end

	M.Println("\n      - - - - - - - - WARNING - - - - - - - - - - - - - -")
	
	for x := 0; x < MAX_ATTEMPTS; x++ {
		C.Println("")
		C.Println("      ", warning_TEXT)
		C.Println("")
		C.Print("       Type: ")
		G.Print(required_input)
		C.Println(" To Continue")
		Y.Print("       RESPONSE: ")
		userResponse := GET_USER_INPUT()

		if userResponse == required_input {
			return true
		} else {
			R.Println("\n ! ! ! ! ! ! INVALID RESPONSE  ! ! ! ! ! !")
		    M.Println("\n     - - - - - - - - - - - - - - - - - - - - - - - - -")			
		}
	} //end of for
	

	//2. If we get this far without a valid response, we will exit the program without proceeding
	if exit_on_fail {
		os.Exit(-9)
	}

	return false

} //end of prompt



func GET_USER_INPUT(ALL_PARAMS ...interface{}) string {

	var showtyped = false

	for _, param := range ALL_PARAMS {
		string_val, is_string := param.(string)
		//int_val, is_int := param.(int)

		if is_string {
			if string_val == "-showtyped" {
				showtyped = true
				continue	
			}
 		}
	} //end

	reader := bufio.NewReader(os.Stdin)
	userTEMP, _ := reader.ReadString('\n')
	userTEMP = strings.TrimSuffix(userTEMP, "\n")

	if showtyped {
		Y.Print("\n     You Typed: ")
		W.Print(userTEMP)
		Y.Println("**")
	}

	return userTEMP

} //end of

func GET_INPUT() string {
	return GET_USER_INPUT()
}



// Simple PressAny Key function
func PressAny(ALL_PARAMS ...interface{}) {
	var be_quiet = false 


	for _, param := range ALL_PARAMS {
		//string_val, is_string := param.(string)

		bool_val, is_bool := param.(bool)

		if is_bool {
			be_quiet = bool_val
		}
	}	

	if be_quiet == false {
		W.Println("")
		W.Println("         ...Press Enter to Continue...")
		W.Println("")
	}

	//1. New way of doing PAK
	b := make([]byte, 10)
	if _, err := os.Stdin.Read(b); err != nil {
		R.Println("Fatal error in PressAny Key: ", err)
	}

} // end of func

func PROMPT(warning_TEXT string, required_input string) {
	VERIFICATION_PROMPT(warning_TEXT, required_input)
}

var PAGE_COUNT = 0
var PAGE_MAX = 5
// This is a basic Paging routine that prompts you to PressAny key
// after x number of items have been shown
func Pager(tmax int) {
	PAGE_MAX = tmax
	PAGE_COUNT++

	if PAGE_COUNT == PAGE_MAX {
		C.Print("   - - PAGER - -")
		PressAny()
		PAGE_COUNT = 0
	}

} //end of Pager


// This generates a serial.. usually used discern between multiple execution runs like in jenkins
func GenSerial(serial_length int) string {

	result := SHUFFLE_STRING("grzbjhuflcekivxmntqpsoadwy527183469")

	part_ONE := result[0:4]
	part_TWO := result[3:serial_length]

	final_res := part_ONE + "-" + part_TWO

	return final_res

} // end of GenSerial





func PRETTY_STRUCT_json(input interface{}) string {
	byte_json, _ := json.MarshalIndent(input, "", "\t")  // Marshall takes a struct and makes it into JSON

	result := string(byte_json)

	return result
}
