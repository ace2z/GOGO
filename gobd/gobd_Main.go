package main

import (
    . "local/_CORE"
    . "local/_BUILD_ENGINE"
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

    CHECK_PreReqs()
    if JUST_TEST { 
        GOMOD_Dependency_Engine()
        RUN_GO_Test()
        DOEXIT()
    }
   
    

    
    if TEST_MOD || MAKE_MOD { 

        GOMOD_Core_Ops() 
    
    } else {

        GO_BUILD_Engine()

    }
}
