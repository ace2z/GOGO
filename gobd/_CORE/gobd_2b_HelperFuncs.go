package _CORE

import (
	"strings"
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)

var CURR_DIR_is_SYMLINK = false



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


	/*
		2. Determine if the current directory is actually a SYMLINK.
		This breaks go mod init and tidy and you get errors about you rimports.
		EXAMPLE:

		go build
		_CORE/hp_2c_PROXY_Scraper.go:15:5: no required module provides package github.com/PuerkitoBio/goquery; to add it:
			go get github.com/PuerkitoBio/goquery
		hp_1.go:6:5: no required module provides package github.com/ace2z/GOGO/Gadgets; to add it:
			go get github.com/ace2z/GOGO/Gadgets
		_CORE/hp_2_Core_COMMON.go:11:2: no required module provides package github.com/ace2z/GOGO/Gadgets/MDC; to add it:
			go get github.com/ace2z/GOGO/Gadgets/MDC		
	*/

	if HAVE_LINK(CURR_DIR())  {
		BW.Print(" = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")
		W.Println("")
		BW.Print(" NOTE: The current directory is a SYMBOLIC LINK                  ")
		W.Println("")
		BW.Print(" GO (go mod init / go build ) doesnt work properly when you are  ")
		W.Println("")
		BW.Print(" building from a SYMLINK directory, and you may see              ")
		W.Println("")
		BW.Print(" some 'no required module provides package' errors               ")
		W.Println("")		
		BW.Print(" To fix this, I am switching to GOGETMODE. So, Instead of        ")
		W.Println("")				
		BW.Print(" using 'go mod tidy' I will retrieve all go imports using        ")
		W.Println("")
		BW.Print(" 'go get [MODULE]' manually                                      ")
		W.Println("")
		BW.Print(" = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = ")
		W.Println("")

		W.Println("")
		
		DO_GOGET_INSTEAD = true
	}

	// All the magic happens here
	GET_Proper_REPO_Paths()

	

}
