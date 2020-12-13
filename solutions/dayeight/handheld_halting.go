package dayeight

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Processor struct {
	acc      int
	index    int
	executed map[int]bool
}

// ExecuteInstructions evaluates a full set of instructions on the processor
func (p *Processor) ExecuteInstructions(instructions []string, alternativeIndex int, altIns string) (int, bool) {
	execute := true
	ins := instructions[0]
	length := len(instructions)
	p.executed[0] = true
	for execute {
		next := p.Execute(ins)
		_, ok := p.executed[next]
		if ok {
			return p.acc, false
		}
		p.executed[next] = true
		if next >= length {
			break
		} else if next == alternativeIndex {
			ins = altIns
		} else {
			ins = instructions[next]
		}
	}
	return p.acc, true
}

// ExecuteInstructionsSuccess will correct the instructions to run successfully
func (p *Processor) ExecuteInstructionsSuccess(instructions []string) int {
	alternativeIndex := 0
	length := len(instructions)
	for alternativeIndex < length {
		ins := instructions[alternativeIndex]
		altIns, ok := alterInstruction(ins)
		if !ok {
			alternativeIndex++
			continue
		}
		p.executed = map[int]bool{}
		p.acc = 0
		p.index = 0
		acc, success := p.ExecuteInstructions(instructions, alternativeIndex, altIns)
		if success {
			return acc
		}
		alternativeIndex++
	}
	return p.acc
}

// Execute evaluates any individual instruction and updates the processor state
func (p *Processor) Execute(instruction string) int {
	components := strings.Split(instruction, " ")
	switch components[0] {
	case "nop":
		p.index = p.index + 1
		return p.index
	case "acc":
		value, err := strconv.Atoi(components[1][1:])
		if err != nil {
			log.Fatal(err)
		}
		if components[1][0] == '+' {
			p.acc = p.acc + value
		} else {
			p.acc = p.acc - value
		}
		p.index = p.index + 1
		return p.index
	case "jmp":
		value, err := strconv.Atoi(components[1][1:])
		if err != nil {
			log.Fatal(err)
		}
		if components[1][0] == '+' {
			p.index = p.index + value
		} else {
			p.index = p.index - value
		}
		return p.index
	default:
		p.index = p.index + 1
		return p.index
	}
}

func alterInstruction(s string) (string, bool) {
	components := strings.Split(s, " ")
	switch components[0] {
	case "nop":
		return fmt.Sprintf("jmp %s", components[1]), true
	case "jmp":
		return fmt.Sprintf("nop %s", components[1]), true
	default:
		return "", false
	}
}

func LoadInstructions(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	processor := Processor{
		executed: map[int]bool{},
	}
	accIni, _ := processor.ExecuteInstructions(instructions, -1, "")
	accSucceed := processor.ExecuteInstructionsSuccess(instructions)
	return accIni, accSucceed
}
