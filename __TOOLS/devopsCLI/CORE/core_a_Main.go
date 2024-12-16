package CORE

import (
	"flag"
	"strings"
)

var CPREFIX = "  ==|"
var MODE_AWS = false
var MODE_AZURE = false
var VERBOSE_MODE = false
var USE_EXISTING_LOGIN_SESSION = false

type StringList []string

// This is for getting multiple , Comma seperatred, string parameters (as strings)
func (m *StringList) String() string {
	return strings.Join(*m, ",")
}
func (m *StringList) Set(value string) error {
	*m = strings.Split(value, ",")
	return nil
}

var JUST_LOGIN = false

// For Azure params
var AZ_TENANT = ""
var AZ_SUB = ""
var AZ_LOCATION = ""
var AZ_USE_GOVCLOUD = false

var AZ_SP_SCOPES = "" //StringList
var AZ_SP_NAME = ""
var AZ_SP_GEN_USING_CERT = false
var AZ_SP_ROLE = "Contributor"
var AZ_SP_EXP_YEARS = "1" //  1 == 1 year... 5 == 5 years
// var AZ_SPCERT_KEYVAULT_NAME = ""
var AZ_LOGOUT = false

// For AWS params
var AWS_ACCOUNT = ""

func CLI_PARAMS_INIT() {

	// Basic admin Params
	flag.BoolVar(&VERBOSE_MODE, "verbose", VERBOSE_MODE, "  Run in VERBOSE mode")

	// Generics that can apply to any cloud platform
	flag.BoolVar(&JUST_LOGIN, "login", JUST_LOGIN, "  JUST Does the login process for whichever cloud env you selected")
	flag.BoolVar(&USE_EXISTING_LOGIN_SESSION, "existing", USE_EXISTING_LOGIN_SESSION, "  Uses the existing Login Session (instead of full login and verification process) Will error if none is available")
	flag.BoolVar(&USE_EXISTING_LOGIN_SESSION, "exist", USE_EXISTING_LOGIN_SESSION, "  (alias) --existing")
	flag.BoolVar(&USE_EXISTING_LOGIN_SESSION, "useexist", USE_EXISTING_LOGIN_SESSION, "  (alias) --existing")
	flag.BoolVar(&USE_EXISTING_LOGIN_SESSION, "persist", USE_EXISTING_LOGIN_SESSION, "  (alias) --existing")

	// = = = AZURE Specific Params
	flag.BoolVar(&MODE_AZURE, "azure", MODE_AZURE, "  Run in AZURE mode (requires az cli installed) or set DEVOPSCLI_DEFAULT_CLOUD=azure")
	flag.BoolVar(&AZ_USE_GOVCLOUD, "usegov", AZ_USE_GOVCLOUD, "  AZURE specifiy if you want to use Azure GovCloud tenants")
	flag.StringVar(&AZ_TENANT, "tenant", AZ_TENANT, "  Specify the AZURE Tenant ID")
	flag.StringVar(&AZ_SUB, "sub", AZ_SUB, "  Specify the AZURE SUBSCRIPTION ID")
	flag.StringVar(&AZ_SUB, "subscription", AZ_SUB, "  (alias) for --sub")
	flag.StringVar(&AZ_LOCATION, "location", AZ_LOCATION, "  Specify Azure Location (ie eastus) if AZURE_DEFAULTS_LOCATION is not set")

	// FOr the service Princiapl Stuff
	flag.StringVar(&AZ_SP_NAME, "sp", AZ_SP_NAME, "  AZURE SP (Service Principal) NAME to use")
	flag.StringVar(&AZ_SP_SCOPES, "scope", AZ_SP_SCOPES, " Azure SCOPE For the SP (a subscription ID or resource ID)")
	flag.BoolVar(&AZ_SP_GEN_USING_CERT, "spcert", AZ_SP_GEN_USING_CERT, "  Generate new Service Principal but creates SelfSigned PEM CERT file instead of password")
	flag.StringVar(&AZ_SP_ROLE, "sprole", AZ_SP_ROLE, "  AZURE SP ROLE That should be assigned to the new SP you generate (ie: Contributor or Owner)")
	flag.StringVar(&AZ_SP_EXP_YEARS, "spexp", AZ_SP_EXP_YEARS, "  AZURE SP Expiration (in years)")
	//flag.StringVar(&AZ_SPCERT_KEYVAULT_NAME, "spkv", AZ_SPCERT_KEYVAULT_NAME, "  AZURE SP when using a CERT. Keyvault NAME of where to store the PEM file")

	flag.BoolVar(&AZ_LOGOUT, "azlogout", AZ_LOGOUT, "  Force logs out of your current AzureCLI Session")

	// = = = AWS Params
	flag.BoolVar(&MODE_AWS, "aws", MODE_AWS, "  Run in AWS mode (requies AWS cli installed) or set DEVOPSCLI_DEFAULT_CLOUD=aws")

}
