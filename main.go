package main

import (
	"advent/solutions/dayone"
	"flag"
	"fmt"
)

func main() {
	var challenge string
	flag.StringVar(&challenge, "challenge", "dayone", "challenge in <day><number> format, with number as the word")
	flag.Parse()

	var res string
	switch challenge {
	case "dayone":
		A, B := dayone.LoadReport("inputs/dayone.txt")
		res = fmt.Sprintf("Day One Results A: %d B: %d", A, B)
	}
	fmt.Println(res)
}
