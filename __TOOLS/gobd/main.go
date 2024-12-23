package main

import (
	. "local/_BUILD_ENGINE"
	. "local/_CORE"
	. "local/_MOD_SUPPORT"

	//. "local/_DATABASE"

	//. "local/_STOCKS"
	//. "local/_ECONOMICS"

	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/EMBED_Ops"
)

var VERSION_NUM = ""

func main() {
	CLI_PARAMS_INIT()

	MASTER_INIT(" gobd - Golang Build Tools ( simplified! )", VERSION_NUM, "-showarch")
	W.Println("")

	PROCESS_OPTIONS()

	if DO_CLEAN || JUST_CLEAN {
		CLEAN_CACHE("clean")
		if JUST_CLEAN {
			DOEXIT()
		}
	} else if DO_PURGE || JUST_PURGE {
		CLEAN_CACHE("purge")
		if JUST_PURGE {
			DOEXIT()
		}
	}

	//2. Check for required Prerequisites
	CHECK_PreReqs()

	//3. Determine if we are buildinga  regular GO program.. or a GO Module
	if INIT_MOD || TEST_MOD || JUST_TEST {
		// If this is just a test.. we don't need to do anything else
		if JUST_TEST {
			GOMOD_Dependency_Engine()
			RUN_GO_Test()
			DOEXIT()
		}
		BUILD_BASIC_GO_PROGRAM = false
		GET_REPO_MetaDATA()
		GOMOD_Core_Ops()
	}

	//4. If this is a regular program
	if BUILD_BASIC_GO_PROGRAM {
		GO_BUILD_Engine()
	}
	//3. Get REPO data

}
