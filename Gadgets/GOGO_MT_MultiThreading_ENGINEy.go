package CUSTOM_GO_MODULE

/*
	Finally got excellent universal multi threading engine

*/
import (
	"sync"
)

type MT_THREAD_OBJ struct {
	THREAD_PREFIX string
	DATA          interface{}
	EXTRA_FLAG    string
}

type CHILD_IMPL_FUNC func(MT_THREAD_OBJ, *sync.WaitGroup)

/*
Multi Thread Engine! Takes in a function implemented like CHILD_IMP_FUNC .. as a parameter
EXAMPLE:

	var chimp = func(item MT_THREAD_OBJ, WAIT_GROUP *sync.WaitGroup) {

		//	if using a a struct, expand data like this
		//	If DATA is just a native string , int.. use
		//	pobj := item.DATA.(string)

		pobj := item.DATA.(PRICE_TA_OBJ)
		prefix := item.THREAD_PREFIX

		C.Println(prefix, " ")
		Y.Print(" PTA_hey Symbole is: ")
		W.Print(" EXTRA FLAG: ", item.EXTRA_FLAG)
		G.Println(pobj.SYMBOL)

		WAIT_GROUP.Done() //call this.. tells the MT Engine this thread is DONE
	}
	// pass a number for total number of threads to use.. pass true/false to "sleep" between go routine invocations
*/
func MT_Engine(desc string, MASTER_LIST []MT_THREAD_OBJ, child CHILD_IMPL_FUNC, PARAMS ...interface{}) {
	var sleep_between_threads = false
	var MAX_THREADS = 10
	var PASSED_EXTRA = ""

	// Get PARAMS
	for _, field := range PARAMS {
		val_int, IS_INT := field.(int)
		//val_float, IS_FLOAT := field.(float64)
		val_string, IS_STRING := field.(string)
		val_bool, IS_BOOL := field.(bool)

		// If an INT is passed, its always going to be the number of threads
		if IS_INT {
			MAX_THREADS = val_int
			continue
		}

		// if a BOOL is passed, this meqns we "sleep between go routeine threads"
		if IS_BOOL {
			sleep_between_threads = val_bool
			continue
		}

		// If a string is passed,  this is an extra_flag we pass along to the child function. The implemenation can do something with this if necessary
		if IS_STRING {
			PASSED_EXTRA = val_string
			continue
		}
	}

	//1. Get started
	SHOW_BOX(desc)

	var TOTAL_ITEMS_TO_PROCESS = len(MASTER_LIST)
	if MAX_THREADS > TOTAL_ITEMS_TO_PROCESS {
		MAX_THREADS = TOTAL_ITEMS_TO_PROCESS
	}
	var WAIT_GROUP sync.WaitGroup
	WAIT_GROUP.Add(MAX_THREADS)

	var THREAD_COUNTER = 0 // MUST ALWAYS START AT 0

	//2. Master Loop. mark the time this STARTS
	start := MARK_START_TIME()

	for _, item := range MASTER_LIST {

		tleft_num := TOTAL_ITEMS_TO_PROCESS - 1
		text_REMAINING_ITEMS := ShowNum(tleft_num)
		text_CURR_THREAD := ShowNum(THREAD_COUNTER + 1)
		item.THREAD_PREFIX = "[ Thrd: " + text_CURR_THREAD + " of " + INT_to_STRING(MAX_THREADS) + ", " + text_REMAINING_ITEMS + " left ] "

		// Now if an extra flag is PASSED, we have to decide if we will use what was PASSED.
		//but first check to see if we already have EXTRA_ITEM set in the list item
		if item.EXTRA_FLAG == "" && PASSED_EXTRA != "" {
			item.EXTRA_FLAG = PASSED_EXTRA
		}
		//2. Maks a map with payload you can pass to the child. Not really needed anymore
		// var DYNDATA = map[string]interface{}{
		// 	"SYMBOL":        tmp_symbol,
		// 	"THREAD_PREFIX": THREAD_PREFIX,
		// 	"PULL_THIS":     pullthis,
		// }

		go child(item, &WAIT_GROUP)
		THREAD_COUNTER++
		TOTAL_ITEMS_TO_PROCESS--

		// sleep delay between thread invocation
		if sleep_between_threads {
			Sleep(1, false)
		}

		//3. Error Hanlding if we are on the LAST item to process
		if TOTAL_ITEMS_TO_PROCESS == 0 {
			WAIT_GROUP.Wait()
			break
		}

		//4. Now lets pause if we hit our MAX_THREAD
		if THREAD_COUNTER == MAX_THREADS {
			THREAD_COUNTER = 0

			//4b. Safety make sure we arent above our total items LEFT to process
			if MAX_THREADS > TOTAL_ITEMS_TO_PROCESS {
				MAX_THREADS = TOTAL_ITEMS_TO_PROCESS
			}

			//5. Now.. lets wait.. while all the threads finish running
			WAIT_GROUP.Wait()

			//6. Now.. after done waiting.. lets queue up the next group of threads
			WAIT_GROUP.Add(MAX_THREADS)
		} //end of if

	} //end of master for

	MARK_END_TIME(start)

	Y.Println("")
	Y.Println(" ****************************************************** ")
	Y.Println(" *** NOTE: All THREADS HAVE COMPLETED !!")
	C.Println(" ****************************************************** ")
}
