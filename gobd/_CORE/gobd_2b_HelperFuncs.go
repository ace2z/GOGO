package _CORE

import (
	"strings"
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)

// Useful Globals
var GO_BIN = ""
var PREFIX = "   -| "

func CHECK_PreReqs() {
	if VERBOSE_MODE {
		C.Print(PREFIX, "Making sure Go is installed..  ")
	}
	result, _ := RUN_COMMAND("go version")
	if strings.Contains(result, "go version") == false {
		W.Println("")
		M.Print(PREFIX, " *** ERROR ***")
		Y.Println(" Go is NOT installed.. or is NOT in your path!")
		os.Exit(-9)
	}

	tmpComm, _ := RUN_COMMAND("which go")
	GO_BIN = strings.TrimSpace(tmpComm)

	if VERBOSE_MODE {
		G.Print("Yes! ")
		Y.Println(GO_BIN)

	
		C.Print(PREFIX, "Making sure GIT is Installed.. ")
	}
	tmpComm, _  = RUN_COMMAND("git version")
	if strings.Contains(tmpComm, "git version") == false {
		W.Println("")
		M.Print(PREFIX, " *** ERROR ***")
		Y.Println(" git command is NOT installed.. or is NOT in your path!")
		os.Exit(-9)
	}	

	if VERBOSE_MODE {
		G.Println("Yes! ")
	}

	// All the magic happens here
	GET_Proper_REPO_Paths()

	

}
