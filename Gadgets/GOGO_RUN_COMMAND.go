package CUSTOM_GO_MODULE

import (
	
	"strings"	
	"os/exec"
	//"math"
)

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
// Much better ver of RUN_COMMAND 
// Specify -verbose -rundir or -showoutput, -hideerror
func RUN_COMMAND(command_text string, ALL_PARAMS ...string) (string, []string) {

	command := strings.Split(command_text, " ")

	var hide_errors = true
	var showoutput = false
	var RUN_FROM_DIR = ""

	//2. Now, see if version was passed
	for x, VAL := range ALL_PARAMS {

		if VAL == "-showoutput" {
			showoutput = true
			continue
		}

		// This is the directory from which to run the program
		if VAL == "-rundir" {
			nnum := x + 1
			if nnum < len(ALL_PARAMS) {
				RUN_FROM_DIR = ALL_PARAMS[nnum]
			}
			continue
		}
		if VAL == "-hideerror" {
			hide_errors = false
			continue
		}

	} //end of for
	
	

	if len(command) < 2 {
		R.Println(" ERROR: Command specified is TOO SHORT")
	}

	cmd := exec.Command(command[0], command[1:]...)

	if RUN_FROM_DIR != "" {
		cmd.Dir = RUN_FROM_DIR
	}
	

	output, err := cmd.CombinedOutput()
	if err != nil {
		if hide_errors == false {
			R.Println(" Error in RUN COMMAND: ", err)
		}
	}

	// If we want to show the output
	if showoutput {
		W.Print("OUTPUT: ")
		C.Println(string(output))
	}

	// Also lets create a mult-line string so we can iterate through these lines if needed
	var str_multi []string
	for _, x := range output {
		str_multi = append(str_multi, string(x))
	}

	// Always return output so we can parse it
	
	return string(output), str_multi
}