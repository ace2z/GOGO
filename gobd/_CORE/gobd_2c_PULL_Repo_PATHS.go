package _CORE

import (
	"os"
	"strings"
	"path/filepath"
	. "github.com/ace2z/GOGO/Gadgets"
)


var REPO_BASEDIR = ""
var MOD_DIRPATH = ""
var JUST_MOD_PACKAGE_NAME = ""
var MODULE_IMPORT_NAME = "local"
var REPO_URL = ""
var PARENT_REPO_NAME = ""



func repo_delims(r rune) bool {
    return r == ' '
}

func url_delims(r rune) bool {
	return r == '@'
}
func GET_Proper_REPO_Paths() {
	cwd, _ := os.Getwd()
	

	if VERBOSE_MODE {
		C.Println("")
		C.Print(PREFIX, "Determining GIT REPO Path data...")
	}
	
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

	REPO_BASEDIR = strings.TrimSuffix(filepath.Base(res2), "\n")
	msplit := strings.Split(cwd, REPO_BASEDIR)
	MOD_DIRPATH = strings.TrimPrefix(msplit[1], "/")
	JUST_MOD_PACKAGE_NAME = filepath.Base(MOD_DIRPATH)

	// Gets the module IMPORT name based on the Github service being used
	WHAT_GIT_Service_Being_Used()

	if VERBOSE_MODE {
		C.Print(PREFIX, "REPO_URL" + ": ")
		if is_private {
			M.Print("(PRIVATE!) ")
		}
	
		Y.Println(REPO_URL)
		C.Print(PREFIX, "PARENT_REPO_NAME" + ": ")
		Y.Println(PARENT_REPO_NAME)
		C.Print(PREFIX, "REPO_BASEDIR" + ": ")
		Y.Println(REPO_BASEDIR)
	}

	if MODULE_IMPORT_NAME == "local" {
		C.Print(PREFIX, "Regular GO Program (not module)" + ": ")
		G.Println("Yes!")
	} else {
		
		C.Print(PREFIX, "MOD_DIRPATH" + ": ")
		Y.Println(MOD_DIRPATH)	
		C.Print(PREFIX, "JUST_MOD_PACKAGE_NAME" + ": ")
		Y.Println(JUST_MOD_PACKAGE_NAME)	
		C.Print(PREFIX, "MODULE_IMPORT_NAME" + ": ")
		Y.Println(MODULE_IMPORT_NAME)
	}
}

/*
	Determine which GIT service we are on:

	Currently just supporting github.com .. eventually will properly support
	- azure devops
	- bitbucket
	- gitlab

	This is due to the fact the path you use to refer to modules when you import and go get them.. DIFFERS with these services
	Github.com tends to just work
*/
func WHAT_GIT_Service_Being_Used() {

	if TEST_MOD == false && MAKE_MOD == false {
		
		MODULE_IMPORT_NAME = "local"

	// Else we determine the proper module name path based on the git service
	} else {

		if strings.Contains(PARENT_REPO_NAME, "github.com") {
			
			MODULE_IMPORT_NAME = GITHUB_Proper_GOMOD_Path_Name()

		} else if strings.Contains(PARENT_REPO_NAME, "dev.azure") {

			MODULE_IMPORT_NAME = AZURE_Proper_GOMOD_Path_Name()
		}

	}	

}
	

		