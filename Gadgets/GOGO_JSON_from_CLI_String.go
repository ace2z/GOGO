package CUSTOM_GO_MODULE

import (
	"regexp"
)

/*
	var RECDATA = map[string]interface{} {
		"PREFIX" : THREAD_PREFIX,
		"MDATA" : m,
	}
*/

// Takes any string with json between "{}" and extracts the key value pairs based on what you pass as lookfor []string
// Returns a MAP of the key value pairs. EXAMPLE:
// ARM_CLIENT_ID := resmap["appId"].(string)
func Extract_JSON_from_CLI_STRING(tmp string, lookfor []string) map[string]interface{} {
	PLACEHOLDER()
	str := tmp

	re := regexp.MustCompile(`{([^}]*)}`) // Regular expression to match text inside brackets
	matches := re.FindAllStringSubmatch(str, -1)

	var PRE_RESULTS = ""
	for _, match := range matches {
		if len(match) <= 0 {
			continue
		}
		PRE_RESULTS = match[1]
	}

	// Now get everything match things between quotes
	re = regexp.MustCompile(`"(.*?)"`) // Regular expression to match text inside quotes
	qmatches := re.FindAllStringSubmatch(PRE_RESULTS, -1)
	dynamicMap := make(map[string]interface{})

	maxlen := len(qmatches) - 1
	for n, qmatch := range qmatches {
		if len(qmatch) <= 0 {
			continue
		}
		p := n + 1
		if p > maxlen {
			break
		}

		tmpmatch := qmatch[1]
		nextrow := qmatches[p]
		tmpval := nextrow[1]

		// Now lets iterate through lookfor and see if there is a match
		for _, findme := range lookfor {
			if findme == tmpmatch {
				dynamicMap[findme] = tmpval
			}
		}

	} //end of dynamicMap For

	/*
		//2. Put in a  recdata payload
		var RECDATA = map[string]interface{} {
			"PREFIX" : THREAD_PREFIX,
			"TRACK" : m.TRACK,
			"DATE" : m.Date,
			"URL" : m.HORSE_URL,
			"NAME" : m.HORSE_NAME,
			"FILENAME" : m.FILENAME,
			"SKIP_FILE" : m.SKIP_FILE,
		}
	*/

	return dynamicMap

} //end of
