/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package GOGO_MDC

import (

	// = = = = = Native Libraries
	"strings"
	"time"

	//"math/rand"

	// = = = = = CUSTOM Libraries

	. "github.com/ace2z/GOGO/Gadgets"
	// = = = = = 3rd Party Libraries
)

func get_TZ_OBJECT(TZ_to_use string) (bool, *time.Location) {
	var TIMEZONE_OBJ = UTC_Location_OBJ

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

		case "local":
			TIMEZONE_OBJ = LOCAL_Location_OBJ
			is_valid_input = true
			break
		}
	}

	return is_valid_input, TIMEZONE_OBJ
}

func FIX_if_needed(input string) string {

	// error handling..we only want to FIX the single digits
	if len(input) != 1 {
		return input
	}
	tmp_num := STRING_to_INT(input)

	if tmp_num < 10 {

		fixed := "0" + input
		return fixed
	}

	// otherwise return
	return input
}

func vsplit_delims(r rune) bool {
	return r == '_' || r == ','
}

// checks for a verbose Date that looks like this: June_1,_2023
// May 28, 2023
func check_for_VERBOSE_DATE(input string) (bool, string, string, string, string, string, string) {

	var sMon, sDay, sYear string
	var sHour = "00"
	var sMin = "00"
	var sSec = "00"

	sd := strings.FieldsFunc(input, vsplit_delims)
	mon_string, _ := GET_MONTH_NUM(input)

	// We have a verbose date! lets splitit out
	if mon_string != "" && len(sd) >= 1 {

		sMon = mon_string
		sDay = FIX_if_needed(sd[1])
		sYear = sd[2]
	} else {
		return false, sMon, sDay, sYear, sHour, sMin, sSec
	}

	return true, sMon, sDay, sYear, sHour, sMin, sSec
}

func conv_SPLIT_delims(r rune) bool {
	return r == '@' || r == ':' || r == '_' || r == '-' || r == '/' || r == ' '
}
func have_SUPPORTED_DEFAULT_FORMAT(inputDate string) (bool, map[string]interface{}) {

	var resMAP map[string]interface{}

	sMon := ""
	sDay := ""
	sYear := ""
	sHour := "0"
	sMin := "0"
	sSec := "0"

	// First check to see if we have a date in verbose format: June 1, 2023
	is_VERBOSE := false
	is_VERBOSE, sMon, sDay, sYear, sHour, sMin, sSec = check_for_VERBOSE_DATE(inputDate)

	//1. If it is NOT verbose...and is a MM/DD/YYYY   (or YYYY/MM/DD) format.. proceed with split
	if is_VERBOSE == false {

		//1. Do a split..
		sd := strings.FieldsFunc(inputDate, conv_SPLIT_delims)
		//error handling check to see if we have enough items
		if len(sd) < 3 {
			return false, resMAP
		}

		// This will fix the parts of the date... if we have 5 .. we get 05 in return
		part_a := sd[0]
		part_b := sd[1]
		part_c := sd[2]

		//2. Check for british format first
		skip_check_time := false
		if len(part_a) == 4 && (len(part_b) == 2 || len(part_b) == 1) {

			sYear = strings.TrimSpace(part_a)

			sMon = FIX_if_needed(part_b)
			sDay = FIX_if_needed(part_c)

			// Now lets see if we have a full british format date with time
			if len(sd) >= 5 {
				sHour = FIX_if_needed(sd[3])
				sMin = FIX_if_needed(sd[4])
				skip_check_time = true
			}

			//3. else check for normal xx/yy/zzzz  format
		} else if (len(part_a) == 2 || len(part_a) == 1) && (len(part_b) == 2 || len(part_b) == 1) && (len(part_c) == 4 || len(part_c) == 2) {
			sMon = FIX_if_needed(part_a)
			sDay = FIX_if_needed(part_b)
			sYear = part_c

			// If they passed a 2 digit year, fix this by prefixing it with 20
			if len(part_c) == 2 {
				sYear = "20" + part_c
			}

		} else {
			return false, resMAP
		}

		//4. Now determine if we have a TIME appended.. via the :
		if skip_check_time == false {
			if strings.Contains(inputDate, ":") {
				sHour = strings.TrimSpace(sd[3])
				sMin = strings.TrimSpace(sd[4])

				//4b.. If seconds was appended.. lets add those as well
				if strings.Count(inputDate, ":") == 2 {
					sSec = strings.TrimSpace(sd[5])
				}
			}
		}
	}

	//5. Finally make the map
	resMAP = map[string]interface{}{
		"month": sMon,
		"day":   sDay,
		"year":  sYear,
		"hour":  sHour,
		"min":   sMin,
		"sec":   sSec,
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
		resMAP = map[string]interface{}{
			"month": strings.TrimSpace(sd[2]),
			"day":   strings.TrimSpace(sd[3]),
			"year":  strings.TrimSpace(sd[4]),
			"hour":  strings.TrimSpace(sd[5]),
			"min":   strings.TrimSpace(sd[6]),
			"sec":   strings.TrimSpace(sd[7]),
		}
	} else {
		return false, resMAP
	}
	return true, resMAP
}
