package _MOD_SUPPORT

import (
	"bufio"
	"strings"
	"os"
	"path"
	"path/filepath"

	. "local/_CORE"

	. "github.com/acedev0/GOGO/Gadgets"
	//. "github.com/acedev0/GOGO_Gadgets/FileOPS"
)


var PACKAGE_LIST []string


/*
	import . "github.com/acedev0/GOGO/Gadgets"
      or  
  import Gadgets "github.com/acedev0/GOGO/Gadgets"

*/
func FIND_all_GO_Files() []string {
	

	var GO_FILE_LIST []string

	err := filepath.Walk("./",
        func(full_file string, info os.FileInfo, err error) error {
            if err != nil {
                return err            
			}

			var just_FILENAME = info.Name()
			var just_EXT = path.Ext(just_FILENAME)

			if just_EXT == ".go" {
				GO_FILE_LIST = append(GO_FILE_LIST, full_file)
			}

            return nil
        })
    if err != nil {
        M.Println(" Error Searching for .go Files", err)        
    }

	C.Print(PREFIX, "Scanning for GO Files...")
	G.Print(len(GO_FILE_LIST))
	C.Println(" files were found..")
	return GO_FILE_LIST
}

// Make sure the line has " (quote) . (period) and / (slash) in it
func SPLIT_delims(r rune) bool {
    return r == '"'
}

func IS_GOMOD_Package(tmp_line string) (bool, string) {
	val_score := 0
	if strings.Count(tmp_line, "\"") == 2 {
		val_score++
	}
	if strings.Count(tmp_line, "/") >= 2 {
		val_score++
	}

	// Bug fix.. If we have back to back .. .. this isx not a go pacage
	if strings.Contains(tmp_line, "..") {
		return false, ""
	}

	// If this is a valid go package
	if val_score == 2 {
		sd := strings.FieldsFunc(tmp_line, SPLIT_delims)
		if len(sd) > 1 {
			packname := sd[1]
			// If there are ANY spaces in this package name.. its not a go package
			if strings.Count(packname, " ") > 0 {
				return false, ""
			}			
			if strings.Count(packname, ".") >= 1 {
				return true, packname
			}			
		}
	}


	// default to false
	return false, ""
}

func SAVE_PACKAGE(packname string) {
	
	
	ALREADY_EXISTS := false
	for _, x := range PACKAGE_LIST {
		if len(x) < 2 {
			continue
		}
		if x == packname {
			ALREADY_EXISTS = true
			break
		}
	}

	if ALREADY_EXISTS == false {
		PACKAGE_LIST = append(PACKAGE_LIST, packname)
	}
}

func EXTRACT_Import_Packages(FILE_LIST []string) {

	for _, full_file := range FILE_LIST {
		file, err := os.Open(full_file)
		if err != nil {
			R.Println(" **ERROR** Cannot open csv_file: ", full_file, err)
			return
		}

		scanner := bufio.NewScanner(file)

		
		SEARCH_via_MULTI := false
		BLOCKED_UNTIL := false
		for scanner.Scan() {

			// First look for import line. lets be careful with this
			// We want to trim any spaces leading or following.. then ubstr for 'import '
			line := scanner.Text()    

			// Trigger on on Import statement (either single.. or multi-line)
			if strings.Contains(line, "import") && BLOCKED_UNTIL == false {
				line = strings.TrimSpace(line)

				// See if this is a multiline statement
				if strings.Contains(line, "(") {
					SEARCH_via_MULTI = true
					continue
				
				// Else this is a single line import statement
				} else {
					found, pack_name := IS_GOMOD_Package(line)
					if found {
						SAVE_PACKAGE(pack_name)
					}
				}				
			}

			if BLOCKED_UNTIL {
				if strings.Contains(line, "*/") {
					BLOCKED_UNTIL = false
				}
				continue
			}
			
			//3. If this is a line that has comments. We skip it
			if strings.Contains(line, "//") {
				line = strings.TrimSpace(line)

				// THe comment may also be at the end of the line. If this is the case.. we may have a package to import
				if strings.Contains(line, " //") {
					// Get everything BEFORE the comment.. check to see if this is a go package
					ksplit := strings.Split(line, " //")
					if len (ksplit) >= 1 {
						newline := ksplit[0]
						found, pack_name := IS_GOMOD_Package(newline)
						if found {
							SAVE_PACKAGE(pack_name)
						}
					}												
				}

				continue
			}

			if strings.Contains(line, "/*") {
				line = strings.TrimSpace(line)
				if strings.Contains(line, " /*") {
					// Get everything BEFORE the comment.. check to see if this is a go package
					ksplit := strings.Split(line, " /*")
					if len (ksplit) >= 1 {
						newline := ksplit[0]
						found, pack_name := IS_GOMOD_Package(newline)
						if found {
							SAVE_PACKAGE(pack_name)
						}
					}
					
					// This blocks until we have a corresponding */
					BLOCKED_UNTIL = true
					continue
				}
			}


			//3. If we get this far and Search Multi is true,, check for go package
			if SEARCH_via_MULTI {
				found, pack_name := IS_GOMOD_Package(line)
				if found {
					SAVE_PACKAGE(pack_name)
				}				
			}
		}
	}

	C.Print(PREFIX, "Total gomod dependencies to import: ")
	G.Println(len(PACKAGE_LIST))
	SHOW_STRUCT(PACKAGE_LIST)
	
}

func CACHE_all_IMPORTS() {	

	// Find all the go files recurseivly
	FILE_LIST := FIND_all_GO_Files()

	// Now extract all the import statements from all the go Files
	EXTRACT_Import_Packages(FILE_LIST)


}
