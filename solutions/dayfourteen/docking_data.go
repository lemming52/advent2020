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

// Memory is a struct that represents the post encoded memory state
type Memory struct {
	mask   string
	memory map[int]uint64
}

// NewMemory instantiates a new type 1 memory
func NewMemory() *Memory {
	return &Memory{
		memory: map[int]uint64{},
	}
}

// UpdateMask sets the value of the memory mask
func (m *Memory) UpdateMask(s string) {
	m.mask = s
}

// AddValue applies the mask to the provided value and stores it
func (m *Memory) AddValue(address, value int) {
	val := float64(value)
	bits := strconv.FormatUint(uint64(value), 2)
	start := 36 - len(bits)
	power := float64(35)
	for i, c := range m.mask {
		switch c {
		case '0':
			val += m.applyMaskBit(-1, power, i, start, bits)
		case '1':
			val += m.applyMaskBit(1, power, i, start, bits)
		}
		power--
	}
	m.memory[address] = uint64(val)
}

func (m *Memory) applyMaskBit(sign int, power float64, position, start int, bits string) float64 {
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

// Total returns the total value of all set memory addresses
func (m *Memory) Total() uint64 {
	total := uint64(0)
	for _, val := range m.memory {
		total += val
	}
	return total
}

// Memory2 is a memory struct that supports multiple address encodings
type Memory2 struct {
	masks           []uint64
	mask            uint64
	floatingIndices []int
	memory          map[uint64]uint64
}

// NewMemory2 instantiates a new type 2 memory with encoding of addresses
func NewMemory2() *Memory2 {
	return &Memory2{
		memory:          map[uint64]uint64{},
		floatingIndices: []int{},
	}
}

// Update mask takes an input mask and computes variations
func (m *Memory2) UpdateMask(s string) {
	mask := uint64(0)
	power := float64(35)
	indices := []int{}
	for i, s := range s {
		if s == 'X' {
			indices = append(indices, 36-i)
		} else if s == '1' {
			mask += uint64(math.Pow(2, power))
		}
		power--
	}
	m.mask = mask
	m.floatingIndices = indices
}

// AddValue applies the mask to the provided addresses and stores the value in each
func (m *Memory2) AddValue(address, value int) {
	addresses := []uint64{uint64(address) | m.mask}

	for _, i := range m.floatingIndices {
		newAddr := []uint64{}
		for _, addr := range addresses {
			newAddr = append(newAddr, addr, addr^(1<<(i-1)))
		}
		addresses = newAddr
	}

	val := uint64(value)
	for _, addr := range addresses {
		m.memory[addr] = val
	}
}

// Total returns the total value of all set memory addresses
func (m *Memory2) Total() uint64 {
	total := uint64(0)
	for _, val := range m.memory {
		total += val
	}
	return total
}

// InitialiseDocking recovers the docking instructions
func InitialiseDocking(path string) (uint64, uint64) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := NewMemory()
	m2 := NewMemory2()
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
		parseLine(scanner.Text(), maskPattern, valPattern, m, m2)
	}
	return m.Total(), m2.Total()
}

type mem interface {
	UpdateMask(string)
	AddValue(int, int)
}

func parseLine(line string, maskPattern, valPattern *regexp.Regexp, m, m2 mem) {
	if strings.HasPrefix(line, "mask = ") {
		match := maskPattern.FindStringSubmatch(line)
		m.UpdateMask(match[1])
		m2.UpdateMask(match[1])
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
		m2.AddValue(address, value)
	}
}
