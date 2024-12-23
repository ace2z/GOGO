package _CORE

import (
	"os"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

func Determine_PATHS_Engine(PARENT_REPO_NAME string, toplevel string) (string, string) {
	LOCAL_ABSOLUTE_PATH, _ := os.Getwd()
	PLACEHOLDER()

	msplit := strings.Split(toplevel, "/")
	if len(msplit) <= 0 {
		M.Println("Error: Could not determine the top level of the repo")
		return "", ""
	}
	lastEL := msplit[len(msplit)-1]

	//for loop in reverse on LOCAL_ABSOLUTE_PATH output
	csplit := strings.Split(LOCAL_ABSOLUTE_PATH, "/")

	var TMP_OFF_MOD_PATH = PARENT_REPO_NAME

	var TPATHS []string

	for i := len(csplit) - 1; i >= 0; i-- {
		tmp := csplit[i]
		//		W.Println("TMP is: **" + tmp + "**")

		lowTMP := strings.ToLower(tmp)
		lowLAST := strings.ToLower(lastEL)
		if lowTMP == lowLAST {
			break
		}
		TPATHS = append(TPATHS, tmp)
	}

	// Now go through TPATHS
	for i := len(TPATHS) - 1; i >= 0; i-- {
		TMP_OFF_MOD_PATH = TMP_OFF_MOD_PATH + "/" + TPATHS[i]
	}

	C.Println(PARENT_REPO_NAME)
	Y.Println("git rev, TopLevel: ", toplevel)
	G.Println("ABS Path: ", LOCAL_ABSOLUTE_PATH)
	W.Println("Top Level LAST Element: **" + lastEL + "** ")
	SHOW_STRUCT(TPATHS)
	Y.Println("TMP_OFF_MOD_PATH: ", TMP_OFF_MOD_PATH)
	PressAny()

	return TMP_OFF_MOD_PATH, LOCAL_ABSOLUTE_PATH
}
