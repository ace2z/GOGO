/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

	 MISC Functions that i sometimes use.. but shouldnt be in the main file

*/

package CUSTOM_GO_MODULE

import (
	"os"
	"time"
	"io"
	"io/ioutil"
	"encoding/hex"
	"math"
	"crypto/md5"
	"bufio"
	"path/filepath"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"	
	"strings"
)



func DOES_FILE_CONTAIN(filename string, lookfor string) bool {
	b, err := ioutil.ReadFile(filename)
    if err != nil {
		M.Print("DOES_FILE_CONTAIN Error: ")
		Y.Println(err)
		return false        
    }
	
    found_string := string(b)
	if strings.Contains(found_string, lookfor) {
		return true
	}


	return false
}

// Easy way to get teh MD5 of a file
func GET_FILE_MD5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string
	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		R.Println(" GET MD5 File ERR", err)
		return returnMD5String, err
	}
	//Tell the program to call the following function when the current function returns
	defer file.Close()
	//Open a new hash interface to write to
	hash := md5.New()
	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		R.Println(" GET MD5 File ERR", err)
		return returnMD5String, err
	}
	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]
	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}



func TOUCH_FILE(filename string) {
	currentTime := time.Now().Local()
	err := os.Chtimes(filename, currentTime, currentTime)
	if err != nil {
		R.Println(" ERROR Touching File!!", err)
	}
}

func makeRound(num float64) int {
	return int(num + math.Copysign(0.5, num))
}




/*
	= = = = =
	= = = = =  Some useful 'does exist' helper functions
	= = = = =
*/

// Determines if a file is a sym link or not
func IS_FILE_LINK(filename string) bool {

	FINFO, err := os.Lstat(filename)
	if err != nil {
		R.Println("IS_FILE_LINK: ", err)
		return false
	}

	//2b. Detect if this is a symlink
	if FINFO.Mode()&os.ModeSymlink != 0 {
		_, err43 := os.Readlink(filename)
		if err43 != nil {
			R.Println(" IS_FILE_LINK Detect SymLink err: ", err43)
			return false
		}

		// Otherwise we have asymlink!
		// Return true and the SOURCE the file points to

		return true
	}	

	return false
}
func HAVE_LINK(input string) bool {
	return IS_FILE_LINK(input)
}
func IS_LINK(input string) bool {
	return IS_FILE_LINK(input)
}

// Also gets the LINK_ORIGIN
func LINK_ORIGIN(filename string) string {
	if HAVE_LINK(filename) {

		FINFO, err := os.Lstat(filename)
		if err != nil {
			R.Println("LINK_ORIGIN err: ", err)
			return ""
		}
	
		//2b. Detect if this is a symlink
		if FINFO.Mode()&os.ModeSymlink != 0 {
			origin_SOURCE, err43 := os.Readlink(filename)
			if err43 != nil {
				R.Println(" LINK_ORIGIN err: ", err43)
				return ""
			}
	
			// Otherwise we have a symlink so return the ORIGIN  of the link

			return origin_SOURCE
		}			

	} else {
		Y.Println(" WARNING: .. this is NOT a SYMLINK!!")		
	}

	return filename
}


func CURR_DIR() string {
	currdir, err := os.Getwd()
    if err != nil {
		M.Print("CURR_DIR Cant determine current directory: ")
		Y.Println(err)
		return ""
    }
	return currdir
}

func HAVE_DIR(input string) bool {
	file, err := os.Open(input)
	if err != nil {
		Y.Println("Warning: ", err)
		return false
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		Y.Println("Warning: ", err)
	}
	
	if fileInfo.IsDir() {
		return true
	}
	
	return false
}
func IS_DIR(input string) bool {
	return HAVE_DIR(input)
}


// This checks to see if a file or directory exists
func FILE_EXISTS(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
} // end of fileExist
// Returns file size in MEGS and BYTES
func GET_FILE_SIZE(bytes int64) (float64, float64) {
	megs := float64(bytes / 1024000)
	gigs := float64(megs / 1024)

	precision := 2

	output := math.Pow(10, float64(precision))
	
	megs = float64(makeRound(megs*output)) / output	
	gigs = float64(makeRound(gigs*output)) / output	

	return megs, gigs
}


