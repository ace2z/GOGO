package AZURE_SUPPORT

import (

	//"io/ioutil"
	. "local/CORE"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

func Azure_TERRA_Service_Principal() {
	if AZ_SP_NAME == "" {
		return
	}

	if AZ_SP_SCOPES == "" {
		M.Println("ERROR: No Scope was specified for the Service Principal")
		Y.Print("Scopes can be a SUBSCRIPTION id or (RESOURCE id within a subscription: ")
		W.Println("/subscriptions/" + AZ_SUB)

		Y.Print("Rerun with ")
		W.Println("--scope ##_SCOPE_ID_###")
		DO_EXIT()
	}

	Y.Println("")
	Y.Println(CPREFIX, "Will generate Azure Service Principal (For Terraform to Use)")
	Y.Print(CPREFIX, " Service Principal Name: ")
	W.Println(AZ_SP_NAME)
	Y.Println(CPREFIX, "On the following Scopes Id: ")
	msplit := strings.Split(AZ_SP_SCOPES, " ")
	if len(msplit) > 0 {
		for _, val := range msplit {
			W.Println("  ", val)
		}
	}

	if DEBUG_MODE {
		PressAny()
	}

	// WORKS
	//  az login --tenant 2de292bc-3103-4b0d-b69e-445f8bc4b6b1 --service-principal --username d4f1f5a1-6a70-4cf5-955d-a29e3ea0c71a --certificate /Users/cowboy/tmpvhd17oy3.pem
	var startLIST = []string{
		"az", // REmove this if you cant get the SuperRunCommand Working
		"ad",
		"sp",
		"create-for-rbac",
	}
	if AZ_SP_GEN_USING_CERT {
		startLIST = append(startLIST, "--create-cert")

		// // Make sure they specified a KEYVAULT to use for storing the cert
		// if AZ_SPCERT_KEYVAULT_NAME != "" {
		// 	ltmp := "--keyvault=" + AZ_SPCERT_KEYVAULT_NAME

		// 	startLIST = append(startLIST, ltmp)

		// 	// ltmp2 := "--cert"
		// 	// startLIST = append(startLIST, ltmp2)
		// }
	}

	var comLIST = []string{
		"--name=" + AZ_SP_NAME,
		"--role=" + AZ_SP_ROLE,
		"--years=" + AZ_SP_EXP_YEARS,
		"--scopes=" + AZ_SP_SCOPES,
	}

	// Now run the command
	full_COMM := append(startLIST, comLIST...)

	// Y.Println(CPREFIX, " Running: ")
	// C.Println("  az", example_command)

	result, _, _ := SUPER_RUN_COMMAND(full_COMM)

	success := false
	if strings.Contains(result, "\"appId\":") {
		success = true
	}

	if success {
		lookup := []string{"appId", "password", "fileWithCertAndPrivateKey"}
		resmap := Extract_JSON_from_CLI_STRING(result, lookup)

		ARM_CLIENT_ID := resmap["appId"].(string)
		ARM_CLIENT_SECRET := resmap["password"].(string)
		tmpCERTPATH := resmap["fileWithCertAndPrivateKey"].(string)

		if AZ_SP_GEN_USING_CERT {
			ARM_CLIENT_SECRET = "Contents of FILE located at: " + tmpCERTPATH
		}

		ARM_TENANT_ID := AZ_TENANT
		ARM_SUBSCRIPTION_ID := AZ_SUB

		G.Println(CPREFIX, "Service Principal Creation SUCCESS!!")
		C.Print(CPREFIX, " NOTE: If you need to use MULTIPLE scopes, rerun the ABOVE ")
		W.Print("az ad sp")
		C.Println(" command")
		C.Println(CPREFIX, "..and include each additional scope delimited by a SPACE")

		G.Println("")
		G.Println("Terraform Helper: ")
		Y.Println("Needs to be exposed in your agent/runners ENVIRONMENT. or your local shell profile")
		W.Println("")
		W.Print("export ARM_CLIENT_ID=")
		C.Println(ARM_CLIENT_ID)

		W.Print("export ARM_CLIENT_SECRET=")
		C.Println(ARM_CLIENT_SECRET)

		W.Print("export ARM_TENANT_ID=")
		C.Println(ARM_TENANT_ID)

		W.Print("export ARM_SUBSCRIPTION_ID=")
		C.Println(ARM_SUBSCRIPTION_ID)

		W.Println("\n(Add these lines to your shell profile)")

		W.Println("")

	}

	W.Println("")

} //end of login
