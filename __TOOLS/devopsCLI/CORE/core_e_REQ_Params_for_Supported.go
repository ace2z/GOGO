package CORE

import (
	"strings"
	//	"strconv"
	"os"

	. "github.com/ace2z/GOGO/Gadgets"
)

func envLookupTool(lookfor string) string {
	envval, exists := os.LookupEnv(lookfor)

	if exists == false {
		W.Println("")
		M.Println(CPREFIX, " *** ERROR: PARAMETER ERROR")
		Y.Print(CPREFIX, "I cannot find an ENV variable for: ")
		W.Println(lookfor)
		Y.Println(CPREFIX, "..and NONE was specified on the command line.")
		Y.Println(CPREFIX, "re-run with --help for more information.")
		DO_EXIT()

	}

	resval := strings.TrimSpace(envval)

	return resval

}

func AZURE_Populate_ENV_Vars() {
	// Only runs if the supported mode is active
	if MODE_AZURE == false {
		return
	}

	// Error handling.. Make sure we have either ENVIRONMENT variables.. Or they are passed via command line
	if AZ_TENANT == "" {
		AZ_TENANT = envLookupTool("AZ_TENANT_ID")
	}
	if AZ_SUB == "" {
		AZ_SUB = envLookupTool("AZ_SUBSCRIPTION_ID")
	}
	if AZ_LOCATION == "" {
		AZ_LOCATION = envLookupTool("AZURE_DEFAULTS_LOCATION")
	}

}

func REQ_Params_Check_for_SUPPORTED() {
	AZURE_Populate_ENV_Vars() // For azure (if AZUREMODE Is activated)
	// PLACEHOLDER FOR AWS
	// PLACEHOLDER FOR GCP ...if i EVER start using this

}
