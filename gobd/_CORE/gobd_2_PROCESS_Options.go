package _CORE

import (
	"strings"
	"os"
	"path/filepath"

	. "github.com/ace2z/GOGO/Gadgets"
)

var ARCH=""			// can be: arm64, or amd64
var OUTFILE=""
var DESTDIR=""
var BIN_TYPE=""		// can be: windows, linux or darwin
var VERSION_to_USE = ""
var FULL_DEST_FILE = ""
var JUST_FILE = ""

var EXTENSIONS = map[string]interface{} {
	"mac" : ".mac",
	"windows" : ".exe",
	"linux" : ".linux",
	"mac_arm" : ".mac",
	"linux_arm" : ".linux",
}	


func SET_VERSION_from_COMMIT() {
	// Make sure we are in a git repo
	result, _ := RUN_COMMAND("git status")

	if strings.Contains(result, "not found") || strings.Contains(result, "not a git repo") {
		VERSION_to_USE = "v1.2.3-beta"

	// Else.. lets get the version
	} else {
		commit_ver, _ := RUN_COMMAND("git rev-parse --short=6 HEAD")
		VERSION_to_USE = strings.TrimSuffix(commit_ver, "\n")
	}

}
func Set_DEFAULTS_based_on_PLATFORM() {

	if CURRENT_OS == "windows" {
		BIN_TYPE = "windows"

	} else if CURRENT_OS == "mac" {
		BIN_TYPE = "darwin"
	} else if CURRENT_OS == "linux" {
		BIN_TYPE = "linux"
	}
	ARCH = CURRENT_ARCH

	// Get the current directory name.. This will be the name of the binary
	cwd, _ := os.Getwd()
	tmp_justfile := filepath.Base(cwd)	
	FULL_DEST_FILE = tmp_justfile
	JUST_FILE = tmp_justfile
	
	// Set the version based on the GIT Commit (if we are in a git repo.. )
	SET_VERSION_from_COMMIT()
}


func GET_EXTENSION_for_FILE() {

	F_EXT := ""
	if BIN_TYPE == "windows" {
		F_EXT = EXTENSIONS["windows"].(string)

	} else if BIN_TYPE == "darwin" {
		F_EXT = EXTENSIONS["mac"].(string)
		if ARCH == "arm64" {
			F_EXT = EXTENSIONS["mac_arm"].(string)
		}		
	} else if BIN_TYPE == "linux" {
		F_EXT = EXTENSIONS["linux"].(string)
		if ARCH == "arm64" {
			F_EXT = EXTENSIONS["linux_arm"].(string)
		}
	}

	FULL_DEST_FILE = FULL_DEST_FILE + F_EXT
}


func opts_delims(r rune) bool {
    return r == '=' || r == ','
}
func PROCESS_OPTIONS () {
	// Find out what platform and arch we are on.. and set the global flags
	Set_DEFAULTS_based_on_PLATFORM()

	var DONT_ADD_EXTENSION = false
	var USE_DEFAULTS = true
	

	sd := strings.FieldsFunc(OPTIONS, opts_delims)
	for n, ropt := range sd {
		nn := n + 1

		VAL := ""
		if nn < len(sd) {
			VAL = sd[nn]
		}

		switch ropt {
			case "ver":
				VERSION_to_USE = VAL
			case "noext":
				DONT_ADD_EXTENSION = true
				USE_DEFAULTS=false
			case "dest":
				DESTDIR=VAL
				USE_DEFAULTS=false
			case "outfile":
				FULL_DEST_FILE=VAL
				JUST_FILE=VAL
				USE_DEFAULTS=false
			case "arch":
				ARCH=VAL
				// a little helper to allow arm and intel
				if ARCH=="arm" {
					ARCH="arm64"
				} else if ARCH=="intel" {
					ARCH="amd64"
				}
			case "platform":
				if VAL=="windows" {
					BIN_TYPE="windows"
					ARCH="amd64"		//as of 05/2023 .. Windows arm isnt a thing.. just force this to intel for convenience
				} else if (VAL=="mac" || VAL=="darwin") {
					BIN_TYPE="darwin"
				} else if VAL=="linux" {
					BIN_TYPE="linux"
				}
		}
	}

	//2. We need to do some fixing for the FILENAME if the defaults have been altered
	if USE_DEFAULTS == false {

		if DESTDIR != "" {
			FULL_DEST_FILE = DESTDIR + "/" + FULL_DEST_FILE
			
			// If this is windows.. we replace the / with \\
			if CURRENT_OS == "windows" {
				DESTDIR = strings.Replace(DESTDIR, "/", "\\\\", -1 )
				FULL_DEST_FILE = strings.Replace(FULL_DEST_FILE, "\\", "\\\\",-1 )
				FULL_DEST_FILE = strings.Replace(FULL_DEST_FILE, "/", "\\\\", -1 )
			}

			// Safety make dir
			os.MkdirAll(DESTDIR, 0777)
		}

		//Check JUST_FILE. If there is ever an extension in the file...we never add an extension
		if strings.Contains(JUST_FILE, ".") {
			DONT_ADD_EXTENSION = true
		}

		//
		if DONT_ADD_EXTENSION == false {
			GET_EXTENSION_for_FILE()
		}

	// Else if we are using the setup defaults
	} else {

		GET_EXTENSION_for_FILE()
	}

	//4. Finally.. safety REMOVE the DEST_FILE
	os.Remove(FULL_DEST_FILE)
}
