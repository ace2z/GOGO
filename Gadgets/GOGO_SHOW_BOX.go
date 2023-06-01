package CUSTOM_GO_MODULE

import (
	"math"
	"strings"
	"unicode/utf8"
	
)


/*
  This padds a string that is passed and makes it a total LENGH of TOTAL_FINAL_LEN
  Pass a character as a string param and it will pad with THAT char instead of space
*/
func PAD_STRING(input string, TOTAL_FINAL_LEN int, ALL_PARAMS ...string) string {

	// This is the default padd string
	var padString = " "

	//2. If you want to use another just pass it as a string
	for p, VAL := range ALL_PARAMS {

		if p == 0 {
			padString = VAL

			//2b. also add space to before and after input so the pad string char isnt running up against it
			input = " " + input + " "
		}

	}	

	// Courtesy of: https://gist.github.com/asessa/3aaec43d93044fc42b7c6d5f728cb039

	var padLength = TOTAL_FINAL_LEN
	var inputLength = len(input)
	var padStringLength = len(padString)

	length := (float64(padLength - inputLength)) / float64(2)
	repeat := math.Ceil(length / float64(padStringLength))
	output := strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]

	return output
}


func DELETE_from_LIST[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}
func REMOVE_from_LIST[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}
func DELETE_ITEM[T any](slice []T, s int) []T {
    return append(slice[:s], slice[s+1:]...)
}


// This has limited color support and if it receives a string with the color in |cyan| format..it prints in that color
func helper_SHOW_with_COLOR(input string, SHOW_OUTPUT bool) (string, string) {

	// var JUST_STRING = temps[2]
	var JUST_STRING = input
	var COLOR = ""

	if strings.Count(JUST_STRING, "|") == 2 {
		temps := strings.Split(input, "|")
		COLOR = temps[1]
		JUST_STRING = temps[2]

		if SHOW_OUTPUT {

			switch COLOR {
			case "cyan":
				C.Print(JUST_STRING)
				break

			case "green":
				G.Print(JUST_STRING)
				break

			case "bluewhite":
				BW.Print(JUST_STRING)
				break

			case "yellow":
				Y.Print(JUST_STRING)
				break
			case "red":
				R.Print(JUST_STRING)
				break				
			default:
				W.Print(JUST_STRING)
				break
			}
		}

		COLOR = "|" + COLOR + "|"

	} else {
		if SHOW_OUTPUT {
			W.Print(JUST_STRING)
		}
	}

	return JUST_STRING, COLOR

} // end of


/*
	This is a nice way of showing a message in a box
	Just pass each line you want in the box as a seperate parameter

╭――――――――――――――――――╮
│                  │
│                  │
│                  │
╰――――――――――――――――――╯
*/
var BOX_INDENT_SPACES = "         "
func SHOW_BOX(ALL_PARAMS ...string) {

	var lines []string

	var SPACE_PREFIX = "       "

	//1. if multiple lines are passed, lets iterate through them
	for _, VAL := range ALL_PARAMS {
		lines = append(lines, VAL)
	}

	//2. FIgure out which line is the LONGEST.. this is how we grow the box length
	largest_len := 0

	for _, l := range lines {

		var JUST_STRING, _ = helper_SHOW_with_COLOR(l, false)
		temp_len := len(JUST_STRING)

		if temp_len > largest_len {
			largest_len = temp_len
		}
	} //end of line len determine for

	largest_len += len(SPACE_PREFIX) + 4

	//3. Now drop top of box
	var BOX_TOP = "┌"
	var BOX_BOTTOM = "└"
	for x := 0; x < largest_len; x++ {
		BOX_TOP += "─"
		BOX_BOTTOM += "─"
	}
	//4. CLose the ends of the BOX
	BOX_TOP += "┐"
	BOX_BOTTOM += "┘"

	//5. MUST use the utf8 way to get the string length since it contains ASCII chars
	var BOX_LEN = utf8.RuneCountInString(BOX_TOP)
	BOX_LEN = BOX_LEN - 2 // We have to do BOXLEN-2 to account for the Right and Left angle Brakcets

	//6. Top of Box
	M.Println(SPACE_PREFIX + BOX_TOP)

	//7. Prints the Lines in between top and bottom
	for _, line := range lines {

		var temp_full_line, MCOLOR = helper_SHOW_with_COLOR(line, false)

		// Most likely the temp_full line is LESS than the BOX_LEN.. so lets padd it
		if len(temp_full_line) < BOX_LEN {
			temp_full_line = PAD_STRING(temp_full_line, BOX_LEN)
		}

		M.Print(SPACE_PREFIX + "│")

		helper_SHOW_with_COLOR(MCOLOR+temp_full_line, true)

		M.Println("│")
	}

	//8. Prints the BOTTOM of box.. we are DONE
	M.Println(SPACE_PREFIX + BOX_BOTTOM)

	//9. Add an indent so the next thing that is C.Println'd ... will be indented under the box
	//C.Println("")


} //end of SHOW_BOX_MESSAGE
