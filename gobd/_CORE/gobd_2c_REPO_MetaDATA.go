package _CORE

import (
	"os"
	"path/filepath"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

var MOD_LOCAL_PATH = ""
var MOD_LOCAL_BASEDIR = ""

var MOD_DIRPATH = ""

// var JUST_MOD_PACKAGE_NAME = ""
var OFFICIAL_MODULE_IMPORT_NAME = "local"
var REPO_URL = ""
var PARENT_REPO_NAME = ""

var REPO_LOCAL_ROOT = ""

func repo_delims(r rune) bool {
	return r == ' '
}

func url_delims(r rune) bool {
	return r == '@'
}
func GET_REPO_MetaDATA() {
	cwd, _ := os.Getwd()

	Y.Println("   -| Determinig REPO Meta Data..")

	result, _ := RUN_COMMAND("git remote show origin 2>&1")

	sd := strings.FieldsFunc(result, repo_delims)
	REPO_URL = strings.TrimSuffix(sd[5], "\n")

	//2. Determine if this is a a private repo.. will have an amphersand that seperate teh U:P credds from the base url
	is_private := false
	tmprd := REPO_URL
	if strings.Contains(REPO_URL, "@") {
		rd := strings.FieldsFunc(REPO_URL, url_delims)
		tmprd = rd[1]
		is_private = true
	} else {
		tmprd = strings.TrimPrefix(tmprd, "https://")
	}
	PARENT_REPO_NAME = strings.TrimSuffix(tmprd, ".git")

	res2, _ := RUN_COMMAND("git rev-parse --show-toplevel")

	//Error handling.. if we are NOT in a GIT repo, we dont do the remaining items
	if strings.Contains(res2, "not a git repository") {
		M.Println(" ( NOT in a GIT REPO )")
		return
	} else {
		if VERBOSE_MODE {
			W.Println("")
		}
	}

	//2. Need all this stuff as is and in this order.. dont change it
	MOD_LOCAL_PATH = strings.TrimSuffix(res2, "\n")
	REPO_LOCAL_ROOT := MOD_LOCAL_PATH
	MOD_LOCAL_BASEDIR = strings.TrimSuffix(filepath.Base(res2), "\n")
	msplit := strings.Split(cwd, MOD_LOCAL_BASEDIR)

	// Y.Println("msplit is: ")
	// SHOW_STRUCT(msplit)
	
	ind := 0
	if len(msplit) > 1 {
		ind = 1
	}

	MOD_DIRPATH = strings.TrimPrefix(msplit[ind], "/")

	//C.Println("Now MOD_DIRPATH is: ", MOD_DIRPATH)

	MOD_LOCAL_BASEDIR = MOD_DIRPATH

	//3. If we are in same directory as REPO LOCAL Root.. and trying to initalize a module
	// Note: right now we're only supporting GITHUB ... for Azure, Gitlab and Bitbucket, we have different paths that are
	// used for the "official module import name".. and need adjust accordingly
	if INIT_MOD or TEST_MOD{
		if REPO_LOCAL_ROOT == cwd {
			MOD_LOCAL_BASEDIR = cwd
			OFFICIAL_MODULE_IMPORT_NAME = PARENT_REPO_NAME
		} else {
			OFFICIAL_MODULE_IMPORT_NAME = PARENT_REPO_NAME + "/" + MOD_LOCAL_BASEDIR
		}

		//3b Cleanup
		OFFICIAL_MODULE_IMPORT_NAME = strings.Replace(OFFICIAL_MODULE_IMPORT_NAME, "//", "/", -1)
		// Remove lsat character if there is an extra /
		OFFICIAL_MODULE_IMPORT_NAME = strings.TrimSuffix(OFFICIAL_MODULE_IMPORT_NAME, "/")
	}

	// Gets the module IMPORT name based on the Github service being used
	//WHAT_GIT_Service_Being_Used()

	if VERBOSE_MODE {
		C.Print(PREFIX, "REPO_URL"+": ")
		if is_private {
			M.Print("(PRIVATE!) ")
		}

		Y.Println(REPO_URL)
		C.Print(PREFIX, "REPO_LOCAL_ROOT"+": ")
		Y.Println(REPO_LOCAL_ROOT)
		C.Print(PREFIX, "PARENT_REPO_NAME"+": ")
		Y.Println(PARENT_REPO_NAME)
		C.Print(PREFIX, "MOD_LOCAL_BASEDIR"+": ")
		Y.Println(MOD_LOCAL_BASEDIR)
	}

	if OFFICIAL_MODULE_IMPORT_NAME == "local" {
		C.Print(PREFIX, "Regular GO Program (not module)"+": ")
		G.Println("Yes!")
	} else {

		C.Print(PREFIX, "Official MODULE Name"+": ")
		G.Println(OFFICIAL_MODULE_IMPORT_NAME)
	}

	// Debug till we have all this right
	//DO_EXIT()
}
