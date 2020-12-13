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
		expectedParent   *Bag
		expectedChildren []*Bag
	}{
		{
			name:  "base",
			input: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			expectedParent: &Bag{
				name: "light red",
			},
			expectedChildren: []*Bag{
				{
					name: "bright white",
				}, {
					name: "muted yellow",
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
			parent, children := ParseBag(tt.input, parent, count, &network)
			if !reflect.DeepEqual(parent, tt.expectedParent) {
				t.Errorf("Parents should match %v %v", parent, tt.expectedParent)
			}
			if !reflect.DeepEqual(children, tt.expectedChildren) {
				t.Errorf("Children should match %v %v", children, tt.expectedChildren)
			}
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
				name:     "light red",
				children: map[string]*Bag{},
				parents:  map[string]*Bag{},
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
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			tt.parent.AddChild(tt.child)
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
				parent, children := ParseBag(input, parent, count, &network)
				for _, child := range children {
					network.AddBag(child)
					parent.AddChild(child)
				}
			}
			assert.Equal(t, tt.expectedCount, len(network.bags[tt.targetBag].parents))
		})
	}
}
