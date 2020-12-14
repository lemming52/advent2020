package main

import (
	"advent/solutions/dayeight"
	"advent/solutions/dayeleven"
	"advent/solutions/dayfive"
	"advent/solutions/dayfour"
	"advent/solutions/daynine"
	"advent/solutions/dayone"
	"advent/solutions/dayseven"
	"advent/solutions/daysix"
	"advent/solutions/dayten"
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
	all := flag.Bool("all", false, "display all results")
	flag.Parse()

	completed := []string{
		"dayone",
		"daytwo",
		"daythree",
		"dayfour",
		"dayfive",
		"daysix",
		"dayseven",
		"dayeight",
		"daynine",
		"dayten",
	}
	if *all {
		for _, c := range completed {
			fmt.Println(RunChallenge(c))
		}
	} else {
		fmt.Println(RunChallenge(challenge))
	}

}

func RunChallenge(challenge string) string {
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
	case "daysix":
		A, B := daysix.LoadQuestions(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayseven":
		A, B := dayseven.LoadBags(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayeight":
		A, B := dayeight.LoadInstructions(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daynine":
		A, B := daynine.LoadNumbers(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayten":
		A, B := dayten.LoadAdapters(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayeleven":
		A := dayeleven.LoadSeats(input)
		res = fmt.Sprintf("%s Results A: %d", challenge, A)
	}
	return res
}
