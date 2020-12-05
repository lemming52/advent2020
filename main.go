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

	switch challenge {
	case "dayone":
		fmt.Println(dayone.LoadReport("inputs/dayone.txt"))
	}
}
