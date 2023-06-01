package _BUILD_ENGINE

import (
	"os"
	"strings"
	"path/filepath"

	"os/exec"

	. "local/_CORE"
	. "local/_MOD_SUPPORT"

	. "github.com/acedev0/GOGO/Gadgets"
)


func GO_BUILD_Engine() {
	SHOW_BOX("Buidling GO Program")

	GOMOD_Dependency_Engine()

	FULL_COMMAND :="GOOS=" + BIN_TYPE + " GOARCH=" + ARCH + " go build -ldflags=\"-s -w -X main.VERSION_NUM=" + VERSION_to_USE + "\" -buildmode=exe -o " + FULL_DEST_FILE
	C.Println(PREFIX, "GO build using: ")
	Y.Println(PREFIX, FULL_COMMAND)
	
	cmd := exec.Command("go", "build", "-ldflags=-s -w -X main.VERSION_NUM=" + VERSION_to_USE, "-buildmode=exe", "-o", FULL_DEST_FILE )
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=" + BIN_TYPE)
	cmd.Env = append(cmd.Env, "GOARCH=" + ARCH)
	output, err := cmd.CombinedOutput()
	if err != nil {
		
		R.Println(" Error in RUN COMMAND: ", err)
	}

	result_OUT := string(output)
	C.Println(result_OUT)

	
	// Check for errors.. particularly: no required module provides package
	if strings.Contains(result_OUT, "no required module provides package") {
		C.Println("")
		Y.Println(" = = =| WARNING (maybe error) | = = =")

		W.Println(" Looks like there is a go.mod Module dependency error / issue")
		W.Print(" This might mean there are too many")
		G.Print(" go.mod ")
		W.Print("files where they ")
		M.Println("SHOULD NOT BE")
		W.Println(" This can ALSO mean: ")
		W.Println("  - One of your import commands has a typo in the package name")
		W.Println("  - You accidentally ran a 'go build' or 'go mod init' in the wrong directory")
		W.Println("  - One of your custom MODULES is using the wrong / old version of ANOTHER CUSTOM Module")
		W.Println("    This is a little harder to detect and you may want to just re-run go mod tidy on all")
		W.Println("    of your CUSTOM modules ")
		W.Print("  - Something is completely fubar'd in GO Cache. So just run")
		Y.Println(" godb --purge")
		W.Println("")
		W.Println(" Check the error message carefully! Also for your convenience,")
		W.Print(" Here is a list of all the ")
		G.Print("go.mod ") 
		W.Println("files from the previous directory down..")
		W.Println("")
		err := filepath.Walk("../", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				M.Println("Error in File Walk: ", err)
				return err
			}

			if strings.Contains(path, "go.mod") {
				Y.Println(path)
			}
			return nil
		})
		if err != nil {
			M.Println("Error in File Walk: ", err)
		}
	}

	
}
