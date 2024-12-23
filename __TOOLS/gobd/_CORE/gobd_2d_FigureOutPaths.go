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
	C.Println("LastELE: ", lastEL)

	for _, el := range msplit {
		C.Println("ELE: ", el)

	}

	PressAny()

	// Debug till we have all this right
	//DO_EXIT()
}
