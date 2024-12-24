/*   GOGO_Math / Date / Conversion Gadget - Useful math and Calucation code to make Go Dev Easier

---------------------------------------------------------------------------------------
NOTE: For Functions or Variables to be globally availble. The MUST start with a capital letter.
	  (This is a GO Thing)


	Aug 28, 2021    v1.23   - Initial Rollout

*/

package CUSTOM_GO_MODULE

import (

	// = = = = = Native Libraries
	"time"
	//"math/rand"
	// = = = = = CUSTOM Libraries
	//. "github.com/ace2z/GOGO/Gadgets"
	//. "github.com/ace2z/GOGO/Gadgets/StringOPS"
	// = = = = = 3rd Party Libraries
)

// Returns a DATE_OBJ as UnixTime ... and MILLISECONDS
func MAKE_Epoch_TIMESTAMP(dateobj time.Time) (int64, int64) {

	//	pretty, _ := SHOW_PRETTY_DATE(dateobj, "basic")
	newUNIX := dateobj.Unix()
	newMILLI := dateobj.UnixMilli()

	return newUNIX, newMILLI
}

func EPOCH_to_DATEOBJ(unix int64) time.Time {

	newDATE := time.Unix(unix, 0)

	return newDATE
}
