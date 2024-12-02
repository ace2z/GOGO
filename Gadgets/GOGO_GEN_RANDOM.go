package CUSTOM_GO_MODULE

import (
	"math/rand"
)

// Generates random range of FLOATS from MIN to MAX
func Gen_RAND_Float(min float64, max float64) float64 {

	res_float := min + rand.Float64()*(max-min)

	return res_float

} //end of genRandomRange

// Generates random range of INTS from MIN to MAX
func Gen_RAND_INT(min int, max int) int {

	resultNum := rand.Intn(max-min) + min

	return resultNum

} //end of genRandomRange
