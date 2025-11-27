// Rewritten with help... original code for part 2 was inefficient and took over 2mins to run.
package day05

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Data struct {
	Start, End, Destination int
}

func PartOne() int {
	// INPUT = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"
	parts := strings.Split(INPUT, "\n\n")
	numbers := strings.TrimPrefix(parts[0], "seeds: ")

	var seeds []*Data
	for _, v := range strings.Fields(numbers) {
		seeds = append(seeds, &Data{Start: convertToNum(v), End: convertToNum(v)})
	}

	return run(parts[1:], seeds)
}

func PartTwo() int {
	// INPUT = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"
	parts := strings.Split(INPUT, "\n\n")
	numbers := strings.TrimPrefix(parts[0], "seeds: ")

	var values []int
	for _, v := range strings.Fields(numbers) {
		values = append(values, convertToNum(v))
	}

	var seeds []*Data
	for i := 0; i < len(values); i += 2 {
		seeds = append(seeds, &Data{Start: values[i], End: values[i] + values[i+1] - 1})
	}

	return run(parts[1:], seeds)
}

func run(parts []string, seeds []*Data) int {
	for _, part := range parts {
		lines := strings.Split(part, "\n")[1:]

		var nextInputs []*Data
		for _, line := range lines {
			var destination, source, rangeLength int
			fmt.Sscanf(line, "%d %d %d", &destination, &source, &rangeLength)

			nextInputs = append(nextInputs, &Data{Start: source, End: source + rangeLength - 1, Destination: destination})
		}

		var newSeeds []*Data
		for _, seed := range seeds {
			newSeeds = append(newSeeds, filterByNext(nextInputs, seed)...)
		}

		seeds = newSeeds
	}

	var res = math.MaxInt64
	for _, seed := range seeds {
		res = min(res, seed.Start)
	}

	return res
}

func filterByNext(nextInputs []*Data, ab *Data) []*Data {
	var res []*Data
	var seeds = []*Data{ab}
	for _, nextInput := range nextInputs {
		var transformed, others []*Data
		for _, seed := range seeds {
			if seed.End < nextInput.Start || seed.Start > nextInput.End {
				others = append(others, seed)
			} else {
				t, o := transformer(seed, nextInput)
				transformed = append(transformed, t)
				others = append(others, o...)
			}
		}

		res = append(res, transformed...)

		seeds = others
		if len(seeds) == 0 {
			return res
		}
	}

	res = append(res, seeds...)

	return res
}

func transformer(seed, nextInput *Data) (transformed *Data, other []*Data) {
	delta := nextInput.Destination - nextInput.Start

	if seed.End < nextInput.Start || seed.Start > nextInput.End {
		return &Data{Start: 0, End: -1}, []*Data{seed}
	}

	if seed.Start >= nextInput.Start && seed.End <= nextInput.End {
		return &Data{Start: seed.Start + delta, End: seed.End + delta}, other
	}

	if seed.End >= nextInput.Start && seed.End < nextInput.End {
		if seed.Start < nextInput.Start {
			other = append(other, &Data{Start: seed.Start, End: nextInput.Start - 1})
		}
		return &Data{Start: nextInput.Start + delta, End: seed.End + delta}, other
	}

	if seed.Start > nextInput.Start && seed.Start <= nextInput.End {
		if seed.End > nextInput.End {
			other = append(other, &Data{Start: nextInput.End + 1, End: seed.End})
		}
		return &Data{Start: seed.Start + delta, End: nextInput.End + delta}, other
	}

	if seed.Start <= nextInput.Start && nextInput.End <= seed.End {
		if seed.Start < nextInput.Start {
			other = append(other, &Data{Start: seed.Start, End: nextInput.Start - 1})
		}
		if nextInput.End < seed.End {
			other = append(other, &Data{Start: nextInput.End + 1, End: seed.End})
		}
		return &Data{Start: nextInput.Start + delta, End: nextInput.End + delta}, other
	}

	return nil, nil
}

func convertToNum(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return result
}
