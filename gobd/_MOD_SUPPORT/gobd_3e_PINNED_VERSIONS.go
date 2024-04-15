package _MOD_SUPPORT

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

var PINNED_VER_FILE = "_PINNED_VERSIONS.conf"

func SAVE_PINNED(inp string, PINNED *[]string) {

	var ALREADY_EXISTS = false
	for _, x := range *PINNED {
		if strings.Contains(x, inp) {
			ALREADY_EXISTS = true
			break
		}
	}

	if ALREADY_EXISTS == false {
		*PINNED = append(*PINNED, inp)
	}
}

func CORRELATE_PINNED(PINNED_LIST []string, GOMOD_CONTENTS []string) []string {

	for _, x := range PINNED_LIST {
		full_line := x
		// split on the one space that is there
		msplit := strings.Split(x, " ")
		if len(msplit) <= 1 {
			continue
		}

		packname := msplit[0]

		// now iterate GOMOD_CONTENTS
		for n, y := range GOMOD_CONTENTS {

			if strings.Contains(y, packname) {

				// msplit := strings.Split(y, " ")
				// if len(msplit) <= 1 {
				// 	continue
				// }
				// gm_pack := msplit[0]
				// gm_everything_else := msplit[1]

				// gm_pack = strings.Replace(y, gm_pack, packname, -1)
				GOMOD_CONTENTS[n] = "\t" + full_line + " // indirect"
			}
		}
	}

	return GOMOD_CONTENTS
}

func CHECK_for_PINNED_VERSIONS_Engine() {

	Y.Println(GLOBAL_PREFIX, "Checking for PINNED versions")

	if FILE_EXISTS(PINNED_VER_FILE) == false {
		Y.Print(GLOBAL_PREFIX, "No Pinned Version file: ")
		W.Println(PINNED_VER_FILE)
		return
	}

	// Otherwise.. open this file
	file, err := os.Open(PINNED_VER_FILE)
	if err != nil {
		M.Print(" **ERROR** Cannot open: ")
		W.Print(PINNED_VER_FILE)
		M.Println(" for some reason! ")
		Y.Println(err)
		return
	}

	//2. Now lets read the file itself
	var PINNED_LIST []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore commented lines
		if strings.Contains(line, "//") || strings.Contains(line, "#") {
			continue
		}
		if line == "" {
			continue
		}

		// Otherwise add ed pinned list
		SAVE_PINNED(line, &PINNED_LIST)

	} //end of for
	file.Close()

	if len(PINNED_LIST) == 0 {
		return
	}

	//3. Now we have all PINNED versions of items loaded.. lets load the go.mod file
	gmfile, gmerr := os.Open("go.mod")
	if gmerr != nil {
		M.Print(" **ERROR** Cannot open: ")
		W.Print("go.mod")
		M.Println(" for some reason! ")
		Y.Println(gmerr)
		return
	}

	scan2 := bufio.NewScanner(gmfile)

	var GOMOD_CONTENTS []string
	for scan2.Scan() {
		line := scan2.Text()

		//		line = strings.Replace(line, "\t", "", -1)

		GOMOD_CONTENTS = append(GOMOD_CONTENTS, line)

	} //end of for
	gmfile.Close()

	final_list := CORRELATE_PINNED(PINNED_LIST, GOMOD_CONTENTS)

	file_path := "go.mod"
	rerr := os.Remove(file_path)
	r2err := os.Remove("go.sum")
	if rerr != nil || r2err != nil {
		Y.Println("Cant Safety remove file: ", rerr, r2err)
		return
	}

	f, err := os.Create(file_path)
	if err != nil {
		Y.Println(" Cant create NEW file: ", err)
		return
	}
	defer f.Close()
	for _, value := range final_list {
		fmt.Fprintln(f, value) // print values to f, one per line
	}

	file.Close()

	//5. Run a final GO mod tidy
	G.Println(GLOBAL_PREFIX, "Running GOMOD Tidy on PINNED versions")
	RUN_COMMAND("go mod tidy")

}
