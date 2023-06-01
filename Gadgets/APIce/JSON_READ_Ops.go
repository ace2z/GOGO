package CUSTOM_GOMOD

import (

//	"bufio"
//	"io"
	"strings"
	"io/ioutil"
	"net/http"
//	"os"
	"context"
	"time"
//	"encoding/json"

	. "github.com/ace2z/GOGO/Gadgets"

	//"github.com/buger/jsonparser"

)


var MAX_JSON_RETRY_ATTEMPTS = 5
var MAX_JSON_SLEEP_VAL = 10





/*
	Reveised 2023 JSON DOWNLOAD
	- You can send it POST statements  (aafter the url... add "DO_POST" as a parameter)
	- If you want to changeg the HTTP timeout.. specifiy an integer (defaults to 30)
	- ...and for --data payloads for POST  .. specifiy with a New Reader as follows: 

	{
		"symbols": ["AAPL", "MSFT", "GOOG"],
		"intervals": ["1h", "1day"],
		"outputsize": 3,
		"methods": [
			"time_series",
			{
				"name": "price"
			}  
		]
	}
	- HEADERS! better header support... specify Headers as a parameter like this (with | pipe):
		"Content-Type|application/json"
		valid, JSON_BYTE, jsonTEXT, _ := JSON_DOWNLOAD(URL, "DO_POST", payload, "Content-Type|application/json")

*/
func JSON_DOWNLOAD(API_URL string, PARAMS ...interface{}) (bool, []byte, string, error) {

	C.Println(" = = JSON Download = =")
	var WAS_SUCCESS = false
	var final_ERROR error
	var TIMEOUT_SECS = 30

	var response_TEXT = ""
	var USE_DATA = false
	var DATA_PAYLOAD *strings.Reader
	var HEADERS []string

	for _, VAL := range PARAMS {
		INT_val, is_INT := VAL.(int)
		STR_val, is_STRING := VAL.(string)

		// If an integrag val was specified, assume is is the timeout in seconds
		if is_INT {
			TIMEOUT_SECS = INT_val
			continue
		}
		
		if is_STRING {
			// If its a string.. check if its DO_POST
			if STR_val == "DO_POST" {
			}

			// Otherwise... assume its a HEADER
			if strings.Contains(STR_val, "|") {
				HEADERS = append(HEADERS, STR_val)
			}			
		}

		// if this is neither INT.. or STRING... its probably a reader (for --data payloads for POSTS) .. assume it is
		if is_INT == false && is_STRING == false {	
			DATA_PAYLOAD = VAL.(*strings.Reader)
			USE_DATA = true
		}
	} // end of PARAMS for

	//1. Setup our REQUEST
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(TIMEOUT_SECS) * time.Second)
	defer cancel()

	var req *http.Request
	var err error

	// If we are using POST mode
	if USE_DATA {

		req, err = http.NewRequestWithContext(ctx, http.MethodPost, API_URL, DATA_PAYLOAD)

	// else default to get
	} else {
		req, err = http.NewRequestWithContext(ctx, http.MethodGet, API_URL, nil)
	}
	
	if err != nil {
		M.Println("JSON_DOWNLOAD NewRequest ERROR: ")
		Y.Println(err)
		final_ERROR = err
	}		

	//2b. Now If headers were specified, lets iterate and add them
	// Since we are expecting JSON back.. we always add the following:
	req.Header.Add("Content-Type", "application/json")

	if len(HEADERS) > 0 {
		for _, x := range HEADERS {
			if strings.Contains(x, "|") {
				msplit := strings.Split(x, "|")
				if len(msplit) > 1 {
					header_KEY := msplit[0]
					header_VAL := msplit[1]
					req.Header.Add(header_KEY, header_VAL)
				}
			}
		} //end of for
	}
/*

		res, err2:= client.Do(req)
		if err2 != nil {
			M.Println(" Client.Do ERROR: ")
			Y.Println(err2)
			return
		}
		defer res.Body.Close()
	  
		body_BYTES, err3 := ioutil.ReadAll(res.Body)
*/
	//4. Next get the response
	client := &http.Client{}
	//resp, err2 := http.DefaultClient.Do(req)
	resp, err2 := client.Do(req)

	if err2 != nil {
		Y.Println(" JSON Read error in RESP: ", err2)
		return false, []byte{}, response_TEXT, final_ERROR
	}

	defer resp.Body.Close()

	if err2 != nil {
		M.Println("JSON_DOWNLOAD Client RESP ERROR: ")
		Y.Println(err2)
		final_ERROR = err2	
	}

	// Now lets read the response we got into parsable json
	JSON_BYTE_OBJ, err3 := ioutil.ReadAll(resp.Body)

	//3. BE SURE TO CLOSE
	resp.Body.Close()

	if err3 != nil {
		// handle error
		M.Println("JSON_DOWNLOAD ReadAll ERROR: ")
		Y.Println(err3)
		final_ERROR = err3

		return false, []byte{}, response_TEXT, final_ERROR
	}
	
	WAS_SUCCESS = true

	//5. SUCCESSS! Convert our response to TEXT
	response_TEXT = string(JSON_BYTE_OBJ)
	

	return WAS_SUCCESS, JSON_BYTE_OBJ, response_TEXT, final_ERROR

} //end of JSON_DOWNLOAD



