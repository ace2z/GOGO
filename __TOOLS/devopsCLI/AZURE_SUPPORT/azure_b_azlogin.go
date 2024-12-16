package AZURE_SUPPORT

import (
	"os"
	"strings"

	//"io/ioutil"

	. "local/CORE"

	. "github.com/ace2z/GOGO/Gadgets"
)

func AZURE_Login() {

	Y.Println(CPREFIX, "AZURE Login Attempt")
	Y.Print(CPREFIX, " Will login to ")
	W.Print("TENANT: ")
	C.Println(AZ_TENANT)
	Y.Print(CPREFIX, " using ")
	W.Print("  SUBSCRIPTION: ")
	C.Println(AZ_SUB)
	if DEBUG_MODE {
		PressAny()
	}

	// cmdOptions := cmd.Options{
	// 	Buffered:  false,
	// 	Streaming: true,
	// }

	// //newComm := cmd.NewCmdOptions(cmdOptions, "./testAZ.sh")
	// newComm := cmd.NewCmdOptions(cmdOptions, "az", "login", "--tenant", AZ_TENANT, "--use-device-code")
	// doneChan := make(chan struct{})
	// go func() {
	// 	defer close(doneChan)
	// 	// Done when both channels have been closed
	// 	// https://dave.cheney.net/2013/04/30/curious-channels
	// 	for newComm.Stdout != nil || newComm.Stderr != nil {
	// 		select {
	// 		case line, open := <-newComm.Stdout:
	// 			if !open {
	// 				newComm.Stdout = nil
	// 				continue
	// 			}
	// 			C.Println(line)

	// 		case line, open := <-newComm.Stderr:
	// 			if !open {
	// 				newComm.Stderr = nil
	// 				continue
	// 			}
	// 			//errstring := newComm.Stderr.String()
	// 			//M.Println(newComm.Stderr)
	// 			Y.Println(CPREFIX, line)
	// 			if strings.Contains(line, "ERROR:") {
	// 				have_error = true
	// 			}
	// 		}
	// 	}
	// }()

	// // Run and wait for Cmd to return, discard Status
	// <-newComm.Start()

	// // Wait for goroutine to print everything
	// <-doneChan

	runcomm := "az login --tenant " + AZ_TENANT + " --use-device-code"
	result, _, _ := SUPER_RUN_COMMAND(runcomm, "stream")
	lower := strings.ToLower(result)

	if strings.Contains(lower, "error:") {

		W.Println("")
		M.Println(" UNKNOWN ERROR:")
		Y.Println(" Are you trying to use GovCloud but this isnt configured on your Tenant?")
		Y.Println(" Did you use the correct TENANT and SUBSCRIPTION ID? ")
		W.Println("")
		DO_EXIT()
	}

	// End of program
	// IF JUST_LOGIN was set.. we exit the program here
	if JUST_LOGIN {
		C.Print("JUST_LOGIN --login was specified, ")
		G.Println("Exiting...")
		W.Println("")
		DO_EXIT()
		os.Exit(0)
	} else {
		W.Println("")
	}

} //end of login
