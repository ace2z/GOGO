package main

import (
	. "local/CORE"
	. "local/RUNMODE_ENGINE"

	// . "local/_BUILD_ENGINE"
	// . "local/_MOD_SUPPORT"

	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/EMBED_Ops"
)

var VERSION_NUM = ""

func main() {
	CLI_PARAMS_INIT()

	MASTER_INIT(" devopsCLI - Azure and AWS CLI helper", VERSION_NUM, "-showarch")
	W.Println("")

	CHECK_PreReqs()

	// Now.. RunMode Engine  - will run based on the parameters that were specified or found

	RUN_MODE_Engine_INIT()
	W.Println("")
}
