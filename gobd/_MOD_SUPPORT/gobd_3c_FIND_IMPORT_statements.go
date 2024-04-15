package _MOD_SUPPORT

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"strings"

	. "local/_CORE"

	. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO_Gadgets/FileOPS"
)

/*
		import . "github.com/ace2z/GOGO/Gadgets"
	      or
	  import Gadgets "github.com/ace2z/GOGO/Gadgets"
*/
var max_depth = 1 // Maximum depth recursiveness we use
func FIND_all_GO_Files() []string {

	var GO_FILE_LIST []string

	err := filepath.Walk("./",
		func(full_file string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Prevent us from going too deep with recursevness when we scan for go files
			var too_deep = false
			if strings.Count(full_file, string(os.PathSeparator)) > max_depth {
				too_deep = true
			}

			var just_FILENAME = info.Name()
			var just_EXT = path.Ext(just_FILENAME)

			if just_EXT == ".go" && too_deep == false {
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

func is_EXTERNAL_GO_IMPORT(tmp_name string) bool {

	// make sure there is a . and a / in the tmp_name

	if strings.Contains(tmp_name, ".") && strings.Contains(tmp_name, "/") {
		return true
	}

	return false

	// val_score := 0
	// if strings.Count(tmp_line, "\"") == 2 {
	// 	val_score++
	// }
	// if strings.Count(tmp_line, "/") >= 2 {
	// 	val_score++
	// }

	// // Bug fix.. If we have back to back .. .. this isx not a go pacage
	// if strings.Contains(tmp_line, "..") {
	// 	return false, ""
	// }

	// // If this is a valid go package
	// if val_score == 2 {
	// 	sd := strings.FieldsFunc(tmp_line, SPLIT_delims)
	// 	if len(sd) > 1 {
	// 		packname := sd[1]
	// 		// If there are ANY spaces in this package name.. its not a go package
	// 		if strings.Count(packname, " ") > 0 {
	// 			return false, ""
	// 		}
	// 		if strings.Count(packname, ".") >= 1 {
	// 			return true, packname
	// 		}
	// 	}
	// }

	// // default to false
	// return false, ""
}

func SAVE_IMPORT(packname string, TMP_LIST *[]string) {

	ALREADY_EXISTS := false
	for _, x := range *TMP_LIST {
		if len(x) < 2 {
			continue
		}
		if x == packname {
			ALREADY_EXISTS = true
			break
		}
	}

	if ALREADY_EXISTS == false {
		*TMP_LIST = append(*TMP_LIST, packname)
	}
}

func GET_everything_BEFORE_COMMENT(line string) string {

	var delim = ""
	if strings.Contains(line, "//") {
		delim = "//"
	} else if strings.Contains(line, "#") {
		delim = "#"
	}

	// Get everything BEFORE the comment.. check to see if this is a go package
	ksplit := strings.Split(line, delim)
	if len(ksplit) >= 1 {
		newline := ksplit[0]
		newline = strings.TrimSpace(newline)

		return newline
	}

	return ""

}
func get_everything_IN_QUOTES(line string) string {
	line = strings.TrimSpace(line)
	var result = ""

	if line == "" {
		return ""
	}
	var tmp_input = line
	//1. First check to see if there is a single line comment on this line
	if strings.Contains(line, "//") {
		tmp_input = GET_everything_BEFORE_COMMENT((line))
		if tmp_input == "" {
			return ""
		}
	}

	// Next, lets get evertthing in quotes
	ksplit := strings.Split(tmp_input, "\"")
	if len(ksplit) > 0 {
		newline := ksplit[1]
		newline = strings.TrimSpace(newline)
		result = newline
	}
	return result
}

func EXTRACT_Import_STATEMENTS(FILE_LIST []string) []string {

	// SHOW_STRUCT(FILE_LIST, "white")
	// Y.Println("File List ABOVE")
	// PressAny()
	var IMPORT_LIST []string

	for _, full_file := range FILE_LIST {
		file, err := os.Open(full_file)
		if err != nil {
			R.Println(" **ERROR** Cannot open csv_file: ", full_file, err)
			return []string{}
		}

		scanner := bufio.NewScanner(file)

		SEARCH_via_MULTI := false
		BLOCKED_FOR_MULTI_COMMENT := false
		for scanner.Scan() {

			// First look for import line. lets be careful with this
			// We want to trim any spaces leading or following.. then ubstr for 'import '
			line := scanner.Text()

			//Safety.. if this is a multi line comment.. Block till we find the closing end
			if strings.Contains(line, "/*") {
				BLOCKED_FOR_MULTI_COMMENT = true
				continue
			}

			// Trigger on on Import statement (either single.. or multi-line)
			if strings.Contains(line, "import") && BLOCKED_FOR_MULTI_COMMENT == false {
				line = strings.TrimSpace(line)

				// See if this is a multiline statement
				if strings.Contains(line, "(") {
					SEARCH_via_MULTI = true
					continue

					// Else this is a single line import statement
				} else {

					// Make sure this is an external GO package improt
					packname := get_everything_IN_QUOTES(line)
					if is_EXTERNAL_GO_IMPORT(packname) {
						SAVE_IMPORT(packname, &IMPORT_LIST)
					}

					continue

				}
			}

			if BLOCKED_FOR_MULTI_COMMENT {
				if strings.Contains(line, "*/") {
					BLOCKED_FOR_MULTI_COMMENT = false
				}
				continue
			}

			//3. If this is a multi_line import
			if SEARCH_via_MULTI {

				//3b Check to see if this is a closing )
				if strings.Contains(line, ")") {
					SEARCH_via_MULTI = false
					continue
				}

				//3c otherwise
				packname := get_everything_IN_QUOTES(line)
				if is_EXTERNAL_GO_IMPORT(packname) {
					SAVE_IMPORT(packname, &IMPORT_LIST)
				}
				continue
			}
		} //end of inner for
	} // end of OUTER for

	SHOW_STRUCT(IMPORT_LIST)

	C.Print(PREFIX, "Total UNIQUE import packages: ")
	G.Println(len(IMPORT_LIST))

	return IMPORT_LIST
}
