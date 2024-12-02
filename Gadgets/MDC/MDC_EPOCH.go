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

func MAKE_Epoch_TIMESTAMP(dateobj time.Time) int {

	//	pretty, _ := SHOW_PRETTY_DATE(dateobj, "basic")
	newEpoch := dateobj.Unix()
	res_ts := int(newEpoch)

	return res_ts
}
