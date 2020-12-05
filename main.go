package main

import (
	"advent/solutions/dayfour"
	"advent/solutions/dayone"
	"advent/solutions/daythree"
	"advent/solutions/daytwo"
	"flag"
	"fmt"
)

// res = fmt.Sprintf("Day One Results A: %d B: %d", A, B)
// res = fmt.Sprintf("Day One Results A: %d", A)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "dayone", "challenge in <day><number> format, with number as the word")
	flag.Parse()

	var res string
	input := fmt.Sprintf("inputs/%s.txt", challenge)
	switch challenge {
	case "dayone":
		A, B := dayone.LoadReport(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daytwo":
		A, B := daytwo.ValidatePasswords(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daythree":
		A, B := daythree.LoadTrees(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayfour":
		A := dayfour.LoadPassports(input)
		res = fmt.Sprintf("Day One Results A: %d", A)
	}
	fmt.Println(res)
}
