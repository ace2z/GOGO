package CUSTOM_GOMOD

import (
	//. "github.com/ace2z/GOGO/Gadgets"
)


type URL_PARAMS struct {
	KEY		string
	Value	string
}

// Gets a value from URL params. Takes in a list of URL Params
func GET_VALUE(KEY string, inputVARS []URL_PARAMS) string {
	//1. FIrst we iterate through the Url Params looking for the one that matches they KEY speicfied
	for _, x := range inputVARS {
		if x.KEY == KEY {
			return x.Value
		}
	} //end of for


	//2. otherwise if we find nothing, return nothing
	return ""

} // end of func

// Alias to GET_VALUE
func FIND_VALUE(KEY string, inputVARS []URL_PARAMS) string {
	return GET_VALUE(KEY, inputVARS)
}
func GET_KEY(KEY string, inputVARS []URL_PARAMS) string {
	return GET_VALUE(KEY, inputVARS)
}

// This is mostly for debug, just shows all values in an URL_PARAMS list
// returns  a json formatted string
func SHOW_ALL_PARAMS(inputVARS []URL_PARAMS) string {

	var JSON_OUTPUT = ``
	for _, x := range inputVARS {
		JSON_OUTPUT += "     " + x.KEY + `:` + x.Value + `,
`		
	} //end of 

	return JSON_OUTPUT	
} //edn of SHOW ALL

