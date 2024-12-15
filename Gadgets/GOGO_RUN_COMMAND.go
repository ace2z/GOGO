package CUSTOM_GO_MODULE

import (
	"strings"
	//"math"
	"github.com/go-cmd/cmd"
)

//  = =  UPDATED RUN COMMAND BEST EEVER  Dec 2024

/*
	NOTE here is teh way to use cmd.Exec WITH environment variables:


	cmd := exec.Command("go", "build", "-ldflags=-s -w -X main.VERSION_NUM=" + VERSION_to_USE, "-buildmode=exe", "-o", FULL_DEST_FILE )
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=" + BIN_TYPE)
	cmd.Env = append(cmd.Env, "GOARCH=" + ARCH)
	output, err := cmd.CombinedOutput()
	if err != nil {

		R.Println(" Error in RUN COMMAND: ", err)
	}
	C.Println(string(output))
*/

func SHOW_ME(line string, have_error bool) {
	if have_error {
		Y.Println(line)
	} else {
		W.Println(line)
	}
}

/*
ULTIMATE run shell commands function
Usage:

result, resultARR, example_command = SUPER_RUN_COMMAND("/opt/SCRIPTS/print_lines.sh")

FOR Realtime live, streaming feedback:

	SUPER_RUN_COMMAND("/opt/SCRIPTS/print_lines.sh", "stream")  // or realtime or live

HIDE all ouput (except if in stream mode)

	SUPER_RUN_COMMAND("/opt/SCRIPTS/print_lines.sh", "quiet")   // or silent or nooutput
*/
func SUPER_RUN_COMMAND(ALL_PARAMS ...interface{}) (string, []string, string) {

	var use_STREAM_MODE = false
	var comLIST = []string{}
	var just_COMMAND = ""
	cmdOptions := cmd.Options{
		Buffered:  true,
		Streaming: false,
	}

	var SHOW_OUTPUT = true // Output is shown.. unless the user specifies otherwise

	for n, param := range ALL_PARAMS {
		string_val, IS_STRING := param.(string)
		arr, IS_ARR := param.([]string)
		//int_val, IS_INT := param.(int)

		//First param is either a string with a cammand we run.. or an array with the command and all its arguments as elements
		if n == 0 {
			if IS_STRING {
				just_COMMAND = string_val

			} else if IS_ARR {
				if len(arr) > 0 {
					comLIST = arr
				}
			}
			continue
		}

		// Everythign else is a parameter
		if IS_STRING {
			if string_val == "silent" || string_val == "quiet" || string_val == "nooutput" {
				SHOW_OUTPUT = false
				continue
			}

			if string_val == "stream" || string_val == "realtime" || string_val == "live" {
				use_STREAM_MODE = true
				cmdOptions.Streaming = true
				cmdOptions.Buffered = false
				continue
			}

		}
	} //end of params FOr

	// Decide if we use the string or the LIST
	// either way we get the first command ... and save it off

	var FIRST_COMM = ""
	var ARG_LIST = []string{}

	// Generate an example command string. This is displayed or returned
	var example_command = just_COMMAND //strings.Join(comLIST, " ")
	if just_COMMAND != "" {

		msplit := strings.Split(just_COMMAND, " ")

		// The first element is the command we run
		//everything else is an argument
		if len(msplit) > 0 {
			FIRST_COMM = msplit[0]
			ARG_LIST = msplit[1:]
		}

	} else {

		FIRST_COMM = comLIST[0]
		ARG_LIST = comLIST[1:]
		example_command = strings.Join(comLIST, " ")
	}

	if SHOW_OUTPUT {
		Y.Println("")
		Y.Println("  ==|", "Now Running: ")
		C.Println(example_command)
		C.Println("")
	}

	// This actually RUNS the command. Returns a newComm CHANNEL object
	newComm := cmd.NewCmdOptions(cmdOptions, FIRST_COMM, ARG_LIST...)

	var res_STD = []string{}
	var res_ERR = []string{}

	// Default mode when running a command (not realtime)
	if use_STREAM_MODE == false {

		status := <-newComm.Start()
		for _, line := range status.Stderr {
			res_ERR = append(res_ERR, line)
			if SHOW_OUTPUT {
				SHOW_ME(line, true)
			}
		}
		for _, line := range status.Stdout {
			res_STD = append(res_STD, line)
			if SHOW_OUTPUT {
				SHOW_ME(line, false)
			}
		}

		// ELSE if we need to get realtime output from the Command and showing onscreen
	} else {

		//newComm := cmd.NewCmdOptions(cmdOptions, FIRST_COMM, ARG_LIST...)
		doneChan := make(chan struct{})
		go func() {
			defer close(doneChan)
			// Done when both channels have been closed
			// https://dave.cheney.net/2013/04/30/curious-channels
			for newComm.Stdout != nil || newComm.Stderr != nil {
				select {
				case line, open := <-newComm.Stderr:
					if !open {
						newComm.Stderr = nil
						continue
					}
					res_ERR = append(res_ERR, line)
					SHOW_ME(line, true) // Because we are in STREAMING mode, we must show the output on the screen

				case line, open := <-newComm.Stdout:
					if !open {
						newComm.Stdout = nil
						continue
					}
					res_STD = append(res_STD, line)
					SHOW_ME(line, false) // Because we are in STREAMING mode, we must show the output on the screen

				}

			} //end of FOR
		}()
		// Run and wait for Cmd to return, discard Status
		<-newComm.Start()

		// Wait for goroutine to print everything
		<-doneChan

	} //end of ELSE

	// Concat the two arrays into one
	final_ARR := append(res_STD, res_ERR...)

	// Join the array into a single string
	final_STR := strings.Join(final_ARR, "\n")

	// Return both

	return final_STR, final_ARR, example_command

}

// alias for SUPER_RUN_COMMAND
func RUN_COMMAND(ALL_PARAMS ...interface{}) (string, []string, string) {

	return SUPER_RUN_COMMAND(ALL_PARAMS...)
}
