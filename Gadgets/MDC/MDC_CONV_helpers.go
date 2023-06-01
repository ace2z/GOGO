/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)

	
	Aug 28, 2021    v1.23   - Initial Rollout

*/

package GOGO_MDC


import (

	// = = = = = Native Libraries
		"time"
		"strings"
		//"math/rand"

	// = = = = = CUSTOM Libraries

	//. "github.com/acedev0/GOGO/Gadgets"


	// = = = = = 3rd Party Libraries

)




func get_TZ_OBJECT(TZ_to_use string) (bool, *time.Location) {
	var TIMEZONE_OBJ = time.Local

	is_valid_input := false

	//2. Time Zone logic
	if TZ_to_use != "" {
		switch TZ_to_use {
			case "est":
				TIMEZONE_OBJ = EST_Location_OBJ
				is_valid_input = true
				break

			case "cst":
				TIMEZONE_OBJ = CST_Location_OBJ
				is_valid_input = true
				break

			case "mdt":
				TIMEZONE_OBJ = MST_Location_OBJ
				is_valid_input = true
				break

			case "mst":
				TIMEZONE_OBJ = MST_Location_OBJ
				is_valid_input = true
				break				

			case "pst":
				TIMEZONE_OBJ = PST_Location_OBJ
				is_valid_input = true
				break

			case "utc":
				TIMEZONE_OBJ = UTC_Location_OBJ
				is_valid_input = true
				break
		}
	}

	return is_valid_input, TIMEZONE_OBJ
}

func conv_SPLIT_delims(r rune) bool {
    return r == '@' || r == ':' || r == '_' || r == '-' || r == '/'
}
func have_DEFAULT_FORMAT(inputDate string) (bool, map[string]interface{}) {

	var resMAP map[string]interface{}
	
	//1. Do a split.. 
	sd := strings.FieldsFunc(inputDate, conv_SPLIT_delims)	
	//error handling check to see if we have enough items
	if len(sd) < 3 {
		return false, resMAP
	}

	sMon := ""
	sDay := ""
	sYear := ""
	sHour := "0"
	sMin := "0"
	sSec := "0"

	//2. Check for british format first
	if len(sd[0]) == 4 && len(sd[1]) == 2 {

		sYear = strings.TrimSpace(sd[0])
		sMon = strings.TrimSpace(sd[1])
		sDay = strings.TrimSpace(sd[2])		

	//3. else check for normal xx/yy/zzz  format
	} else if len(sd[0]) == 2 && len(sd[1]) == 2 && len(sd[2]) == 4 {
		sMon = strings.TrimSpace(sd[0])
		sDay = strings.TrimSpace(sd[1])
		sYear = strings.TrimSpace(sd[2])
	} else {
		return false, resMAP
	}

	
	//4. Now determine if we have a TIME appended.. via the :
	if strings.Contains(inputDate, ":") {
		sHour = strings.TrimSpace(sd[3])
		sMin = strings.TrimSpace(sd[4])

		//4b.. If seconds was appended.. lets add those two
		if strings.Count(inputDate, ":") == 2 {
			sSec = strings.TrimSpace(sd[5])
		}
	}

	//5. Finally make the map
	resMAP = map[string]interface{} {
		"month" : sMon,
		"day" : sDay,
		"year" : sYear,
		"hour": sHour,
		"min": sMin,
		"sec": sSec,
	}	

	return true, resMAP
}

func epoch_SPLIT_delims(r rune) bool {
    return r == ' ' || r == ':' || r == '-'
}
// EPOCH format looks like this: must confrom EPOCH_Time_CONVERT
// Mon Jan 02-01-2006 15:04:05
func have_EPOCH_FORMAT(input string) (bool, map[string]interface{}) {
	valid_score := 0
	var resMAP map[string]interface{}

	if strings.Count(input, ":") == 2 {
		valid_score++
	}
	if strings.Count(input, "-") == 2 {
		valid_score++
	}
	if strings.Count(input, " ") == 3 {
		valid_score++
	}
	if valid_score == 3 {
		sd := strings.FieldsFunc(input, epoch_SPLIT_delims)
		resMAP = map[string]interface{} {
			"month" : strings.TrimSpace(sd[2]),
			"day" : strings.TrimSpace(sd[3]),
			"year" : strings.TrimSpace(sd[4]),
			"hour": strings.TrimSpace(sd[5]),
			"min": strings.TrimSpace(sd[6]),
			"sec": strings.TrimSpace(sd[7]),
		}
	} else {
		return false, resMAP
	}
	return true, resMAP
}
