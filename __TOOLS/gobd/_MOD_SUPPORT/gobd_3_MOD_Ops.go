package _MOD_SUPPORT

import (
	"os"
	"strings"

	//"io/ioutil"
	. "local/_CORE"

	. "github.com/ace2z/GOGO/Gadgets"
)

func QUIK_COMMIT() {

	G.Println(PREFIX, "COMMITING To GIT Repo...")
	_, date_obj := GET_CURRENT_TIME()
	commdate := SHOW_PRETTY_DATE(date_obj, "timestamp_nozone_noday")

	MESSAGE := "Auto QuickCommit for GOMOD " + commdate
	if COMMIT_MESSAGE != "" {
		MESSAGE = COMMIT_MESSAGE
	}

	// Replace spaces in the message with _ .. for some reason git commit running this way has a problem with them
	MESSAGE = strings.Replace(MESSAGE, " ", "_", -1)

	RUN_COMMAND("git add -A")
	RUN_COMMAND("git add .")
	RUN_COMMAND("git add --all")
	RUN_COMMAND("git add -u")
	COMMIT_COMM := "git commit -m '" + MESSAGE + "' ."
	RUN_COMMAND(COMMIT_COMM)
	RUN_COMMAND("git push -f", "-showoutput")

}

func RUN_GO_Test() bool {
	SHOW_BOX(" TESTING Go MODULE/Program")

	result, _, _ := RUN_COMMAND("go test")

	if strings.Contains(result, "no test files") || (strings.Contains(result, "PASS") && strings.Contains(result, "ok")) {
		G.Println(PREFIX, "GO Test Compile SUCCESS!!")
		W.Println("")

		return true

	} else {
		W.Println("")
		W.Println(result)
		M.Println(PREFIX, "*** ERROR *** ")
		W.Print(PREFIX, "Compile Test: ")
		Y.Println(" FAILED!!")
		W.Println("")

	}

	return false
}

func DO_GOMOD_Init() {
	C.Print(PREFIX, "Running: ")
	W.Print("go mod init ")
	Y.Println(OFFICIAL_MODULE_IMPORT_NAME)
	os.Remove("go.mod")
	os.Remove("go.sum")
	RUN_COMMAND("go mod init " + OFFICIAL_MODULE_IMPORT_NAME)
	// RUN_COMMAND("go mod tidy")
}

// Takes care of manageming modules.. Uses either go mod tidy... (or go get if --usegoget is specified
func GOMOD_Dependency_Engine() {

	//1. Clear mod cahe and do DO_GOMOD_Init
	DO_GOMOD_Init()

	var dont_tidy = false
	//2. Then, Run either go mod tidy or go get
	if DO_GOGET_INSTEAD {

		SHOW_BOX("Running", "|yellow|go get -u", "to find all Dependencies")
		START_Spinner()

		var GO_FILES = FIND_all_GO_Files()

		// Now extract all the import statements from all the go Files
		var IMPORT_LIST = EXTRACT_Import_STATEMENTS(GO_FILES)

		// This will run go get for each found inmport statement and update go.mod
		RUN_GOGET_for_IMPORTS(IMPORT_LIST)
		dont_tidy = true
	}

	//3. Run go mod tidy if needed
	if dont_tidy == false {
		SHOW_BOX("Running go mod tidy to find DEPENDENCIES")
		START_Spinner()
		if VERBOSE_MODE {
			RUN_COMMAND("go mod tidy", "-showoutput")
		} else {
			RUN_COMMAND("go mod tidy")
		}
	}

	//4. Otherwise.. lets look for a PINNED_VERSIONS file. If it exists, we go through go.mod and replace anything that matches a PINNED VERSION with that version in go.mod
	CHECK_for_PINNED_VERSIONS_Engine()

	STOP_Spinner()
}

func GOMOD_Core_Ops() {

	GOMOD_Dependency_Engine()

	if RUN_GO_Test() == false {
		DOEXIT()
	}

	if TEST_MOD {
		DOEXIT()
	}

	// IF they are also building this module for pushing to the repo..
	SHOW_BOX(" BUILDING and PUBLISHING GO Module")

	QUIK_COMMIT()

}
