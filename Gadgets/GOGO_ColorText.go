/*   GOGO_Gadgets  - Useful multi-purpose GO functions to make GO DEV easier
	 by TerryCowboy

*/

package CUSTOM_GO_MODULE

import (

	"github.com/fatih/color"

)

// -=-=-= COMMON COLOR GLOBAL references =-=-=-=-

var R = color.New(color.FgRed, color.Bold)
var G = color.New(color.FgGreen, color.Bold)
var Y = color.New(color.FgYellow, color.Bold)
var B = color.New(color.FgBlue, color.Bold)
var M = color.New(color.FgMagenta, color.Bold)
var C = color.New(color.FgCyan, color.Bold)
var W = color.New(color.FgWhite, color.Bold)

// This shows with UNDERLINE
var Ru = color.New(color.FgRed, color.Underline)
var Gu = color.New(color.FgGreen, color.Underline)
var Yu = color.New(color.FgYellow, color.Underline)
var Bu = color.New(color.FgBlue, color.Underline)
var Mu = color.New(color.FgMagenta, color.Underline)
var Cu = color.New(color.FgCyan, color.Underline)
var Wu = color.New(color.FgWhite, color.Underline)


// Some helpful REVERSe colors

var RY = color.New(color.FgRed, color.BgYellow, color.Bold)
var GB = color.New(color.FgGreen, color.BgBlue, color.Bold)
var BW = color.New(color.FgBlue, color.BgWhite, color.Bold)
var MW = color.New(color.FgMagenta, color.BgWhite, color.Bold)

var WB = color.New(color.FgWhite, color.BgBlue, color.Bold)
var WM = color.New(color.FgWhite, color.BgMagenta, color.Bold)
var WG = color.New(color.FgWhite, color.BgGreen, color.Bold)

var YB = color.New(color.FgYellow, color.BgBlue, color.Bold)
var YR = color.New(color.FgYellow, color.BgRed, color.Bold)
var YM = color.New(color.FgYellow, color.BgMagenta, color.Bold)

// Reverse w UNDERLINE.. Only setting up the ones that work best on Black background terminals
var RYu = color.New(color.FgRed, color.BgYellow, color.Bold, color.Underline)
var GBu = color.New(color.FgGreen, color.BgBlue, color.Bold, color.Underline)
var BWu = color.New(color.FgBlue, color.BgWhite, color.Bold, color.Underline)

var WBu = color.New(color.FgWhite, color.BgBlue, color.Bold, color.Underline)
var WMu = color.New(color.FgWhite, color.BgMagenta, color.Bold, color.Underline)
var WGu = color.New(color.FgWhite, color.BgGreen, color.Bold, color.Underline)

var YBu = color.New(color.FgYellow, color.BgBlue, color.Bold, color.Underline)
var YRu = color.New(color.FgYellow, color.BgRed, color.Bold, color.Underline)
var YMu = color.New(color.FgYellow, color.BgMagenta, color.Bold, color.Underline)



// This makes foreground and background color text
/*
func SetForeBack_COLOR(fore string, back string) {
	var s_fore = color.FgYellow
	var s_back = color.FgMagenta

	if fore == "yellow" { s_fore = color.Set(color.FgYellow) }
	if fore == "cyan" { s_fore = color.Set(color.FgYellow) }

	/ Use handy standard colors
	color.Set(s_fore, s_back, color.Bold)
	

	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")
	
	color.Unset() // Don't forget to unset	
}
*/


