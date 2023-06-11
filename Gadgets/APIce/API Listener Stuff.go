package CUSTOM_GOMOD

import (

	"flag"
//	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"unicode"
//	"context"
//	"time"
	"encoding/json"

	. "github.com/ace2z/GOGO/Gadgets"
)

/*
	###
	###		Globals!!
	###
*/

var DEFAULT_LISTEN_PORT = "8888" // This is the default "listen" port for all webservices.. if using ports below 1024 you must run as root
var USE_API_SETUP_DEFAULT_PARAMS = false
var JSON_CLEAN_FLAG	   = false	// When set to TRUE, we return "MINIFIED" json  (it omits whatever spaces and indents you used get stripped out )

var SSL_ENABLE_FLAG     = false	// If set to true.. we listen in SSL mode.. meaning we need a PEM file and a KEY specified below
var SSL_CERT_PEM_FILE   = ""		// Full path to wherever this CERT/ pem file is			(this is just example: /opt/SSL_CERTS/biolab_COMBINED.pem  )
var SSL_KEY_FILE        = ""		// Full path to wherever the KEY file of this cert is (this is just example: /opt/SSL_CERTS/biolab_ALPHA-SSL.key )

type HEADER_OBJ struct {
	NAME    string
	VALUE   string
}
type API_JSON_OBJ struct {
	DATA 		[]interface{}		`json:"data"`
}

// This makes "API" json you can retrieve from jquery or VUEjs
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



/*

 Sends a response to HTTP requests. Defaults to a generic JSON response

*/
func SEND_Response_2_Client(JSON_Response string, ClientResponseWriter *http.ResponseWriter) {

	//M.Print("\n\n   *** Sending Response Back to HTTP Client ...\n")

	//1. Set the outgoing headers for this message type:
	(*ClientResponseWriter).Header().Set("Content-Type", "application/json") 	// ---- Set THIS header for sending back an actual json object
	(*ClientResponseWriter).Header().Set("Access-Control-Allow-Origin", "*") 	//  --- always need this regardless

	//2. This is a default message.. i keep this for a way to debug (meaning the REST stuff works but the JSON is just fucked up)
	defaultMessage := `
    [
		{
			"ERROR (default Resp #1)" : "No parameters specified",
		  	"ERROR (default Resp #2)" : "or.. Invalid API endpoint requested"
		}
    ]
`
	//2b. If no json provided, we give a default response
	if JSON_Response == "" {

		JSON_Response = defaultMessage

	}

	//3. Now lets check the JSON_CLEAN_FLAG .. if it is true, we strup out all the spaces and tabs
	if JSON_CLEAN_FLAG {
		JSON_Response = MinifyJSON(JSON_Response)
	}

	/*
			5. OLD  Method..uses NewEncoder
			err := json.NewEncoder(*ClientResponseWriter).Encode(JSON_Response)

			5c. This method just pushes the already formatted outgoing JSON directly on the ClientResponseWriter

			 rawIn := json.RawMessage(JSON_Response)
			 bytes, err2 := rawIn.MarshalJSON()
			if err2 != nil {
			 }	
			byteObj := []byte(bytes)
	*/


	//4. This formats the string and dumps it on the HTTP response writer
	byteObj := []byte(JSON_Response)
	_, err := (*ClientResponseWriter).Write(byteObj)
	if err != nil {
		R.Print("*** ERROR ***: Weird Error responding to the Client!!!")
		Y.Println(err.Error())
	}

} //end of SEND_Response_2_Client



/*
  This is the 'http service / routing handler' definition .. that is called... from here we call other "routeCOmmand(s)"
  which form the business logic. These are Controllers for the routes that are used/passed
  anyhting after the route/entry point is a "variable" that is consumed by the routeCommand
  You must ALWAYS include ?data= to trigger this.. the service will error out otherwise
*/
type GENERIC_API_ENDPOINT_HANDLER func([]URL_PARAMS) string
/*
	1. EXAMPLE of an ENDPOINT HANLDER.. Should look like this:

			func HANDLER_for_API_endpoint (URL_input []string) {

				//1. Do some business logic based on the URL_input that was passed.. ie: localhost:8020/getuser/username123

				OUTGOING_JSON = `
				[
					{
						"HelloWorld":"This is valid JSON sending back to client!!"
					}
				]
				`

				return OUTGOING_JSON
			}



	//2. Now when you startup your program, great a service endpoint TO this handler like this:

	CREATE_SERVICE_ENDPOINT("/getuser", HANDLER_for_API_endpoint) 

	//3. Now start your listener service:
	
	Start_LISTENER_SERVICE_Engine()

*/

