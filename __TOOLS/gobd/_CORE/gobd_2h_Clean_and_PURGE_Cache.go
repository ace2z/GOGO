package _CORE

import (
	"os"
	// "strings"
	// "io/ioutil"
	//. "local/_CORE"

	. "github.com/ace2z/GOGO/Gadgets"
)

func CLEAN_CACHE(ctype string) {

	SHOW_BOX(" CLEANING Go Caches")

	RUN_COMMAND("go clean -modcache -i -r", "silent")
	os.Remove("go.mod")
	os.Remove("go.sum")

	// If we are doing a basic 'clean' just clear the mod cache
	if ctype == "clean" {

	} else if ctype == "purge" {
		Y.Println("")

		VERIFICATION_PROMPT("Hey you are about to DEEP PURGE your whole GO CACHE.\n       This is a pretty destructive operation and will\n       reset all your vscode GO tools \n       You may want to do a --clean instead", "YES", "-exit_on_fail")

		RUN_COMMAND("go clean -cache -i -r")

		// Also lets delete the actual gocache and go mod cach directories
		// They will get re-created
		cachedir, _, _ := RUN_COMMAND("go env GOCACHE", "silent")
		modcache, _, _ := RUN_COMMAND("go env GOMODCACHE", "silent")
		gopath, _, _ := RUN_COMMAND("go env GOPATH", "silent")

		err := os.RemoveAll(cachedir)
		if err != nil {
			C.Println("Cachedir ERR", err)
		}
		err = os.RemoveAll(modcache)
		if err != nil {
			C.Println("modcache ERR", err)
		}

		err = os.RemoveAll(gopath)
		if err != nil {
			C.Println("gopath ERR", err)
		}

		SHOW_BOX("FULL Go Cache PURGE ", "|magenta|COMPLETED ..")

	}
	// var FILE_LIST []string

	// cwd, _ := os.Getwd()
	// files, err := ioutil.ReadDir(cwd)
	// if err != nil {
	// 	M.Println(" Cant read local dir: ", err)
	// }
	// for _, file := range files {
	// 	if file.IsDir() {
	// 		continue
	// 	}
	// 	if strings.Contains(file.Name(), ".go") {
	// 		FILE_LIST = append(FILE_LIST, file.Name())
	// 	}
	// }
}
