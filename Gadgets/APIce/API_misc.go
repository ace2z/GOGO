package CUSTOM_GOMOD

import (

	"net/url"
	"strings"
	"unicode"
//	"context"
//	"time"
	"encoding/json"

	. "github.com/ace2z/GOGO/Gadgets"
)

// This makes "API" json you can retrieve from jquery or VUEjs
type API_JSON_OBJ struct {
	DATA 		[]interface{}		`json:"data"`
}
func MAKE_API_JSON(tmpOBJ interface{}) string {

	var a API_JSON_OBJ
	a.DATA = append(a.DATA, tmpOBJ)

	JSON_RESULT, err := json.MarshalIndent(a, "", "\t")  // Marshall takes a struct and makes it into JSON
	
	if err != nil {
		R.Println(" error in the MAKE_API_JSON ")
		W.Println(err)
		return ""
	}	

	return string(JSON_RESULT)
}

func MinifyJSON(str string) string {
	return strings.Map(func(r rune) rune {
		 if unicode.IsSpace(r) {
			  return -1
		 }
		 
		 return r
	}, str)
} //end of 




/* This is meant to be passed a keymap of url.Values
Returns true and expects a POINTER to hold the value that it finds
If keyname is found in the map, its VALUE is returned
NOTE: this lets you send URL vars of ANY case

FIND_URL_key_using_POINTER("data", URL_MAP, &result):
*/
func FIND_URL_key(keyname string, UMAP url.Values, myresult *string) bool {

	//1. this allows you to specify case INSENSITIVE keyNames
	kUpper := strings.ToUpper(keyname)
	klow := strings.ToLower(keyname)
	firstLetter := UpperFirst(keyname)

	if keyValue_ARRAY, ok := UMAP[kUpper]; ok {

		*myresult = keyValue_ARRAY[0]

		return true
	}

	if keyValue_ARRAY, ok := UMAP[klow]; ok {

		*myresult = keyValue_ARRAY[0]

		return true
	}

	if keyValue_ARRAY, ok := UMAP[firstLetter]; ok {

		*myresult = keyValue_ARRAY[0]

		return true
	}

	return false
}

/*

	FIND_URL_key
	Takes in a key to search for
	this is similar to FIND_URL_key but doesnt use a pointer
*/
func Simple_FIND_URL_key(keyname string, UMAP url.Values) (bool, string) {

	//1. this allows you to specify case INSENSITIVE keyNames
	kUpper := strings.ToUpper(keyname)
	klow := strings.ToLower(keyname)
	firstLetter := UpperFirst(keyname)

	if keyValue_ARRAY, ok := UMAP[kUpper]; ok {

		return true, keyValue_ARRAY[0]
	}

	if keyValue_ARRAY, ok := UMAP[klow]; ok {

		return true, keyValue_ARRAY[0]

	}

	if keyValue_ARRAY, ok := UMAP[firstLetter]; ok {

		return true, keyValue_ARRAY[0]

	}

	return false, ""
} //end of func


