package CUSTOM_GOMOD

import (

	"flag"
	"net/http"
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)



// This is a mini/micro Web Server 
func MINI_WEB(WEBROOT string, listenPort string, START_MESSAGE string, USE_PROD_MODE bool) {

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

		//flag.BoolVar(&USE_PROD_MODE,       "restapiprod", USE_PROD_MODE, "  Enables PROD MODE Rest service.. listens on ALL interfaces not just 127.0.0.1")

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
