package AZURE_SUPPORT

import (

	//"io/ioutil"
	. "local/CORE"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

// / Params for Azure Support.. based on whats sent in CLI
var AZ_NEED_SP = false

func Check_EXISTING_Azure_Session() {
	// Check if we have an existing Azure session
	result, _, _ := RUN_COMMAND("az account show", "silent")
	if strings.Contains(result, "run 'az login'") {
		M.Println("ERROR: No EXISTING Azure Login Sessions Found!")
		Y.Println("Rerun with WITHOUT the --existing parameter")
		DO_EXIT()
	}
}

func AZURE_Support_Engine() {
	if MODE_AZURE == false {
		return
	}

	// For logging out of Azure
	if AZ_LOGOUT {
		Y.Print(CPREFIX, "Completely Logged OUT of ")
		M.Println("Azure CLI Session")
		RUN_COMMAND("az logout", "silent")
		RUN_COMMAND("az account show", "justfeedback")
		DO_EXIT()
	}

	// Overrides for Azure GovCloud (not complete. Only supports US)
	if AZ_USE_GOVCLOUD {
		RUN_COMMAND("az cloud set --name AzureUSGovernment", "silent")
		C.Print(CPREFIX, " Azure ")
		G.Print("AzureGovCloud MODE ")
		C.Println("is CONFIGURED!!")

	} else {
		C.Println(CPREFIX, "Azure Cloud is Configured")
		RUN_COMMAND("az cloud set --name AzureCloud", "silent")
	}

	// Always start with Login (if IF JUST_LOGIN --login was specified set.. we exist after)
	if USE_EXISTING_LOGIN_SESSION {
		Check_EXISTING_Azure_Session()
	} else {
		AZURE_Login()
	}
	G.Print(CPREFIX, " AZURE az-cli Login Success! ")
	G.Println("")

	//2. All the additoinal functionality functions listed here.. 
	// Everything is based on a FLAG being set.. So if it is not set, they will exit quietly
	Azure_TERRA_Service_Principal()

}
