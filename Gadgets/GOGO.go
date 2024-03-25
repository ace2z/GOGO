/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier

------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter. (This is a GO Thing)

	Nov 24, 2022	- MAJOR CLEANUP AND RIVSIONS

	Oct 22, 2022	- Cleanup and revisions
	Aug 27, 2021	- Ripped out a bunch of stuff to make this smaller. They are
	Jun 05, 2014    - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (
	"flag"
	"math"
	"math/rand"
	"os"
	"os/exec"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/google/uuid"
	//mini "github.com/janeczku/go-spinner"
)

/*
- - - -
- - - -
- - - - START OF GLOBALS WE NEED - - - - - -
- - - -
- - - -
*/
var GLOBAL_PREFIX = " ==| "

var PROG_START_TIME string
var PROG_START_TIMEOBJ time.Time

var GLOBAL_CURR_DATE = "" // Current Actual Date in the Timezone we specified

var CHECK_VERSION = false
var DEBUG_MODE = false

// Generic null ints and floats to use .. when using just 0 isnt sufficient
var NULL_INT = -69696969
var NULL_FLOAT = -69.69696969

func NEW_UUID() string {

	return uuid.New().String()
}

// If first is more recent than second, we return true
func MOST_RECENT_DATE(first time.Time, second time.Time) bool {

	if first.After(second) {
		return true
	}

	return false
}

// var s *mini.Spinner
// func MINI_SpinStart() {
// 	s = mini.StartNew("")
// }
// func MINI_SpinSTOP() {
// 	s.Stop()
// }

type IP struct {
	Query string
}

// Easy way to get your public IP
func GET_PUBLIC_IP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func getProcessOwner() string {
	stdout, err := exec.Command("ps", "-o", "user=", "-p", strconv.Itoa(os.Getpid())).Output()
	if err != nil {
		R.Println("Error with AM_I_ROOT_ getProcess Owner", err)
		os.Exit(1)
	}

	//Y.Println(" RETURNED is", string(stdout))
	return string(stdout)
}
func AM_I_ROOT() bool {

	result := getProcessOwner()
	if strings.Contains(result, "root") {
		return true
	}

	return false

}

// Easy way to quick exit a program without having to remember to import os
func DO_EXIT(EXTRA_ARGS ...string) {

	var verbose = false

	//1. Parse out EXTRA_ARGS
	for _, VAL := range EXTRA_ARGS {

		if VAL != "" {
			if VAL == "-verbose" {
				verbose = true
			}
		}

	} // end of for

	if verbose {
		SHOW_BOX("|red| DO_EXIT of Program!")
	}
	os.Exit(-9)
}

// alias
func DOEXIT(EXTRA_ARGS ...string) {
	DO_EXIT(EXTRA_ARGS...)
}

/*
	TRIM_FIRST  not needed.. use
	TrimSuffix  and TrimPrefix instead
*/

func IS_ODD_Num(inum int) bool {
	var result float64

	// dividing the said number by 2 and storing the result in a variable
	result = math.Mod(float64(inum), 2)

	if result != 0 {
		return true
	}
	return false
}

// This makes sure we're running as run.. program exists otherwise
func MAKE_Sure_Running_As_ROOT() {
	if AM_I_ROOT() == false {

		M.Println(" WARNING: You HAVE to use sudo to run this program!")
		C.Println("")
		C.Println("")
		DOEXIT("silent")
	}

}

func MARK_START_TIME() time.Time {
	_, START_TIMEOBJ := GET_CURRENT_TIME()

	return START_TIMEOBJ
}

