package daytwelve

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	FORWARD = 'F'
	LEFT    = 'L'
	RIGHT   = 'R'
	NORTH   = 'N'
	SOUTH   = 'S'
	EAST    = 'E'
	WEST    = 'W'
)

// Ship manages the position and heading of a ship
type Ship struct {
	x              int
	y              int
	directions     []rune
	directionIndex int
}

// NewShip instantiates a ship for the default conditions
func NewShip() *Ship {
	return &Ship{
		x:              0,
		y:              0,
		directions:     []rune("NESW"),
		directionIndex: 1,
	}
}

// Instruct executes a given navigation instruction on the ship
func (s *Ship) Instruct(instruction string) {
	ins := rune(instruction[0])
	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatal(err)
	}
	switch ins {
	case FORWARD:
		s.MoveCardinal(s.directions[s.directionIndex], distance)
	case LEFT:
		s.Rotate(-1, distance)
	case RIGHT:
		s.Rotate(1, distance)
	default:
		s.MoveCardinal(ins, distance)
	}
}

// MoveCardinal moves the ship in a given cardinal direction by the distance
func (s *Ship) MoveCardinal(heading rune, distance int) {
	switch heading {
	case NORTH:
		s.y += distance
	case SOUTH:
		s.y -= distance
	case EAST:
		s.x += distance
	case WEST:
		s.x -= distance
	}
}

// Rotate adjusts the ships heading between the cardinal directions
func (s *Ship) Rotate(direction, rotation int) {
	shift := (rotation * direction) / 90
	s.directionIndex = (s.directionIndex + shift + 4) % 4
}

// Manhattan returns the manhattan distance of the ship from the origin
func (s *Ship) Manhattan() int {
	return abs(s.x) + abs(s.y)
}

func (s *Ship) print() string {
	return fmt.Sprintf("%d %d %d %s %d", s.x, s.y, s.directionIndex, string(s.directions[s.directionIndex]), s.Manhattan())
}

// WaypointShip manages the position and vector of a ship that navigates by waypoint
type WaypointShip struct {
	x         int
	y         int
	waypointX int
	waypointY int
}

// NewWaypointShip instantiates a waypoint ship for the default conditions
func NewWaypointShip() *WaypointShip {
	return &WaypointShip{
		x:         0,
		y:         0,
		waypointX: 10,
		waypointY: 1,
	}
}

// Instruct executes a given navigation instruction on the ship
func (s *WaypointShip) Instruct(instruction string) {
	ins := rune(instruction[0])
	distance, err := strconv.Atoi(instruction[1:])
	if err != nil {
		log.Fatal(err)
	}
	switch ins {
	case FORWARD:
		s.FollowWaypoint(distance)
	case LEFT:
		s.Rotate(-1, distance)
	case RIGHT:
		s.Rotate(1, distance)
	default:
		s.MoveWaypoint(ins, distance)
	}
}

// FollowWaypoint moves the ship in the direction of the waypoint multiple times
func (s *WaypointShip) FollowWaypoint(multiple int) {
	s.x += s.waypointX * multiple
	s.y += s.waypointY * multiple
}

// MoveWaypoiny moves the waypoint in a given cardinal direction by the distance
func (s *WaypointShip) MoveWaypoint(heading rune, distance int) {
	switch heading {
	case NORTH:
		s.waypointY += distance
	case SOUTH:
		s.waypointY -= distance
	case EAST:
		s.waypointX += distance
	case WEST:
		s.waypointX -= distance
	}
}

// Rotate adjusts the ships heading around the ship
func (s *WaypointShip) Rotate(direction, rotation int) {
	shift := (rotation * direction) / 90
	rot := (shift + 4) % 4
	switch rot {
	case 1:
		s.waypointX, s.waypointY = s.waypointY, -s.waypointX
	case 2:
		s.waypointX, s.waypointY = -s.waypointX, -s.waypointY
	case 3:
		s.waypointX, s.waypointY = -s.waypointY, s.waypointX
	}
}

// Manhattan returns the manhattan distance of the ship from the origin
func (s *WaypointShip) Manhattan() int {
	return abs(s.x) + abs(s.y)
}

// abs is a convenience function for absolute int values
func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Navigate(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ship := NewShip()
	waypoint := NewWaypointShip()
	for scanner.Scan() {
		line := scanner.Text()
		ship.Instruct(line)
		waypoint.Instruct(line)
	}
	return ship.Manhattan(), waypoint.Manhattan()
}
