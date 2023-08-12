/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	"strings"
	"os"
	"path/filepath"	
)


type FILE_INFO_OBJ struct {
	NAME 	string
	TYPE 	string
	IS_LINK 	bool
	LINK_ORIGIN string
	IS_DIR 		bool
	PATH 	string
	FULL_PATH string
	SIZE 	float64
	DATE 	string
	TIME 	string
}


func GET_FILE_INFO(filename string) FILE_INFO_OBJ {

	FINFO, err := os.Lstat(filename)
	if err != nil {
		R.Println("GET_FILE_INFO: ", err)
		return FILE_INFO_OBJ{}
	}

	var f FILE_INFO_OBJ
	f.NAME = FINFO.Name()
	f.LINK_ORIGIN = "n/a"
	
	//2b. Detect if this is a symlink	
	if FINFO.Mode()&os.ModeSymlink != 0 {
		f.IS_LINK = true
		f.TYPE = "SYMLINK"

		//since this is a link.. lets save the link ORIGIN
		origin_SOURCE, err43 := os.Readlink(filename)
		if err43 != nil {
			M.Println(" GET_FILE_INFO cant determine LINK ORIGIN: ", err43)
		} else {
			f.LINK_ORIGIN = origin_SOURCE
		}
	}
	//2c. Detect if DIRECTORY 
	if FINFO.IsDir() {
		f.IS_DIR = true
		f.TYPE = "DIRECTORY"
	}

	// Otherwise assume its a file
	if f.TYPE == "" {
		f.TYPE = "FILE"
	}

	//2d. Get the file SIZE
	f.SIZE = float64(FINFO.Size())
	f.TIME = FINFO.ModTime().String()

	f.FULL_PATH, _ = filepath.Abs(filename)
	f.PATH = filepath.Dir(filename)

	msplit := strings.Split(f.TIME, " ")
	tmpdate := msplit[0]

	ksplit := strings.Split(tmpdate, "-")
	year := ksplit[0]
	mon := ksplit[1]
	day := ksplit[2]
	year = GET_LAST_CHARs(year, 2)

	mfinal := mon + "-" + day + "-" + year
	f.DATE = mfinal

	return f
	
}