package CORE

import (
	"strings"
	//	"strconv"
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)

// Useful Globals
var PREFIX = "   -| "
var ERR_CODE = 1

func SET_FLAG_Based_on_SUPPORTED(cloudval string) {
	if cloudval == "aws" {
		MODE_AWS = true
		return
	}

	if cloudval == "azure" {
		MODE_AZURE = true
		return
	}
}

func CHECK_PreReqs() {
	if VERBOSE_MODE {
		C.Println(PREFIX, "Checking Prereqs...  ")
	}
	fin_res := ""
	if MODE_AWS {
		result, _, _ := RUN_COMMAND("aws --version", "silent")
		if strings.Contains(result, "aws-cli") == false {
			W.Println("")
			M.Print(PREFIX, " *** ERROR ***")
			Y.Println(" AWS cli is NOT installed.. or is NOT in your path!")
			os.Exit(ERR_CODE)
		}
		fin_res = result

	} else if MODE_AZURE {
		result, _, _ := RUN_COMMAND("az --version", "silent")
		if strings.Contains(result, "azure-cli") == false {
			W.Println("")
			M.Print(PREFIX, " *** ERROR ***")
			Y.Println(" AZURE az-cli is NOT installed.. or is NOT in your path!")
			os.Exit(ERR_CODE)
		}
		fin_res = result
		// // Pull out JUST the az cli version
		// msplit := strings.Split(result, "\n")[0]
		// if len(msplit) > 0 {
		// 	Y.Println(msplit[0])
		// }
		// Else .. lets see if DEVOPSCLI_DEFAULT_CLOUD is set
	} else {

		envval, exists := os.LookupEnv("DEVOPSCLI_DEFAULT_CLOUD")
		if exists == false {
			M.Println(" ERROR: No Supported Cloud platform was specified!")
			Y.Println(" Please rerun with --aws or --azure (or another supported cloud platform)")
			Y.Print(" Or set the ")
			W.Print(" export DEVOPSCLI_DEFAULT_CLOUD=aws")
			Y.Println(" environment variable")
			Y.Println(" ...also, you can rerun with --help for more information")
			Y.Println("")
			DO_EXIT()
		}

		// Otherwise set whichever flag based on what we support
		SET_FLAG_Based_on_SUPPORTED(envval)

	}

	//3. If we get this far.. Lets show validation (or error out)

	if fin_res != "" {
		if VERBOSE_MODE {
			Y.Println(fin_res)
		}
		if MODE_AWS {
			G.Println(CPREFIX, "AWS CLI Found")
		} else if MODE_AZURE {
			G.Println(CPREFIX, "AZURE CLI Found")
		}
		W.Println("")
	} else {
		M.Println(" ERROR: No valid CLI detected")
		os.Exit(ERR_CODE)
	}

	//4. Now.. we need to run the check for Params..ENVIRONMENT variables (or whatever was specified on the CLI)
	REQ_Params_Check_for_SUPPORTED()

}