// Convenient way to write to a file. either append or overwrite
func WRITE_FILE(FULL_FILENAME string, TEXT_for_FILE string, ALL_PARAMS ...string) bool {

	var verbose = false
	var doOverwrite = false

	for x, VAL := range ALL_PARAMS {

		//1. First param specified, means we OVERWRITE the file (instead of the default which is append)
		if x == 0 {
			if VAL == "verbose" {
				verbose = true
			} else if VAL == "overwrite" || VAL == "true" {
				doOverwrite = true
			}
		}
	} // end of for

	//2. If set, we erase the file before writing to it
	if doOverwrite {
		if FILE_EXISTS(FULL_FILENAME) {
			err := os.Remove(FULL_FILENAME)
			if err != nil {
				if verbose {
					R.Println(" WRITE_FILE weird error checking for exists: ", err)
				}
				return false
			}
		}
	}

	file_OBJ, err := os.OpenFile(FULL_FILENAME, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		if verbose {
			R.Print("*** WRITE_FILE ERROR: ")		
			Y.Println(err)
		}
		return false
	}


	//3. Otherwise, write the line we have TO the file specified
	datawriter := bufio.NewWriter(file_OBJ)     
	_, _ = datawriter.WriteString(TEXT_for_FILE + "\n")
	datawriter.Flush()
	file_OBJ.Close()

	if verbose {
		SHOW_BOX("Writing TO file named: ", "|green|" + FULL_FILENAME)	
	}


	return true
}




/*
 DownloadFile will download a url to a local file. It's efficient because it will
 write as it downloads and not load the whole file into memory.

  Courtesy of: https://golangcode.com/download-a-file-from-a-url/
*/
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}



/*
 Saves a Struct to DISK: 
 Inspired by:
   https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d
*/

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
	return nil, err
	}
	return bytes.NewReader(b), nil
}







/*
EXTRA_ARGS ...interface{}) {

	var have_alt, alt_val = GET_EXTRA_ARG(0,  EXTRA_ARGS...)
	var verbose, _  = GET_EXTRA_ARG("verbose", EXTRA_ARGS...)
*/
func SAVE_Struct_2_DISK(dest_file string, v interface{}) error {
	// var lock sync.Mutex
	// lock.Lock()
	// defer lock.Unlock()	
	var verbose = true


	Y.Print(" SAVING Struct to file: ")
	G.Println(dest_file)
	
	var justPATH = filepath.Dir(dest_file)
	// fd, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
    // var justFILE = filepath.Base(filename)	

	// Error Handling, lets safety create the parent directory
	os.MkdirAll(justPATH, 0777)

	f, err := os.Create(dest_file)
	if err != nil {
		if verbose {
	  		R.Println(" Error IN SaveStruct: ", err)
		}
	  return err
	}
	defer f.Close()


	r, err := Marshal(v)
	if err != nil {
		if verbose {
	  		R.Println("Error in SaveStruct: ", err)
		}
	  return err
	}
	_, err = io.Copy(f, r)

	if err != nil {
		if verbose {
			R.Println(" Error in SaveStruct: ", err)
		}
	}


	return err
}

// Unmarshal is a function that unmarshals the data from the
// reader into the specified value.
// By default, it uses the JSON unmarshaller.
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}





func DOES_FILE_EXIST(FULL_path_2_file string, ALL_PARAMS ...string) bool{

	var verbose = false 

	for x, VAL := range ALL_PARAMS {

		//1. First param specified, means we OVERWRITE the file (instead of the default which is append)
		if x == 0 {
			if VAL == "verbose" {
				verbose = true
				continue
			}						
			

		} 
	} // end of for	

	//2. Determine if file exists:
	if _, err := os.Stat(FULL_path_2_file); err == nil {

		return true

	} else if errors.Is(err, os.ErrNotExist) {

		if verbose {
			R.Print(" ERROR in DOES_FILE_EXIST: ")
			Y.Println(err)
		}
		// path/to/whatever does *not* exist
		return false

	} else {
		if verbose {
			R.Print(" ERROR in DOES_FILE_EXIST (weird ERROR):")
			Y.Println(err)
		}


		return false
	}	

} //end of 


// Load loads the file at path into v.
// Use os.IsNotExist() to see if the returned error is due
// to the file being missing.
func LOAD_Struct_from_FILE(FULL_path_2_file string, v interface{}, bequiet bool) error {

	if bequiet == false {
		Y.Print(" Reading STRUCT from: ")
		W.Println(FULL_path_2_file)
	}

	// var lock sync.Mutex
	// lock.Lock()
	// defer lock.Unlock()
	f, err := os.Open(FULL_path_2_file)
	if err != nil {
		R.Print(" ERROR! during OPEN")
		Y.Println(err)
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}






// Opens a file and returns a file object
func OPEN_FILE(path_to_file string) *os.File {

	file_obj, err := os.Open(path_to_file)
	if err != nil {
		R.Print(" ** ERROR ** Cannot open the file: ")
		W.Println(path_to_file)
		Y.Println(err.Error())
	}

	return file_obj

} //end of func