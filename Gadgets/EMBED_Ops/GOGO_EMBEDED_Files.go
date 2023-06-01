/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GOMOD

import (
	"os"
	"strings"

	. "github.com/ace2z/GOGO/Gadgets"
)

/*
	This is how you use EMBEDDED support with GO ... The below items are for EMBED support.. 
	be sure to specify the files you want to embed IN the go file OUTSIDE a function
	
	1. Add this to the Top of your file (use a seperate file from MAIN
	_ "embed"

	2. EXAMPLE below embeds the steam_plugin_isntall_helper into the variable steam_plugin_helper
	   Add these delcreations with the path to the file you wish to embed
	   They will be made available as vars you see below

//go:embed _EMBEDDED_FILES/steam_plugin_install_helper.sh
var steam_plugin_helper string				// For TEXT files

//go:embed _EMBEDDED_FILES/ca.pem.gz
var pem []byte								// For BINARY files

    3. When all the embeddeds have been declared, you can use SAVE_EMBED (to save the EMBEDED files to the EMBED_FILES Structure)
*/
type EMBED_OBJ struct {
	Name 	        string
	TYPE			string

	DATA  			[]byte
	TEXT			string
	// BIN 			[]byte
}

var EMBED_FILES []EMBED_OBJ

// This saves items in EMBED_FILES.. and always returns totall number of items in EMBED so far
func SAVE_EMBED(filename string, fileOBJ interface{} ) int {

	var T EMBED_OBJ
	T.Name = filename

	// Determine the type of fileOBJ (either string or byte)
	STRING_TEXT, FOUND_string := fileOBJ.(string)
	BYTE_DATA, FOUND_bytes := fileOBJ.([]byte)

	if FOUND_string {
		T.TYPE = "text"
		T.DATA = []byte(STRING_TEXT)


	} else if FOUND_bytes {

		T.TYPE = "binary"
		T.DATA = BYTE_DATA
	} else {

		R.Println(" Cant Save Embed for some reason! Weird or Invalid Data Type??")
		return 0
	}

	EMBED_FILES = append(EMBED_FILES, T)	

	return len(EMBED_FILES)
}
func WRITE_EMBEDDED(filename string, destpath string, EXTRA_ARGS ...interface{}) {

	var alt_filename = GET_EXTRA_ARG(0,  EXTRA_ARGS...).(string)
	var verbose_mode  = GET_EXTRA_ARG("verbose", EXTRA_ARGS...).(bool)

	for _, f := range EMBED_FILES {

		if f.Name == filename {

			var FULL_FILE_PATH = destpath + "/" + f.Name
		
			if alt_filename != "" {
				FULL_FILE_PATH = destpath + "/" + alt_filename
			}


			// Also safety remove file
			err1 := os.RemoveAll(FULL_FILE_PATH)

			// This will only fail if we dont have permission or file isnt there
			if err1 != nil {

				if verbose_mode {
					R.Println(" WRITE_EMBED Error: ", err1)
					return
				}
			}

			 // Text written as 644
			 if f.TYPE == "text" {
				if err := os.WriteFile(FULL_FILE_PATH, f.DATA, 0644); err != nil {
					if verbose_mode {
						R.Print("ERROR! Cant save the file: ")
						Y.Println(FULL_FILE_PATH)
						R.Println(err)
					}
				}
			 }


			 // Binarys are written as 755
			 if f.TYPE == "binary" {
				if err := os.WriteFile(FULL_FILE_PATH, f.DATA, 0755); err != nil {
					if verbose_mode {
						R.Print("ERROR! Cant save the file: ")
						Y.Println(FULL_FILE_PATH)
						R.Println(err)
					}
				}
			 }


			 /// Also check to see if this was a gz file.. if so, we extract it
			 if strings.Contains(f.Name, ".gz") {

				NOGZ_FILENAME := strings.Replace(FULL_FILE_PATH, ".gz", "", 1)
				os.RemoveAll(NOGZ_FILENAME)

				RUN_COMMAND("gunzip " + FULL_FILE_PATH )
			 }

			return
		}

	}


	// else if they specified somethign we didnt have, tell them

	R.Println(" ERRROR! Cant find the specified Embedded file!!")
}

func SHOW_EMBED() {

	C.Print("\n Showing EMBEDDED Files: ")
	G.Println(len(EMBED_FILES))
	C.Println("")

	for _, v := range EMBED_FILES {
		
		if v.TYPE == "text" {
			Y.Print("     ", v.TYPE, " - ")
		} else {
			Y.Print("   ", v.TYPE, " - ")
		}
		
		C.Println(v.Name)
	}

	Y.Print("\n Total Embedded Files: ")
	C.Println(len(EMBED_FILES))	

	C.Println("")
}
