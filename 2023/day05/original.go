// Works but part 2 takes almost 3 minutes (2m52.662119708s).
package day05

// import (
// 	"math"
// 	"strconv"
// 	"strings"
// )

// var (
// 	seedToSoilHash            []*Data
// 	soilToFertilizerHash      []*Data
// 	fertilizerToWaterHash     []*Data
// 	waterToLightHash          []*Data
// 	lightToTemperatureHash    []*Data
// 	temperatureToHumidityHash []*Data
// 	humidityToLocationHash    []*Data
// 	hashesForDays             [][]*Data
// )

// type Data struct {
// 	Destination int
// 	Source      int
// 	Range       int
// }

// func PartOne() int {
// 	// INPUT = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"
// 	seeds := parseInputs(INPUT)
// 	closestLocation := math.MaxInt

// 	for _, v := range seeds {
// 		for _, hashMap := range hashesForDays {
// 			v = lookup(v, hashMap)
// 		}

// 		if v < closestLocation {
// 			closestLocation = v
// 		}
// 	}

// 	return closestLocation
// }

// func PartTwo() int {
// 	// INPUT = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4"
// 	seeds := parseInputs(INPUT)

// 	var seedPairs [][]int
// 	for i := 0; i < len(seeds); i += 2 {
// 		seedPairs = append(seedPairs, []int{seeds[i], seeds[i+1]})
// 	}

// 	closestLocation := math.MaxInt

// 	for _, seedSet := range seedPairs {
// 		start := seedSet[0]
// 		length := seedSet[1]

// 		for i := 0; i < length; i++ {
// 			result := start + i
// 			for _, hashMap := range hashesForDays {
// 				result = lookup(result, hashMap)
// 			}

// 			if result < closestLocation {
// 				closestLocation = result
// 			}
// 		}
// 	}

// 	return closestLocation
// }

// func parseInputs(input string) []int {
// 	parts := strings.Split(input, "\n\n")
// 	seedStrings := strings.TrimPrefix(parts[0], "seeds: ")

// 	var seeds []int
// 	for _, v := range strings.Fields(seedStrings) {
// 		seeds = append(seeds, stringToNum(v))
// 	}

// 	seedToSoilHash = convertToData(strings.Split(parts[1], "\n")[1:])
// 	soilToFertilizerHash = convertToData(strings.Split(parts[2], "\n")[1:])
// 	fertilizerToWaterHash = convertToData(strings.Split(parts[3], "\n")[1:])
// 	waterToLightHash = convertToData(strings.Split(parts[4], "\n")[1:])
// 	lightToTemperatureHash = convertToData(strings.Split(parts[5], "\n")[1:])
// 	temperatureToHumidityHash = convertToData(strings.Split(parts[6], "\n")[1:])
// 	humidityToLocationHash = convertToData(strings.Split(parts[7], "\n")[1:])
// 	hashesForDays = [][]*Data{seedToSoilHash, soilToFertilizerHash, fertilizerToWaterHash, waterToLightHash, lightToTemperatureHash, temperatureToHumidityHash, humidityToLocationHash}

// 	return seeds
// }

// func convertToData(lines []string) []*Data {
// 	result := []*Data{}

// 	for _, line := range lines {
// 		parts := strings.Split(line, " ")

// 		source := stringToNum(parts[1])
// 		destination := stringToNum(parts[0])
// 		rangeLen := stringToNum(parts[2])

// 		d := &Data{
// 			Source:      source,
// 			Destination: destination,
// 			Range:       rangeLen,
// 		}

// 		result = append(result, d)
// 	}
// 	return result
// }

// func lookup(source int, hash []*Data) int {
// 	for _, data := range hash {
// 		if source >= data.Source && source < data.Source+data.Range {
// 			return data.Destination + (source - data.Source)
// 		}
// 	}

// 	return source
// }

// func stringToNum(input string) int {
// 	result, err := strconv.Atoi(input)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return result
// }
