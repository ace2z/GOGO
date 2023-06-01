package CUSTOM_GO_MODULE

import (

	"strings"
	//"crypto/md5"
	//"encoding/hex"		

	
	//"github.com/dustin/go-humanize"
)



// Gets a numeric month from a text indicator
func GET_MONTH_NUM(month_name string) (string, int) {

	month_name = strings.ToLower(month_name)
	if len(month_name) < 3 {
		M.Println(" Invalid GET_MONTH name")
		return "", -6969
	}

	// ust get the first 3 chars .. (That way january matches jan)
	input :=  month_name[0:3]
	
	       if input == "jan" { return "01", 1
	} else if input == "feb" { return "02", 2
	} else if input == "mar" { return "03", 3
	} else if input == "apr" { return "04", 4
	} else if input == "may" { return "05", 5
	} else if input == "jun" { return "06", 6
	} else if input == "jul" { return "07", 7
	} else if input == "aug" { return "08", 8
	} else if input == "sep" { return "09", 9
	} else if input == "oct" { return "10", 10
	} else if input == "nov" { return "11", 11
	} else if input == "dec" { return "12", 12 }


	return "", -6969
}