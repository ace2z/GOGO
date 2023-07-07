package CUSTOM_GO_MODULE

import (
    "os"
    "time"
    "strings"
    "encoding/json"

	//. "local/_CORE"
    
    //. "local/_DATABASE"
    . "github.com/ace2z/GOGO/Gadgets"
    //. "github.com/ace2z/GOGO/Gadgets/Mongolian"

    "github.com/PuerkitoBio/goquery"

    "crypto/tls"
    "io"
	"net/http"
	uu "net/url"    

)


var API_KEY = "f26e62fe199a53b7746d712da5c2a5fde4320703"
var COUNTRY = ""          // Check zenrows if you need another country  .. otherise select 'us'  ..or leave blank!

var TIMEOUT = 120   //60 
var RESP_ERROR_MATRIX = []string{
    "Could not get content",
    "returned a 404",
    "The concurrency limit was reached",
} 



// Returns errors the proxy might kick out if it has a problem
func found_RESP_ERROR(input string) bool {
    for _, xerr := range RESP_ERROR_MATRIX {
        if strings.Contains(input, xerr) {
            return true
        }
    }
    return false
}

// adds errors to the RESP_ERROR_MATRIX
func ADD_RESP_ERROR(newerror string) {
    RESP_ERROR_MATRIX = append(RESP_ERROR_MATRIX, newerror)
}





// This scrapes by way of using a proxy (Zenrows)... you need an account.. 
func PROXY_SCRAPE( url string, ALL_PARAMS ...interface{}) (bool, *goquery.Document, string) {
    var p_filename = ""
    var p_timeout = TIMEOUT
    var p_skipheaders = false

    var doc *goquery.Document

	//1. Get tvars passed
	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		int_val, IS_INT := param.(int)
        //bool_val, IS_BOOL := param.(bool)

        
        // For string params
        if IS_STRING && string_val != "" {

            // If this is a file.. nexxt param is a filename (with full path) to save it to
            if param == "-savefile" {
                p_filename = ALL_PARAMS[n+1].(string)
                continue
            }

            if param == "-skipheader" {
                p_skipheaders = true
                continue
            }

            continue
        }

        // If an int is passed, assume its a timeout
        if IS_INT && int_val != 8675309 {
            p_timeout = int_val
            continue
        }

	} //end of ARGS


	var country_url_section = "&proxy_country=" + COUNTRY
	if COUNTRY == "" {
		country_url_section = ""
	}
    if country_url_section == "" {}

	FULL_URL := "http://" + API_KEY + ":js_render=true&antibot=true&premium_proxy=true" + country_url_section + "@proxy.zenrows.com:8001"
    // W.Println("URL: ")
    // C.Println(FULL_URL)
    //FULL_URL := "http://f26e62fe199a53b7746d712da5c2a5fde4320703:js_render=true&antibot=true&premium_proxy=true@proxy.zenrows.com:8001"
    proxy, _ := uu.Parse(FULL_URL)

    httpClient := &http.Client{
        Timeout: time.Duration(p_timeout) * time.Second,
        Transport: &http.Transport{
            Proxy:           http.ProxyURL(proxy),
            TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        },
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
		M.Print("ERROR in PROXY_ REQUEST: ")
		Y.Println(err)
		return false, doc, ""
    }
    //otherwise. Save Main req.. Also get all the headerssaved
    var MAIN_REQ = req
	var ALL_HEADERS = ""
    if p_skipheaders == false {
        if reqHeadersBytes, err := json.Marshal(MAIN_REQ.Header); err != nil {
            Y.Println(" WARNING: Cant show request PROXY_")
            W.Println(err)
        } else {
            ALL_HEADERS = string(reqHeadersBytes)
        }
    }

    
    resp, err2 := httpClient.Do(req)
    if err2 != nil {
		M.Print("ERROR in PROXY_ RESSP ")
		Y.Println(err2)
		return false, doc, err2.Error()
    }

    // otherwise
    body, err3 := io.ReadAll(resp.Body)
    if err3 != nil {
		M.Print("ERROR in PROXY_ BODY ")
		Y.Println(err3)
		return false, doc, err3.Error()
    }    

    // Finally almost at the end. cpature the body
	BODY_TEXT := string(body)
    
    if found_RESP_ERROR(BODY_TEXT) {
        return false, doc, BODY_TEXT
    }
    
    //3. and if specified, savev the body to the filename
    if p_filename != "" {
        
        os.Remove(p_filename) 
        WRITE_FILE(p_filename, BODY_TEXT)

        //Create a empty file.. delete any existing
        /* this doesnt work 
            if FILE_EXISTS(p_filename) {
                os.Remove(p_filename)            
            }
            file, err := os.Create(p_filename)
            if err != nil {
                M.Println("PROXY_FILE_ERROR: ", err)
                return false, doc, ""
            }
            defer file.Close()

            //Write the bytes to the fiel
            _, err = io.Copy(file, resp.Body)
            if err != nil {
                M.Println("PROXY_FILE_ERROR_SAVE_to_DISK: ", err)
                return false, doc, ""
            }
        */

    }
    
	// If we get this far we are good.. call goquery
	//5. Now finally, lets create our DOM object using goquery and empty the reader into the DOM object
	doc, err2 = goquery.NewDocumentFromReader(resp.Body)
	if err2 != nil {
		M.Print("ERROR in PROXY_ GOQUERY: ")
		Y.Println(err2)
		return false, doc, BODY_TEXT
	}

    //5b. Explictly close the body
    resp.Body.Close()

    // SUCCESS!!!!. Append headers to the body_TEXT from now on    
    BODY_TEXT = ALL_HEADERS + BODY_TEXT

	return true, doc, BODY_TEXT
}



