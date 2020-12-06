package main

import (
	"advent/solutions/dayfive"
	"advent/solutions/dayfour"
	"advent/solutions/dayone"
	"advent/solutions/daythree"
	"advent/solutions/daytwo"
	"flag"
	"fmt"
)

// res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
// res = fmt.Sprintf("%s Results A: %d", challenge, A)

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
		A, B := dayfour.LoadPassports(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayfive":
		A, B := dayfive.LoadBoardingPasses(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	}
	fmt.Println(res)
}
