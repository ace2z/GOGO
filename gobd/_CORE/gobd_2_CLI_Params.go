package _CORE

import (
	"flag"
)

var TEST_MOD = false
var MAKE_MOD = false
var DO_NORMALIZE = false

var DO_GOGET_INSTEAD = false
var VERBOSE_MODE = false

var COMMIT_MESSAGE = ""
var JUST_TEST = false

var OPTIONS=""

var DO_CLEAN = false
var DO_PURGE = false
var JUST_CLEAN = false
var JUST_PURGE = false

func CLI_PARAMS_INIT() {

	// Basic admin Params
    flag.BoolVar(&TEST_MOD,  "testmod", TEST_MOD,      "  Test the current Go Module you are in (must have a go.mod file) but doesnt commit it to the repo")
	flag.BoolVar(&MAKE_MOD,  "initmod", MAKE_MOD,      "  Builds, Tests and commits module to repo (if successful)")
	flag.BoolVar(&MAKE_MOD,  "makemod", MAKE_MOD,      "  alias for initmod")	
	flag.BoolVar(&VERBOSE_MODE,  "verbose", VERBOSE_MODE,      "  Verbose Mode (more messages and errors)")	

	flag.BoolVar(&JUST_TEST,  "test", JUST_TEST,      "  Runs a basic go test on the GO program or module")	

	flag.BoolVar(&DO_CLEAN,  "clean", DO_CLEAN,      "  Runs a basic go clean")	
	flag.BoolVar(&DO_PURGE,  "purge", DO_PURGE,      "  Runs a DEEP Purge.. kills all the GO Caches")	

	flag.BoolVar(&JUST_CLEAN,  "justclean", JUST_CLEAN,      "  Runs a a clean and EXITS")	
	flag.BoolVar(&JUST_PURGE,  "justpurge", JUST_PURGE,      "  Runs a DEEP Purge..and EXITS")

	flag.BoolVar(&DO_GOGET_INSTEAD,  "usegoget", DO_GOGET_INSTEAD,      "  Instead of using 'go mod tidy', we use go get to install package dependencies (example: Golang Mongo Driver ver 1.11.5 broke go mod tidy on 5/3/2023). Use this if you have issues with go mod tidy on other modules")

	flag.StringVar(&COMMIT_MESSAGE,  "message", COMMIT_MESSAGE,      "  The message to git should use when you are publishing your GO MOD")	

	flag.StringVar(&OPTIONS,  "opts", OPTIONS,      "  Specify modifier options for the build process (see --docs)")	
}