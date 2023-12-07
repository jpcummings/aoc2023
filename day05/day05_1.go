package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("Starting...")

	scanner := bufio.NewScanner(os.Stdin)

	var seeds []int
	var s2smaps [][]int
	var s2fmaps [][]int
	var f2wmaps [][]int
	var w2lmaps [][]int
	var l2tmaps [][]int
	var t2hmaps [][]int
	var h2lmaps [][]int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		tag := strings.Split(line,":")
		switch tag[0] {
			case "seeds" :
				fmt.Println("Found seeds")
				s := strings.Split(tag[1]," ")
				for _,n := range s {
					if n != "" {
						seed,_ := strconv.Atoi(n)
						seeds = append(seeds, seed)
					}
				}
				fmt.Println(seeds)
			case "seed-to-soil map" :
				// assign maplets to s2smaps <= list of [dest_start, source_start, len]
				fmt.Println("Found seed-to-soil map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					s2smaps = append(s2smaps, mkmaplet(line))
				}
				fmt.Println("s2smaps: ", s2smaps)
			case "soil-to-fertilizer map" :
				fmt.Println("Found soil-to-fertilizer map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					s2fmaps = append(s2fmaps, mkmaplet(line))
				}
				fmt.Println("s2fmaps: ", s2fmaps)
			case "fertilizer-to-water map" :
				fmt.Println("Found fertilizer-to-water map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					f2wmaps = append(f2wmaps, mkmaplet(line))
				}
				fmt.Println("f2wmaps: ", f2wmaps)
			case "water-to-light map" :
				fmt.Println("Found water-to-light map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					w2lmaps = append(w2lmaps, mkmaplet(line))
				}
				fmt.Println("w2lmaps: ", w2lmaps)
			case "light-to-temperature map" :
				fmt.Println("light-to-temperature map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					l2tmaps = append(l2tmaps, mkmaplet(line))
				}
				fmt.Println("l2tmaps: ", l2tmaps)
			case "temperature-to-humidity map" :
				fmt.Println("temperature-to-humidity map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					t2hmaps = append(t2hmaps, mkmaplet(line))
				}
				fmt.Println("t2hmaps: ", t2hmaps)
			case "humidity-to-location map" :
				fmt.Println("humidity-to-location map")
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					h2lmaps = append(h2lmaps, mkmaplet(line))
				}
				fmt.Println("h2lmaps: ", h2lmaps)
		}
	}

	min := -1
	for _,seed := range seeds {
		fmt.Println("seed: ",seed)
		soil := mapping(seed, s2smaps)
		fmt.Println(soil)
		fertilizer := mapping(soil, s2fmaps)
		fmt.Println(fertilizer)
		water := mapping(fertilizer, f2wmaps)
		fmt.Println(water)
		light := mapping(water, w2lmaps)
		fmt.Println(light)
		temperature := mapping(light, l2tmaps)
		fmt.Println(temperature)
		humidity := mapping(temperature, t2hmaps)
		fmt.Println(humidity)
		location := mapping(humidity, h2lmaps)
		fmt.Println("location: ",location)
		if min == -1 {
			min = location
		}
		if location < min {
			min = location
		}
	}
	fmt.Println(min)
	fmt.Println("Done ...")
}

func mkmaplet(line string) []int {
	var maplet []int
	ss := strings.Split(line, " ")
	for _,s := range ss {
		n,_ := strconv.Atoi(s)
		maplet = append(maplet, n)
	}
	return maplet
}

func mapping(input int, maps [][]int) int {
	output := input
	for _,m := range maps {
		dest_start := m[0]
		source_start := m[1]
		len := m[2]
		if input >= source_start && input < source_start + len {
			output = dest_start + (input - source_start)
		}
	}
	return output
}