package main

import "fmt"

func main() {

	input := "01MAR2034"
	res := make([]string, 3)
	res[0] = input[:2]
	res[1] = input[2:5]
	res[2] = input[5:]
	replaceMon(&res[1])
	fmt.Printf("%v\n", res)
}

/*
The function make replacement with NO new string allocation
*/
func replaceMon(s *string) {
	switch *s {
	case "JAN":
		*s = "01"
	case "FEB":
		*s = "02"
	case "MAR":
		*s = "03"
	case "APR":
		*s = "04"
	case "MAY":
		*s = "05"
	case "JUN":
		*s = "06"
	case "JUL":
		*s = "07"
	case "AUG":
		*s = "08"
	case "SEP":
		*s = "09"
	case "OCT":
		*s = "10"
	case "NOV":
		*s = "11"
	case "DEC":
		*s = "12"
	default:
		*s = "??"
	}
}