// Takes start_TIME and gets the total time elapsed and displays a status
func MARK_END_TIME(start_TIME time.Time) {

	endTime, endOBJ := GET_CURRENT_TIME()
	difftemp := endOBJ.Sub(start_TIME)
	TIME_DIFF := difftemp.String()

	Y.Println("\n\n ****************************************************** ")

	W.Print("              Start Time:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("                End Time:")
	M.Println(" " + endTime)
	C.Print("      Total PROGRAM DURATION: ")
	G.Println(" ", TIME_DIFF)
	C.Println("******************************************************")
}

func GET_CURRENT_TIME(ALL_PARAMS ...string) (string, time.Time) {
	var LOCAL_Location_OBJ, _ = time.LoadLocation("Local")
	var EST_Location_OBJ, _ = time.LoadLocation("America/New_York")
	var CST_Location_OBJ, _ = time.LoadLocation("America/Chicago")     // aka CST	}
	var MST_Location_OBJ, _ = time.LoadLocation("America/Denver")      // MDT / Mountain Standard
	var PST_Location_OBJ, _ = time.LoadLocation("America/Los_Angeles") // aka PST
	var UTC_Location_OBJ, _ = time.LoadLocation("UTC")

	var TZ_OBJECT = LOCAL_Location_OBJ

	var output_format = "basic"
	// Get the Prams
	for _, param := range ALL_PARAMS {

		// If they specify a timezone to use
		if param == "est" {
			TZ_OBJECT = EST_Location_OBJ
			continue
		}
		if param == "cst" {
			TZ_OBJECT = CST_Location_OBJ
			continue
		}
		if param == "mst" || param == "mdt" {
			TZ_OBJECT = MST_Location_OBJ
			continue
		}
		if param == "pst" || param == "pdt" {
			TZ_OBJECT = PST_Location_OBJ
			continue
		}
		if param == "utc" {
			TZ_OBJECT = UTC_Location_OBJ
			continue
		}

		// Otherwise assume they are specifying the OUTPUT format to SHOW_PRETTY (defaults to basic)
		// First paramn is always the STRUCT
		output_format = param
	}

	//1. Default ot the local machines time zone

	dateOBJ := time.Now().In(TZ_OBJECT)
	result, _ := SHOW_PRETTY_DATE(dateOBJ, output_format)

	return result, dateOBJ

} //end of func

// Super Useful to show structures... ANY Structure .. in JSON Format
func SHOW_STRUCT(ALL_PARAMS ...interface{}) {

	var tmpSTRUCT interface{}

	var COLOR = "yellow" // whill show struct in yellow

	// Collects the input params specified... supports INT and FLOAT dynamically
	for n, param := range ALL_PARAMS {

		// First paramn is always the STRUCT
		if n == 0 {
			tmpSTRUCT = param
			continue

		}

		// If second paramis sepectified.. This will be the color in whichw e display thie struct
		if n == 1 {
			COLOR = param.(string)
		}
	}

	if COLOR == "red" {
		R.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	} else if COLOR == "magenta" {
		M.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	} else if COLOR == "green" {
		G.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	} else if COLOR == "white" {
		W.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	} else if COLOR == "cyan" {
		C.Println(PRETTY_STRUCT_json(tmpSTRUCT))
		// The default color
	} else {
		Y.Println(PRETTY_STRUCT_json(tmpSTRUCT))
	}
}

var SPINNER_SPEED = 100
var SPINNER_CHAR = 4
var spinOBJ = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

// Creates a cool "im busy right now" status spinner so you know the program is running
func START_Spinner() {

	sduration := time.Duration(SPINNER_SPEED)

	spinOBJ = spinner.New(spinner.CharSets[SPINNER_CHAR], sduration*time.Millisecond)
	spinOBJ.Start()
}

func STOP_Spinner() {

	spinOBJ.Stop()
}

// this is a simple sleep function
func Sleep(seconds int, ALL_PARAMS ...bool) {

	var showOutput = false

	for x, BOOL_VAL := range ALL_PARAMS {

		//1. First Param is allthat is used
		if x == 0 {
			showOutput = BOOL_VAL
			continue
		}
	} // end of for

	if showOutput == true {
		secText := ""
		suffix := "seconds"
		sectemp := seconds

		if seconds >= 119 {
			sectemp = seconds / 60
			suffix = "minutes"
		}
		secText = strconv.Itoa(sectemp)
		C.Println("        ** Sleeping for: "+secText+" ", suffix, "...")
	}

	duration := time.Duration(seconds) * time.Second
	time.Sleep(duration)

} //end of sleep function

// a simple placeholder.. if you arent using anything in GOGO Gadgets main (at the moment)
// this will prevent you getting a "imported but not used" error
func PLACE_HOLDER() {
}
func PLACEHOLDER() {

}

// We need to determine what the CURRENT running platform is
var ON_LINUX = false
var ON_WINDOWS = false
var ON_MAC = false
var ON_ARM = false
var USING_ARM = false
var CURRENT_OS = ""
var CURRENT_ARCH = ""

func DETERMINE_Current_OS_and_PLATFORM() {

	// First lets run uname. MAC and Linux always have this command
	// so if it comes back as BLANK.. we know we are on windows
	output, _ := RUN_COMMAND("uname -a")
	res_out := strings.ToLower(output)

	if strings.Contains(res_out, "darwin") {
		ON_MAC = true
		CURRENT_OS = "MAC"
	} else if strings.Contains(res_out, "linux") {
		ON_LINUX = true
		CURRENT_OS = "LINUX"

		// Otherwise.. this is WINDOWS!!
	} else {
		ON_WINDOWS = true
		CURRENT_OS = "Windows"
	}

	// determine platform.. arm or amd
	CURRENT_ARCH = "Intel-AMD"
	if strings.Contains(res_out, "aarch64") || strings.Contains(res_out, "arm64") {
		ON_ARM = true
		USING_ARM = true
		CURRENT_ARCH = "ARM"
	}
}

// This gets the platform we are running on (mac, linux, windows)
func GET_BINARY_BUILT_FOR() (string, string) {
	OSTYPE := ""
	ARCH_TYPE := ""

	if runtime.GOOS == "linux" {
		OSTYPE = "LINUX"

		//2. Otherwise see if this is MAC
	} else if runtime.GOOS == "darwin" {
		OSTYPE = "MAC"

		//3. otherwise.. its windows.. it wins by default!!
	} else if runtime.GOOS == "windows" {
		OSTYPE = "Windows"

		//4. If we get this far, means we have some weird unrecognizable OS:
	} else {
		OSTYPE = "- - UNKNOWN OS - -"
	}

	if strings.Contains(runtime.GOARCH, "amd64") {
		ARCH_TYPE = "Intel-AMD"

		// Otherwise this was built for ARM
	} else {
		ARCH_TYPE = "ARM"
	}

	return OSTYPE, ARCH_TYPE

} //end of getOsType

// Returns a randomly generated number within a given range (returns a STRING AND an int)
func GenRandomRange(min int, max int) (int, string) {

	resultNum := rand.Intn(max-min) + min
	resultText := strconv.Itoa(resultNum)

	// Always return a string with a 0 prefix
	if resultNum < 10 {
		resultText = "0" + resultText
	}

	return resultNum, resultText

} //end of genRandomRange

func SHOW_PROG_VERSION(vernum string) {
	C.Print("     ")
	BW.Print("                 ")
	C.Println("")
	C.Print("     ")
	MW.Print(" Version: ")
	BW.Print(vernum, " ")
	C.Println("")
	C.Print("     ")
	BW.Print("                 ")
	W.Println("")
	W.Println("")
	DO_EXIT()
}

/*
  MUST ALWAYS CALL THIS in the MAIN of every program..
  This is how command line params get initted
   Also make sure it is the LAST 'init' type function called (for example.. BEFORE AWS_INIT
*/

func MASTER_INIT(PROGNAME string, VERSION string, ALL_PARAMS ...string) {
	var verbose = false

	var VERSION_NUM = VERSION
	var SHOW_ARCH = false

	//2. Now, see if version was passed
	for _, VAL := range ALL_PARAMS {

		if VAL == "-verbose" {
			verbose = true
			continue
		}
		if VAL == "-showarch" {
			SHOW_ARCH = true
			continue
		}
	} //end of for

	//3. Do a flag parse.. in the program itself, your params should be declared BEFORe you call master_INIT
	flag.BoolVar(&CHECK_VERSION, "version", CHECK_VERSION, "  Check Version ")
	flag.BoolVar(&DEBUG_MODE, "debug", DEBUG_MODE, "  Enables a global DEBUG_MODE flag useful for debug")
	flag.Parse()

	//3b. For Showing the version

	if CHECK_VERSION {
		SHOW_PROG_VERSION(VERSION_NUM)
	}

	//3. And, Always init the random number seeder
	rand.Seed(time.Now().UTC().UnixNano())

	//3b. And get current OS Data
	DETERMINE_Current_OS_and_PLATFORM()
	build_OS, build_ARCH := GET_BINARY_BUILT_FOR()

	//4. Setup the prog start time globals
	PROG_START_TIME, PROG_START_TIMEOBJ = GET_CURRENT_TIME()
	GLOBAL_CURR_DATE = PROG_START_TIME

	if verbose {

		if VERSION_NUM == "" {
			SHOW_BOX(PROGNAME, "|cyan|Current OS: "+CURRENT_OS+","+CURRENT_ARCH)

		} else {
			SHOW_BOX(PROGNAME, "|green|ver: "+VERSION_NUM, "|bluewhite| Current OS: "+CURRENT_OS+","+CURRENT_ARCH, " |redyellow| ( compiled For: "+build_OS+","+build_ARCH+" ) ")
		}

	} else {
		if SHOW_ARCH && VERSION_NUM != "" {
			SHOW_BOX(PROGNAME, "|green|ver: "+VERSION_NUM, "|bluewhite| Current OS: "+CURRENT_OS+","+CURRENT_ARCH, " |redyellow| ( compiled For: "+build_OS+","+build_ARCH+" ) ")

		} else if VERSION_NUM != "" {
			SHOW_BOX(PROGNAME, "|bluewhite|ver: "+VERSION_NUM)

		} else {
			SHOW_BOX(PROGNAME)
		}

	}
	W.Println("")
} //end of func

// Shows the amount of time a program ran (and start and end time)
func Show_Total_RUNTIME() {

	endTime, endOBJ := GET_CURRENT_TIME()
	difftemp := endOBJ.Sub(PROG_START_TIMEOBJ)
	TIME_DIFF := difftemp.String()

	Y.Println("\n\n ****************************************************** ")

	W.Print("              Start Time:")
	B.Println(" " + PROG_START_TIME)
	Y.Print("                End Time:")
	M.Println(" " + endTime)
	C.Print("      Total PROGRAM DURATION: ")
	G.Println(" ", TIME_DIFF)
	C.Println("******************************************************")
}

/*
	Kept here as filler /template /example / Reference

.. anything you put in this function will run  when the module is imported
*/
func init() {

	// Add stuff here

} // end of main
