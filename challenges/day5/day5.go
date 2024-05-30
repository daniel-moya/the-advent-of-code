package day5

import "github.com/daniel-moya/aoc/challenges/utils"
import "strings"
import "strconv"

const SeedToSoil = "seed-to-soil"
const SoildToFertilizer = "soil-to-fertilizer"
const FertilizerToWater = "fertilizer-to-water"
const WaterToLight = "water-to-light"
const LightToTemperature = "light-to-temperature"
const TemperatureToHumidity = "temperature-to-humidity"
const HumidityToLocation = "humidity-to-location"

var MapChain = [...]string{SeedToSoil, SoildToFertilizer, FertilizerToWater, WaterToLight, LightToTemperature, TemperatureToHumidity, HumidityToLocation}

type Data struct {
	Target int
	Source int
	Offset int
}

func Run(inputPath string) int {
	slice := utils.GetInputFromFile(inputPath)
	var maps map[string][]Data = make(map[string][]Data)
	maps[SeedToSoil] = []Data{}
	maps[SoildToFertilizer] = []Data{}
	maps[FertilizerToWater] = []Data{}
	maps[WaterToLight] = []Data{}
	maps[LightToTemperature] = []Data{}
	maps[TemperatureToHumidity] = []Data{}
	maps[HumidityToLocation] = []Data{}

	var seeds []int
	var mapKey string
	for row, line := range slice {
		if line == "" {
			continue
		}

		if seeds == nil && row == 0 {
			for _, seedStr := range strings.Fields(strings.Split(line, ":")[1]) {
				seed, _ := strconv.Atoi(seedStr)
				seeds = append(seeds, seed)
			}
			continue
		}

		lineSlice := strings.Fields(line)

		if lineSlice[1] == "map:" {
			mapKey = lineSlice[0]
			continue
		}
		targetPoint, _ := strconv.Atoi(lineSlice[0])
		sourcePoint, _ := strconv.Atoi(lineSlice[1])
		rangeSize, _ := strconv.Atoi(lineSlice[2])

		maps[mapKey] = append(maps[mapKey], Data{Source: sourcePoint, Target: targetPoint, Offset: rangeSize})
	}

	result := 0
	for _, seed := range seeds {
		currSource := seed

		for _, currMap := range MapChain {
			target := getTarget(currMap, currSource, maps)
			currSource = target
		}

		if result == 0 {
			result = currSource
			continue
		}

		if currSource < result {
			result = currSource
		}
	}

	return result
}

func getTarget(mapKey string, source int, maps map[string][]Data) int {
	var dataSlice []Data = maps[mapKey]
	target := source
	for _, data := range dataSlice {
		if (data.Source+data.Offset) >= source && source >= data.Source {
			target = source + (data.Target - data.Source)
		}
	}

	return target
}
