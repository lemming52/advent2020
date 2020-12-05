package main

import (
	"advent/solutions/dayone"
	"advent/solutions/daytwo"
	"flag"
	"fmt"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "dayone", "challenge in <day><number> format, with number as the word")
	flag.Parse()

	var res string
	input := fmt.Sprintf("inputs/%s.txt", challenge)
	switch challenge {
	case "dayone":
		A, B := dayone.LoadReport(input)
		res = fmt.Sprintf("Day One Results A: %d B: %d", A, B)
	case "daytwo":
		A := daytwo.ValidatePasswords(input)
		res = fmt.Sprintf("Day One Results A: %d", A)
	}
	fmt.Println(res)
}