type URL_PARAMS struct {
	KEY		string
	Value	string
}

/*
 This parses the input prameters passed.. Supports the following formats

	?name=Mary&job=librarian


*/
func PARAM_PARSER(input string, PARAMLINE string) []URL_PARAMS {

	var RESULTS []URL_PARAMS


	//1. Lets determine what format the params are in
	// if a ? is specified, we have parameters specified on the get line
	if strings.Contains(input, "=") {

		//2. If there are multiple params, we need to split on &
		if strings.Contains(input, "&") {
			ms := strings.Split(input, "&")

			for _, x := range ms {			

				sp := strings.Split(x, "=")
				var v URL_PARAMS
				v.KEY = sp[0]
				v.Value = sp[1]

				if v.KEY != "" && v.Value != "" {
					RESULTS = append(RESULTS, v)
				}
			} //end of for

		//2b. Otherwise there is a SINGLE parameter here.. Short and sweet
		} else {
			sp := strings.Split(input, "=")
			var v URL_PARAMS
			v.KEY = sp[0]
			v.Value = sp[1]

			if v.KEY != "" && v.Value != "" {
				RESULTS = append(RESULTS, v)
			}
		}
	}

	/*
		** DISABLING passing URL vars using /name/JohnDoe
			or

			/name/Valerie/age/37
 
		.. it is not reliable (and difficult to parse the router from the variables that follow it (hte mUX router reads all of them)
		BECAUSE THIS FORMAT IS INCONSISTANT 
	 3. Otherwise they have specified the parameter as a / values.. which means the FIRST value is the KEY (second param is the value).. 	
	  You can have multiples of these

		} else {

			//4. Split on /
			pp := strings.Split(PARAMLINE, "/")

			//4b, Iterate through them
			for x := 0; x < len(pp); x++ {

				item_key := pp[x]

				R.Println("KEY IS: ", item_key)

				if item_key == "" {
					x++
					continue
				}

				// a little error handling
				if strings.Contains(item_key, "favicon") || item_key == "" {
					continue
				}

				var v URL_PARAMS
				v.KEY = item_key

				//4c. The value of this item will be the very next element
				valnum := x+1

				//4d. But we need to make sure this exists in the array (or we will get an index out of range error)
				if valnum < len(pp) {
					v.Value = pp[valnum]
					x++

				//4e. IF no value is avalable, dont save this parameter item
				} else {
					continue
				}

				//5. Otherwise, save this parameter into our array			
				RESULTS = append(RESULTS, v)
				
			}
		}
	*/

	return RESULTS
} //end of func



func (RouteEntry_Handler_FUNC_to_USE GENERIC_API_ENDPOINT_HANDLER) ServeHTTP(outgoingResponseObj http.ResponseWriter, inReq *http.Request) {

	//1. Remember,the outgoingResponseObj is passed in by the MUX... we have to send a pointer to this object forward
	//  MUST ALWAYS DO THIS so the response gets sent
	
	//2. This splits the URL path and vars on the forward slash	
	REQ_PARAMS := PARAM_PARSER(inReq.URL.RawQuery, inReq.URL.Path)
	for _, p := range REQ_PARAMS {
		W.Print("\n **DEBUG PARAMS** KEY: ")
		Y.Print(p.KEY)
		W.Print(" | Value: ")
		G.Println(p.Value)


		//2b. hard coded means of turning "pretty" json on and off.. pretty shows the json however you formatted it in your handler
		// specificing ?jsonformat=clean   .. means it strips all that out
		if (p.KEY == "jsonformat" || p.KEY == "json") && p.Value == "clean" {
			JSON_CLEAN_FLAG = true
		}
	} //end of for

	
	
	//3. Send it to the appropriate "controller/handler"   .. Every handler should respond with formatted JSON
	JSON_response := RouteEntry_Handler_FUNC_to_USE(REQ_PARAMS)

	//4. Finally we send a response TO the client.. meaning we are done.HandleIncomingData
	SEND_Response_2_Client(JSON_response, &outgoingResponseObj)

} //end of http service / routing handler

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


type API_ENDPOINT_DEFINITION struct {
	API_Endpoint string
	API_Handler  GENERIC_API_ENDPOINT_HANDLER
}

var ALL_SERVICE_ENDPOINTS []API_ENDPOINT_DEFINITION

