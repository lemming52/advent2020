package dayseven

import (
	"log"
	"reflect"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBag(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		parentName       string
		childName        string
		expectedParent   *Bag
		expectedChildren []*Bag
	}{
		{
			name:       "base",
			input:      "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			parentName: "light red",
			expectedParent: &Bag{
				name: "light red",
				children: map[string]*Bag{
					"bright white": &Bag{
						name: "bright white",
					},
					"muted yellow": &Bag{
						name: "muted yellow",
					},
				},
				childCounts: map[string]int{
					"bright white": 1,
					"muted yellow": 2,
				},
			},
		},
	}
	parent, err := regexp.Compile(BagPattern)
	if err != nil {
		log.Fatal(err)
	}
	count, err := regexp.Compile(NumberedBagPattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			network := BagNetwork{
				bags: map[string]*Bag{},
			}
			ParseBag(tt.input, parent, count, &network)
			if !reflect.DeepEqual(network.bags[tt.parentName].childCounts, tt.expectedParent.childCounts) {
				t.Errorf("Parents should match %v %v", network.bags[tt.parentName], tt.expectedParent)
			}
			assert.Equal(t, len(tt.expectedParent.children), len(network.bags[tt.parentName].children))
		})
	}
}

func TestAddChild(t *testing.T) {
	tests := []struct {
		name           string
		child          *Bag
		parent         *Bag
		expectedParent *Bag
		expectedChild  *Bag
	}{
		{
			name: "base",
			child: &Bag{
				name:     "bright white",
				children: map[string]*Bag{},
				parents:  map[string]*Bag{},
			},
			parent: &Bag{
				name:        "light red",
				children:    map[string]*Bag{},
				parents:     map[string]*Bag{},
				childCounts: map[string]int{},
			},
			expectedChild: &Bag{
				name: "bright white",
				parents: map[string]*Bag{
					"light red": &Bag{
						name: "light red",
					},
				},
			},
			expectedParent: &Bag{
				name: "light red",
				children: map[string]*Bag{
					"bright white": &Bag{
						name: "bright white",
					},
				},
				childCounts: map[string]int{
					"bright white": 1,
				},
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			tt.parent.AddChild(tt.child, "1")
			assert.Equal(t, len(tt.parent.children), len(tt.expectedParent.children))
			assert.Equal(t, len(tt.child.parents), len(tt.child.parents))
		})
	}
}

func TestSeveralBags(t *testing.T) {
	tests := []struct {
		name          string
		input         []string
		targetBag     string
		expectedCount int
		expectedBags  int
	}{
		{
			name: "base",
			input: []string{
				"light red bags contain 1 bright white bag, 2 muted yellow bags.",
				"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
				"bright white bags contain 1 shiny gold bag.",
				"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
				"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
				"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
				"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
				"faded blue bags contain no other bags.",
				"dotted black bags contain no other bags.",
			},
			targetBag:     "shiny gold",
			expectedCount: 4,
			expectedBags:  32,
		}, {
			name: "alt",
			input: []string{
				"shiny gold bags contain 2 dark red bags.",
				"dark red bags contain 2 dark orange bags.",
				"dark orange bags contain 2 dark yellow bags.",
				"dark yellow bags contain 2 dark green bags.",
				"dark green bags contain 2 dark blue bags.",
				"dark blue bags contain 2 dark violet bags.",
				"dark violet bags contain no other bags.",
			},
			targetBag:     "shiny gold",
			expectedCount: 0,
			expectedBags:  126,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
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
			for _, input := range tt.input {
				ParseBag(input, parent, count, &network)
			}
			assert.Equal(t, tt.expectedCount, len(network.bags[tt.targetBag].parents))
			assert.Equal(t, tt.expectedBags, network.bags[tt.targetBag].CountBags())
		})
	}
}
