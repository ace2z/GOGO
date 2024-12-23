package _CORE

import (
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)

func Determine_PATHS_Engine(PARENT_REPO_NAME string, revout string) {
	cwd, _ := os.Getwd()
	PLACEHOLDER()

	C.Println(PARENT_REPO_NAME)
	Y.Println(revout)
	G.Println(cwd)
	PressAny()
	// Debug till we have all this right
	//DO_EXIT()
}