/*
	Call this to create a new REST API Endpoint /THISADDY 
*/
func CREATE_SERVICE_ENDPOINT(api_endpoint string, handler_name GENERIC_API_ENDPOINT_HANDLER) {

	var etmp API_ENDPOINT_DEFINITION
	etmp.API_Endpoint = api_endpoint
	etmp.API_Handler = handler_name

	ALL_SERVICE_ENDPOINTS = append(ALL_SERVICE_ENDPOINTS, etmp)
} //end func

var USE_PROD_MODE = false     // PROD MODE makes the listener listen on ALL interfaces (not just 127).. which is needd when running on windows locally
func Start_LISTENER_SERVICE_Engine() { 

	//1. Error handling
	slen := len(ALL_SERVICE_ENDPOINTS)

	// C.Println(" Hey Slen is: ", slen)
	if slen == 0 {
		R.Println(" ** No Entries in the ALL_SERVICE_ENDPOINTS list")
		return
	}


	//2. Formatting defualt port this way so we can use it in the MUX/ listener call
	// when running on windows local, have to explicitly specify 127 so windows firewall doesnt fuck things up
	colon_PORT := "127.0.0.1:" + DEFAULT_LISTEN_PORT

	// Otherwise, use the "prod mode" which listens on all interfaces
	if USE_PROD_MODE {
		colon_PORT = ":" + DEFAULT_LISTEN_PORT

	}	
	//1. First lets create the SERVICE object

	SERVICE_MUX_OBJ := http.NewServeMux()

	//1b. Show a hostname prefix
	hname := "http://" + colon_PORT
	if SSL_ENABLE_FLAG {
		hname = "https://" + colon_PORT
	}


	//2. Now lets create a loop that goes through the DEFINITIONS that were passed in
	// and creates http handlers handlers
	for x, ep := range ALL_SERVICE_ENDPOINTS {

		if x == 0 {
			C.Println("")
		}

		C.Print(" ** Configuring REST Endpoint for: ")
		M.Println(hname + ep.API_Endpoint)
		SERVICE_MUX_OBJ.Handle(ep.API_Endpoint, GENERIC_API_ENDPOINT_HANDLER(ep.API_Handler))

		// SERVICE_MUX_OBJ.Handle(ep.API_Endpoint, ep.API_Handler)

	} //end of for




	//3. If we are running in default mode.. we WONT be using SSL
	if SSL_ENABLE_FLAG == false {

			G.Println(" REST API is now Ready!")			
			
			err := http.ListenAndServe(colon_PORT, SERVICE_MUX_OBJ)

			if err != nil {
				R.Print("Start_LISTENER_SERVICE_Engine ERROR!! Trying to use PORT: ")
				C.Print(DEFAULT_LISTEN_PORT, "\n\n")
				M.Println(err.Error())
			}

	//5. Otherwise we start the service in SSL Mode
	} else {

		C.Print(" REST API")
		Y.Print(", USING SSL")
		C.Println(" is now Listening")

		err := http.ListenAndServeTLS(colon_PORT, SSL_CERT_PEM_FILE, SSL_KEY_FILE, SERVICE_MUX_OBJ)

		if err != nil {
			R.Print("Start_LISTENER_SERVICE_Engine IN SSL MODE ERROR!! Trying to use PORT: ")
			C.Print(DEFAULT_LISTEN_PORT, "\n\n")
			M.Println(err.Error())
		}
	}

} //end of Start_SERVICE







func init() {
  
	if USE_API_SETUP_DEFAULT_PARAMS {
		
		flag.StringVar(&DEFAULT_LISTEN_PORT, "port", DEFAULT_LISTEN_PORT,  "  The port the listener should listen on")
		flag.BoolVar(&SSL_ENABLE_FLAG,       "enableSSL", SSL_ENABLE_FLAG, "  Enables SSL support. Requires --cert and --key")
		flag.StringVar(&SSL_CERT_PEM_FILE,   "cert", SSL_CERT_PEM_FILE,    "  Full path to the CERT / PEM file for SSL (requires --enableSSL )")
		flag.StringVar(&SSL_KEY_FILE,        "key", SSL_KEY_FILE,        "    Full path to the KEY FILE for the SSL CERT (requires --enableSSL )")

		flag.BoolVar(&USE_PROD_MODE,       "restapiprod", USE_PROD_MODE, "  Enables PROD MODE Rest service.. listens on ALL interfaces not just 127.0.0.1")

	}


	//2. Error handling for the SSL stuff
	if SSL_ENABLE_FLAG {
		if SSL_CERT_PEM_FILE == "" || SSL_KEY_FILE == "" {

			R.Println("ERROR: You MUST specify the CERT and KEY file with --cert and --key")
			R.Println("       To run in SSL mode ")

			os.Exit(-9)
		}
	}

}
