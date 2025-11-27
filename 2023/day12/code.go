package day12

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	OPERATIONAL = '.'
	DAMAGED     = '#'
	UNKNOWN     = '?'
)

var cache = map[string]int{}

func PartOne() int {
	// INPUT = "???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1" // 21
	data := parse(INPUT)
	total := 0

	for _, d := range data {
		a := d.Arrangements()
		total += a
	}

	return total
}

func PartTwo() int {
	// INPUT = "???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1" // 525152@5
	data := parse(INPUT)
	total := 0

	for _, d := range data {
		d.Expand(5)
		a := d.Arrangements()
		total += a
	}

	return total
}

type Data struct {
	condition string
	groups    []int
}

func (d *Data) key() string {
	return fmt.Sprintf("%s-%v", d.condition, d.groups)
}

func (d *Data) iters() int {
	result := len(d.groups) - 1

	for _, val := range d.groups {
		result += val
	}

	return len(d.condition) - result
}

func (d *Data) Expand(times int) {
	newCond := ""
	newGroups := []int{}

	for i := 0; i < times; i++ {
		newCond += d.condition + "?"
		newGroups = append(newGroups, d.groups...)
	}

	d.condition = strings.TrimSuffix(newCond, "?")
	d.groups = newGroups
}

func (d *Data) Arrangements() int {
	if val, exist := cache[d.key()]; exist {
		return val
	} else if len(d.groups) == 0 && !strings.Contains(d.condition, string(DAMAGED)) {
		cache[d.key()] = 1
		return 1
	} else if len(d.groups) == 0 {
		cache[d.key()] = 0
		return 0
	}

	total := 0
	g := d.groups[0]

	for i := 0; i <= d.iters(); i++ {
		if strings.Contains(d.condition[0:i], string(DAMAGED)) || strings.Contains(d.condition[i:i+g], string(OPERATIONAL)) {
			continue
		}
		if len(d.condition) > i+g && d.condition[i+g] == DAMAGED {
			continue
		}
		if g+i < len(d.condition)-1 {
			next := Data{condition: d.condition[g+i+1:], groups: d.groups[1:]}
			total += next.Arrangements()
			continue
		}

		total += 1
	}

	cache[d.key()] = total
	return total
}

func parse(input string) []Data {
	data := []Data{}

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")

		groupInts := []int{}
		for _, g := range strings.Split(parts[1], ",") {
			gNum, err := strconv.Atoi(g)
			if err != nil {
				panic(err)
			}
			groupInts = append(groupInts, gNum)
		}

		data = append(data, Data{condition: parts[0], groups: groupInts})
	}

	return data
}
