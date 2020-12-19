package dayfourteen

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Memory struct {
	mask    string
	minMask uint64
	maxMask uint64
	memory  map[int]uint64
}

func NewMemory() *Memory {
	return &Memory{
		memory:  map[int]uint64{},
		minMask: 0,
		maxMask: 0,
	}
}

func (m *Memory) UpdateMask(s string) {
	m.mask = s
}

func (m *Memory) AddValue(address, value int) {
	val := float64(value)
	bits := strconv.FormatUint(uint64(value), 2)
	start := 36 - len(bits)
	power := float64(35)
	for i, c := range m.mask {
		switch c {
		case '0':
			val += m.applyMask(-1, power, i, start, bits)
		case '1':
			val += m.applyMask(1, power, i, start, bits)
		}
		power--
	}
	m.memory[address] = uint64(val)
}

func (m *Memory) applyMask(sign int, power float64, position, start int, bits string) float64 {
	if position < start {
		if sign == 1 {
			return math.Pow(2, power)
		}
		return 0
	} else {
		if m.mask[position] != bits[position-start] {
			return float64(sign) * math.Pow(2, power)
		}
		return 0
	}
}

func (m *Memory) Total() uint64 {
	total := uint64(0)
	for _, val := range m.memory {
		total += val
	}
	return total
}

func InitialiseDocking(path string) uint64 {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := NewMemory()
	maskPattern, err := regexp.Compile(`mask = ([0,1,X]{36})`)
	if err != nil {
		log.Fatal(err)
	}
	valPattern, err := regexp.Compile(`mem\[(\d+)\] = (\d+)`)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(scanner.Text(), maskPattern, valPattern, m)
	}
	return m.Total()
}

func parseLine(line string, maskPattern, valPattern *regexp.Regexp, m *Memory) {
	if strings.HasPrefix(line, "mask = ") {
		match := maskPattern.FindStringSubmatch(line)
		m.UpdateMask(match[1])
	} else {
		match := valPattern.FindStringSubmatch(line)
		address, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		value, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		m.AddValue(address, value)
	}
}
