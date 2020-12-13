package dayseven

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

const BagPattern = `([a-z]+ [a-z]+) bag`
const NumberedBagPattern = `([0-9]) ([a-z]+ [a-z]+) bag`
const TargetBag = "shiny gold"

// BagNetwork is the full set of all bags
type BagNetwork struct {
	bags map[string]*Bag
}

func (n *BagNetwork) AddBag(bag *Bag) {
	_, ok := n.bags[bag.name]
	if !ok {
		n.bags[bag.name] = bag
	}
}

func (n *BagNetwork) GetBag(name string) *Bag {
	bag, ok := n.bags[name]
	if ok {
		return bag
	}
	n.bags[name] = &Bag{
		name:        name,
		children:    map[string]*Bag{},
		parents:     map[string]*Bag{},
		childCounts: map[string]int{},
	}
	return n.bags[name]
}

// Bag is a struct that represents any single bag and possible containers or containees
type Bag struct {
	name        string
	parents     map[string]*Bag
	children    map[string]*Bag
	childCounts map[string]int
}

// AddParent adds the name of a parent ot a list of possible parents of a bag
// and propagates that name to all of it's children.
func (b *Bag) AddParent(parent *Bag) {
	_, ok := b.parents[parent.name]
	if !ok {
		b.parents[parent.name] = parent
	}
	for _, child := range b.children {
		child.AddParent(parent)
	}
}

// AddChild adds a node as a child to the current bag,
// and propagates the current bag as a parent down the children
func (b *Bag) AddChild(child *Bag, number string) {
	count, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	for _, bag := range b.parents {
		child.AddParent(bag)
	}
	b.children[child.name] = child
	b.childCounts[child.name] = count
	child.AddParent(b)
}

// CountBags recursively counts the bags down.
func (b *Bag) CountBags() int {
	count := 0
	for name, bag := range b.children {
		count = count + b.childCounts[name]
		count = count + (b.childCounts[name] * bag.CountBags())
	}
	return count
}

// ParseBag takes a line of the input file and recovers the parent bag and any children
func ParseBag(s string, parentPattern, childPattern *regexp.Regexp, network *BagNetwork) {
	parentComponents := parentPattern.FindStringSubmatch(s)
	parent := network.GetBag(parentComponents[1])
	children := childPattern.FindAllStringSubmatch(s, -1)
	for _, child := range children {
		parent.AddChild(network.GetBag(child[2]), child[1])
	}
}

func LoadBags(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	parent, err := regexp.Compile(BagPattern)
	if err != nil {
		log.Fatal(err)
	}
	count, err := regexp.Compile(NumberedBagPattern)
	if err != nil {
		log.Fatal(err)
	}
	network := BagNetwork{
		bags: map[string]*Bag{},
	}
	for scanner.Scan() {
		line := scanner.Text()
		ParseBag(line, parent, count, &network)
	}
	return len(network.bags[TargetBag].parents), network.bags[TargetBag].CountBags()
}
