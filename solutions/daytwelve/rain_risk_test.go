package daytwelve

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name          string
		direction     int
		rotation      int
		initialIndex  int
		expectedIndex int
	}{
		{
			name:          "R90",
			direction:     1,
			rotation:      90,
			initialIndex:  3,
			expectedIndex: 0,
		}, {
			name:          "L180",
			direction:     -1,
			rotation:      180,
			initialIndex:  1,
			expectedIndex: 3,
		}, {
			name:          "L270",
			direction:     -1,
			rotation:      270,
			initialIndex:  1,
			expectedIndex: 2,
		}, {
			name:          "L360",
			direction:     -1,
			rotation:      360,
			initialIndex:  0,
			expectedIndex: 0,
		}, {
			name:          "R360",
			direction:     1,
			rotation:      360,
			initialIndex:  3,
			expectedIndex: 3,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewShip()
			ship.directionIndex = tt.initialIndex
			ship.Rotate(tt.direction, tt.rotation)
			assert.Equal(t, tt.expectedIndex, ship.directionIndex)
		})
	}
}

func TestInstruct(t *testing.T) {
	tests := []struct {
		name          string
		instruction   string
		expectedX     int
		expectedY     int
		expectedIndex int
	}{
		{
			name:          "forward",
			instruction:   "F15",
			expectedX:     15,
			expectedY:     0,
			expectedIndex: 1,
		}, {
			name:          "north",
			instruction:   "N15",
			expectedX:     0,
			expectedY:     15,
			expectedIndex: 1,
		}, {
			name:          "east",
			instruction:   "E15",
			expectedX:     15,
			expectedY:     0,
			expectedIndex: 1,
		}, {
			name:          "left",
			instruction:   "L90",
			expectedX:     0,
			expectedY:     0,
			expectedIndex: 0,
		}, {
			name:          "right",
			instruction:   "R90",
			expectedX:     0,
			expectedY:     0,
			expectedIndex: 2,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewShip()
			ship.Instruct(tt.instruction)
			assert.Equal(t, tt.expectedX, ship.x)
			assert.Equal(t, tt.expectedY, ship.y)
			assert.Equal(t, tt.expectedIndex, ship.directionIndex)
		})
	}
}

func TestNavigate(t *testing.T) {
	tests := []struct {
		name             string
		input            []string
		expectedX        int
		expectedY        int
		expectedIndex    int
		expectedDistance int
	}{
		{
			name: "base",
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expectedX:        17,
			expectedY:        -8,
			expectedIndex:    2,
			expectedDistance: 25,
		}, {
			name: "sample",
			input: []string{
				"R90",
				"F56",
				"R90",
				"F56",
				"R90",
				"R180",
				"W5",
				"L90",
				"E2",
				"L90",
				"S5",
				"E1",
				"F11",
				"L90",
				"F46",
				"S2",
				"E2",
				"S1",
				"E2",
				"E3",
				"N4",
			},
			expectedX:        -97,
			expectedY:        -49,
			expectedIndex:    3,
			expectedDistance: 146,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewShip()
			for _, ins := range tt.input {
				ship.Instruct(ins)
			}
			assert.Equal(t, tt.expectedX, ship.x)
			assert.Equal(t, tt.expectedY, ship.y)
			assert.Equal(t, tt.expectedIndex, ship.directionIndex)
			assert.Equal(t, tt.expectedDistance, ship.Manhattan())
		})
	}
}

func TestWaypointRotate(t *testing.T) {
	tests := []struct {
		name      string
		direction int
		rotation  int
		expectedX int
		expectedY int
	}{
		{
			name:      "R90",
			direction: 1,
			rotation:  90,
			expectedX: 1,
			expectedY: -10,
		}, {
			name:      "L180",
			direction: -1,
			rotation:  180,
			expectedX: -10,
			expectedY: -1,
		}, {
			name:      "L270",
			direction: -1,
			rotation:  270,
			expectedX: 1,
			expectedY: -10,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewWaypointShip()
			ship.Rotate(tt.direction, tt.rotation)
			assert.Equal(t, tt.expectedX, ship.waypointX)
			assert.Equal(t, tt.expectedY, ship.waypointY)
		})
	}
}

func TestWaypointInstruct(t *testing.T) {
	tests := []struct {
		name              string
		instruction       string
		expectedX         int
		expectedY         int
		expectedWaypointX int
		expectedWaypointY int
	}{
		{
			name:              "forward",
			instruction:       "F15",
			expectedX:         150,
			expectedY:         15,
			expectedWaypointX: 10,
			expectedWaypointY: 1,
		}, {
			name:              "north",
			instruction:       "N15",
			expectedX:         0,
			expectedY:         0,
			expectedWaypointX: 10,
			expectedWaypointY: 16,
		}, {
			name:              "east",
			instruction:       "E15",
			expectedX:         0,
			expectedY:         0,
			expectedWaypointX: 25,
			expectedWaypointY: 1,
		}, {
			name:              "west",
			instruction:       "W25",
			expectedX:         0,
			expectedY:         0,
			expectedWaypointX: -15,
			expectedWaypointY: 1,
		}, {
			name:              "left",
			instruction:       "L90",
			expectedX:         0,
			expectedY:         0,
			expectedWaypointX: -1,
			expectedWaypointY: 10,
		}, {
			name:              "right",
			instruction:       "R90",
			expectedX:         0,
			expectedY:         0,
			expectedWaypointX: 1,
			expectedWaypointY: -10,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewWaypointShip()
			ship.Instruct(tt.instruction)
			assert.Equal(t, tt.expectedX, ship.x)
			assert.Equal(t, tt.expectedY, ship.y)
			assert.Equal(t, tt.expectedWaypointX, ship.waypointX)
			assert.Equal(t, tt.expectedWaypointY, ship.waypointY)
		})
	}
}

func TestWaypointNavigate(t *testing.T) {
	tests := []struct {
		name              string
		input             []string
		expectedX         int
		expectedY         int
		expectedWaypointX int
		expectedWaypointY int
		expectedDistance  int
	}{
		{
			name: "base",
			input: []string{
				"F10",
				"N3",
				"F7",
				"R90",
				"F11",
			},
			expectedX:         214,
			expectedY:         -72,
			expectedWaypointX: 4,
			expectedWaypointY: -10,
			expectedDistance:  286,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			ship := NewWaypointShip()
			for _, ins := range tt.input {
				ship.Instruct(ins)
			}
			assert.Equal(t, tt.expectedX, ship.x)
			assert.Equal(t, tt.expectedWaypointX, ship.waypointX)
			assert.Equal(t, tt.expectedWaypointY, ship.waypointY)
			assert.Equal(t, tt.expectedDistance, ship.Manhattan())
		})
	}
}
