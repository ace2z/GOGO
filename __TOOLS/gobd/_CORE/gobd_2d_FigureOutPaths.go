package _CORE

import (
	"os"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

func Determine_PATHS_Engine(PARENT_REPO_NAME string, toplevel string) {
	cwd, _ := os.Getwd()
	PLACEHOLDER()

	C.Println(PARENT_REPO_NAME)
	Y.Println(toplevel)
	G.Println(cwd)

	msplit := strings.Split(toplevel, "/")
	if len(msplit) <= 0 {
		M.Println("Error: Could not determine the top level of the repo")
		return
	}
	lastEL := msplit[len(msplit)-1]
	W.Println("Last Element: ", lastEL)

	//for loop in reverse on cwd output
	csplit := strings.Split(cwd, "/")

	var OFFICIAL_MOD_PATH = PARENT_REPO_NAME

	var TPATHS []string

	for i := len(csplit) - 1; i >= 0; i-- {
		tmp := csplit[i]
		if tmp == lastEL {
			break
		}
		TPATHS = append(TPATHS, tmp)
	}
	SHOW_STRUCT(TPATHS)

	// Now go through TPATHS
	for i := len(TPATHS) - 1; i >= 0; i-- {
		OFFICIAL_MOD_PATH = OFFICIAL_MOD_PATH + "/" + TPATHS[i]
	}

	G.Println("OFFICIAL_MOD_PATH: ", OFFICIAL_MOD_PATH)

	PressAny()

	// Debug till we have all this right
	//DO_EXIT()
}