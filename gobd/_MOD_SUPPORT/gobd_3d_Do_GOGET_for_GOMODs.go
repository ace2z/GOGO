package _MOD_SUPPORT

import (	

	. "local/_CORE"

	. "github.com/ace2z/GOGO/Gadgets"
)

/*
	import . "github.com/ace2z/GOGO/Gadgets"
      or  
  import Gadgets "github.com/ace2z/GOGO/Gadgets"

*/
func RUN_GOGET_for_IMPORTS() {


	//Then init a new gomod file .. if this is a "program" it is called local.. so we can easily use local packages.
	// if this is meant to be a go module.. We determine that based on the repo and current directory name

	DO_GOMOD_Init()

	for _, pack := range PACKAGE_LIST {		
		//COMM := GO_BIN + " get -u " + pack 
		COMM := "go get -u " + pack 
		C.Print(PREFIX, "Now running: ")
		Y.Println(COMM)

		result, _ := RUN_COMMAND(COMM)

		if VERBOSE_MODE {
			Y.Println(result)
		}
	}
}
