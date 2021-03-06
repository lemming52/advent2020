package main

import (
	"advent/solutions/dayeight"
	"advent/solutions/dayeighteen"
	"advent/solutions/dayeleven"
	"advent/solutions/dayfifteen"
	"advent/solutions/dayfive"
	"advent/solutions/dayfour"
	"advent/solutions/dayfourteen"
	"advent/solutions/daynine"
	"advent/solutions/daynineteen"
	"advent/solutions/dayone"
	"advent/solutions/dayseven"
	"advent/solutions/dayseventeen"
	"advent/solutions/daysix"
	"advent/solutions/daysixteen"
	"advent/solutions/dayten"
	"advent/solutions/daythirteen"
	"advent/solutions/daythree"
	"advent/solutions/daytwelve"
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
		"dayeleven",
		"daytwelve",
		"daythirteen",
		"dayfourteen",
		"dayfifteen",
		"daysixteen",
		"dayseventeen",
		"dayeighteen",
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
		A, B := dayeleven.LoadSeats(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daytwelve":
		A, B := daytwelve.Navigate(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daythirteen":
		A, B := daythirteen.LoadShuttles(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayfourteen":
		A, B := dayfourteen.InitialiseDocking(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayfifteen":
		A, B := dayfifteen.Play(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daysixteen":
		A, B := daysixteen.LoadTickets(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayseventeen":
		A, B := dayseventeen.LoadCube(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "dayeighteen":
		A, B := dayeighteen.LoadOperations(input)
		res = fmt.Sprintf("%s Results A: %d B: %d", challenge, A, B)
	case "daynineteen":
		A := daynineteen.LoadMessages(input)
		res = fmt.Sprintf("%s Results A: %d", challenge, A)
	}
	return res
}
