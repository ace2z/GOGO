/*	APIceBerg - A bunch of SUPER lean REST and JSON helper functions to make REST dev in GO ..FUN!

	v1.54	Feb 22, 2021	- Removed SCrape Master.. moved a better version to TERRY_COMMON
	v1.52	Feb 03, 2020	- Added JSON_DOWNLOAD. Downloads a json Byte object you can read or iterate
	v1.50	Jan 06, 2020	- Rebranded just making this the generic name for this Module
	v1.45	Dec 27, 2019	- Added parameter parse support. sends them in an array you can access
	v1.40	Feb 02, 2019	- Now  This supports SSL and non-SSL (just http) mode.
				              				If you need SSL, specifiy a PEM and KEY file on the command line
							  				with SSL PEM and Key file specified

	v1.30	Nov 05, 2018	- Added support for Spinnig up a micro webserver
	v1.24	Sep 24, 2018	- Initial Release

*/

package CUSTOM_GOMOD

import (

	"flag"
	"net/http"
	"os"

	. "github.com/acedev0/GOGO/Gadgets"
)



// This is a mini/micro Web Server 
func MINI_WEB(WEBROOT string, listenPort string, START_MESSAGE string) {

	listenPort = ":" + listenPort

	webURL := "http://localhost" + listenPort

	MESSAGE := "Micro WEB Server is RUNNING!"

	if START_MESSAGE != "" {
		MESSAGE = START_MESSAGE
	}

	C.Println(" - - - -", MESSAGE)
	lastChar := WEBROOT[len(WEBROOT)-1:]
	if lastChar != "/" {
		// C.Println("\n - - - INFO: Appending / (forward slash) to WEBROOT\n")
		WEBROOT = WEBROOT + "/"
	}	

	M.Print("\n     * Listening On: ")
	G.Println(webURL)
	CLIPBOARD_COPY(webURL)
	Y.Println("     (Saved to clipboard), Just paste in your browser!")

	W.Println("\n\n  .....CTRL -C to Exit!")
	W.Println("")
	W.Println("")

	http.Handle("/", http.FileServer(http.Dir(WEBROOT)))
	http.ListenAndServe(listenPort, nil)


} //end of


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
