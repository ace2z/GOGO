package _CORE

import (
	"os"
	"strings"
	"path/filepath"
	. "github.com/ace2z/GOGO/Gadgets"
)

var MOD_LOCAL_PATH = ""
var MOD_LOCAL_BASEDIR = ""

var MOD_DIRPATH = ""
var JUST_MOD_PACKAGE_NAME = ""
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

	MOD_LOCAL_PATH = strings.TrimSuffix(res2, "\n")	
	REPO_LOCAL_ROOT := MOD_LOCAL_PATH
	MOD_LOCAL_BASEDIR = strings.TrimSuffix(filepath.Base(res2), "\n")
	msplit := strings.Split(cwd, MOD_LOCAL_BASEDIR)
	MOD_DIRPATH = strings.TrimPrefix(msplit[1], "/")
	JUST_MOD_PACKAGE_NAME = filepath.Base(MOD_DIRPATH)


	

	curr_wd, _ := os.Getwd()
	C.Println("REPO LOCAL ROOT: ", REPO_LOCAL_ROOT)
	C.Println(" CWD: ", curr_wd)
	//C.Println("MOD_LOCAL_BASEDIR: ", MOD_LOCAL_BASEDIR, " and", MOD_DIRPATH, "|", MOD_LOCAL_PATH, "|", curr_wd)

	if curr_wd != REPO_LOCAL_ROOT {
		tmp_wd := strings.Replace(curr_wd, MOD_LOCAL_PATH, "", -1)
		MOD_LOCAL_PATH = MOD_LOCAL_PATH + tmp_wd
	}





	// Gets the module IMPORT name based on the Github service being used
	WHAT_GIT_Service_Being_Used()


	if strings.Contains(OFFICIAL_MODULE_IMPORT_NAME, "local") == false {

		OFFICIAL_MODULE_IMPORT_NAME = OFFICIAL_MODULE_IMPORT_NAME + tmp_wd
		OFFICIAL_MODULE_IMPORT_NAME = strings.Replace(OFFICIAL_MODULE_IMPORT_NAME, "//", "/", -1)

		// Remove lsat character if there is an extra /
		OFFICIAL_MODULE_IMPORT_NAME = strings.TrimSuffix(OFFICIAL_MODULE_IMPORT_NAME, "/")
	}



	//3. Fix.. If we are in the ROOT dir of the repo.. We will change the offical module name accordingly
	// We check this by looking at MOD_DIRPATH.. if it is blank.. means we are in ROOT

	IN_REPO_ROOT := false
/*
	if MOD_DIRPATH == "" {
		IN_REPO_ROOT = true
	}

	//C.Println(" Before: ", OFFICIAL_MODULE_IMPORT_NAME, "cwd: ", cwd, MOD_DIRPATH, "basedir: ", MOD_LOCAL_BASEDIR, "LOCAL: ", MOD_LOCAL_PATH)

	if IN_REPO_ROOT {
		OFFICIAL_MODULE_IMPORT_NAME = strings.TrimSuffix(OFFICIAL_MODULE_IMPORT_NAME, "/")		
	}


	//4. Now.. lets check to see if we are in the ROOT of the repo.. or a subdirectory of the repo
	if MOD_LOCAL_PATH != cwd && BUILD_BASIC_GO_PROGRAM == false {
		tmp_basedir := filepath.Base(cwd)
		//Y.Println("tmpbasedir: ", tmp_basedir)
		OFFICIAL_MODULE_IMPORT_NAME += "/" + tmp_basedir
	}
*/

	if VERBOSE_MODE {
		C.Print(PREFIX, "REPO_URL" + ": ")
		if is_private {
			M.Print("(PRIVATE!) ")
		}
	
		Y.Println(REPO_URL)
		C.Print(PREFIX, "REPO_LOCAL_ROOT" + ": ")
		Y.Println(REPO_LOCAL_ROOT)
		C.Print(PREFIX, "PARENT_REPO_NAME" + ": ")
		Y.Println(PARENT_REPO_NAME)
		C.Print(PREFIX, "MOD_LOCAL_PATH" + ": ")
		if IN_REPO_ROOT {
			Y.Println(MOD_LOCAL_PATH)
		} else {
			Y.Println(MOD_LOCAL_PATH + "/" + MOD_DIRPATH)
			
		}
	}

	if OFFICIAL_MODULE_IMPORT_NAME == "local" {
		C.Print(PREFIX, "Regular GO Program (not module)" + ": ")
		G.Println("Yes!")
	} else {
		
		C.Print(PREFIX, "Official MODULE Name" + ": ")
		G.Println(OFFICIAL_MODULE_IMPORT_NAME)
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
		
		OFFICIAL_MODULE_IMPORT_NAME = "local"

	// Else we determine the proper module name path based on the git service
	} else {

		if strings.Contains(PARENT_REPO_NAME, "github.com") {
			
			OFFICIAL_MODULE_IMPORT_NAME = GITHUB_Proper_GOMOD_Path_Name()

		} else if strings.Contains(PARENT_REPO_NAME, "dev.azure") {

			OFFICIAL_MODULE_IMPORT_NAME = AZURE_Proper_GOMOD_Path_Name()
		}

	}	

}
	

		