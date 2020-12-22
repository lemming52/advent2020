package daynineteen

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	CHARACTER  = iota
	SEQUENCE   = iota
	ORSEQUENCE = iota

	RulePattern       = `(\d+): (.+)`
	OrSequencePattern = `\d{1,3}`
)

type Codex struct {
	rules      map[int]*Rule
	orSequence *regexp.Regexp
}

func NewCodex() *Codex {
	return &Codex{
		rules:      map[int]*Rule{},
		orSequence: regexp.MustCompile(OrSequencePattern),
	}
}

func (c *Codex) AddRule(index int, rule string) {
	if strings.Contains(rule, "\"") {
		c.rules[index] = &Rule{
			name:     index,
			ruleType: CHARACTER,
			codex:    c,
			value:    rune(rule[1]),
		}
		return
	}
	if strings.Contains(rule, "|") {
		match := c.orSequence.FindAllStringSubmatch(rule, -1)
		indices := []int{}
		for _, m := range match {
			i, err := strconv.Atoi(m[0])
			if err != nil {
				log.Fatal(err)
			}
			indices = append(indices, i)
		}
		var seq, sec []int
		if len(indices) == 4 {
			seq = []int{indices[0], indices[1]}
			sec = []int{indices[2], indices[3]}
		} else {
			seq = []int{indices[0]}
			sec = []int{indices[1]}
		}
		c.rules[index] = &Rule{
			name:      index,
			ruleType:  ORSEQUENCE,
			codex:     c,
			sequence:  seq,
			secondary: sec,
		}
		return
	}
	components := strings.Split(rule, " ")
	indices := []int{}
	for _, c := range components {
		i, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		indices = append(indices, i)
	}
	c.rules[index] = &Rule{
		name:     index,
		ruleType: SEQUENCE,
		codex:    c,
		sequence: indices,
	}
}

type Rule struct {
	name      int
	ruleType  int
	codex     *Codex
	value     rune
	sequence  []int
	secondary []int
}

func (r *Rule) Apply(s string) (bool, int) {
	switch r.ruleType {
	case CHARACTER:
		return rune(s[0]) == r.value, 1
	case SEQUENCE:
		var increment int
		for _, id := range r.sequence {
			rule := r.codex.rules[id]
			success, del := rule.Apply(s[increment:])
			if !success {
				return false, 0
			}
			increment += del
		}
		return true, increment
	case ORSEQUENCE:
		success := true
		var del, increment int
		for _, id := range r.sequence {
			rule := r.codex.rules[id]
			success, del = rule.Apply(s[increment:])
			if !success {
				break
			}
			increment += del
		}
		if success {
			return true, increment
		}
		increment = 0
		for _, id := range r.secondary {
			rule := r.codex.rules[id]
			success, del = rule.Apply(s[increment:])
			if !success {
				return false, 0
			}
			increment += del
		}
		return true, increment
	}
	return false, 0
}

func LoadMessages(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	pattern := regexp.MustCompile(RulePattern)
	c := NewCodex()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		match := pattern.FindStringSubmatch(line)
		id, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		c.AddRule(id, match[2])
	}

	count := 0
	zeroth := c.rules[0]
	for scanner.Scan() {
		line := scanner.Text()
		success, increment := zeroth.Apply(scanner.Text())
		if success && increment == len(line) {
			count++
		}
	}
	return count
}
