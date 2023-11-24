package _CORE

import (
	"flag"
)

var BUILD_BASIC_GO_PROGRAM = true
var TEST_MOD = false
var INIT_MOD = false
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


// These are build overrides
var BUILD_MAC = false
var BUILD_LINUX = false
var BUILD_WIN = false

var BUILD_ARM = false
var BUILD_INTEL = false

var NOEXT = false

var ALT_OUTPUT = ""
func CLI_PARAMS_INIT() {

	// Basic admin Params
    flag.BoolVar(&TEST_MOD,  "testmod", TEST_MOD,      "  Test the current Go Module you are in (must have a go.mod file) but doesnt commit it to the repo")
	flag.BoolVar(&INIT_MOD,  "initmod", INIT_MOD,       "  Initializes New Go Module, Tests and commits module to repo (if successful)")
	flag.BoolVar(&INIT_MOD,  "initmods", INIT_MOD,      "  Initializes New Go Module, Tests and commits module to repo (if successful)")
	flag.BoolVar(&VERBOSE_MODE,  "verbose", VERBOSE_MODE,      "  Verbose Mode (more messages and errors)")	

	flag.BoolVar(&JUST_TEST,  "test", JUST_TEST,      "  Runs a basic go test on the GO program or module")	

	flag.BoolVar(&DO_CLEAN,  "clean", DO_CLEAN,      "  Runs a basic go clean")	
	flag.BoolVar(&DO_PURGE,  "purge", DO_PURGE,      "  Runs a DEEP Purge.. kills all the GO Caches")	

	flag.BoolVar(&JUST_CLEAN,  "justclean", JUST_CLEAN,      "  Runs a a clean and EXITS")	
	flag.BoolVar(&JUST_PURGE,  "justpurge", JUST_PURGE,      "  Runs a DEEP Purge..and EXITS")

	flag.BoolVar(&DO_GOGET_INSTEAD,  "usegoget", DO_GOGET_INSTEAD,      "  Instead of using 'go mod tidy', we use go get to install package dependencies (example: Golang Mongo Driver ver 1.11.5 broke go mod tidy on 5/3/2023). Use this if you have issues with go mod tidy on other modules")

	flag.StringVar(&COMMIT_MESSAGE,  "message", COMMIT_MESSAGE,      "  The message to git should use when you are publishing your GO MOD")	

	// = = = LOTS OF OPTIONS
	flag.StringVar(&OPTIONS,  "opts", OPTIONS,      "  Specify modifier options for the build process (see --docs)")
	flag.StringVar(&ALT_OUTPUT,  "o", ALT_OUTPUT,      "  Specify an alternate FILENAME.. or DIRECTORY where the binary should be exported to")	
	flag.StringVar(&ALT_OUTPUT,  "out", ALT_OUTPUT,      "  alias for --o")	

	flag.BoolVar(&NOEXT,  "noext", NOEXT,      "  DOES NOT add an EXTENSION when generating the OUTPUT file")

	// Quick opts for building linux and windows
	flag.BoolVar(&BUILD_MAC,  "mac", BUILD_LINUX,      "  Builds a MAC OS binary")	
	flag.BoolVar(&BUILD_LINUX,  "linux", BUILD_LINUX,      "  Builds a LINUX binary")	
	flag.BoolVar(&BUILD_WIN,  "windows", BUILD_WIN,      "  Builds for Windows")	
	flag.BoolVar(&BUILD_WIN,  "win", BUILD_WIN,      "  alias to build for Windows")	

	flag.BoolVar(&BUILD_ARM,  "arm", BUILD_ARM,     "  forces a ARM (or Apple Silicone) binary compile ")	
	flag.BoolVar(&BUILD_INTEL,  "intel", BUILD_INTEL,   "  forces an INTEL/AMD Architecture binary compile")	
}