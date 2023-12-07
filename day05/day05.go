package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {

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
		tag := strings.Split(line,":")
		switch tag[0] {
			case "seeds" :
				s := strings.Split(tag[1]," ")
				for _,n := range s {
					if n != "" {
						seed,_ := strconv.Atoi(n)
						seeds = append(seeds, seed)
					}
				}
			case "seed-to-soil map" :
				// assign maplets to s2smaps <= list of [dest_start, source_start, len]
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					s2smaps = append(s2smaps, mkmaplet(line))
				}
			case "soil-to-fertilizer map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					s2fmaps = append(s2fmaps, mkmaplet(line))
				}
			case "fertilizer-to-water map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					f2wmaps = append(f2wmaps, mkmaplet(line))
				}
			case "water-to-light map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					w2lmaps = append(w2lmaps, mkmaplet(line))
				}
			case "light-to-temperature map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					l2tmaps = append(l2tmaps, mkmaplet(line))
				}
			case "temperature-to-humidity map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					t2hmaps = append(t2hmaps, mkmaplet(line))
				}
			case "humidity-to-location map" :
				for scanner.Scan() {
					line := scanner.Text()
					if line == "" { break }
					h2lmaps = append(h2lmaps, mkmaplet(line))
				}
		}
	}

	min := -1
	for _,seed := range seeds {
		soil := mapping(seed, s2smaps)
		fertilizer := mapping(soil, s2fmaps)
		water := mapping(fertilizer, f2wmaps)
		light := mapping(water, w2lmaps)
		temperature := mapping(light, l2tmaps)
		humidity := mapping(temperature, t2hmaps)
		location := mapping(humidity, h2lmaps)
		if min == -1 {
			min = location
		}
		if location < min {
			min = location
		}
	}
	fmt.Println(min)
	min = -1
	for i:=0; i<len(seeds); i+=2 {
		fmt.Println("range length: ", seeds[i+1])
		for seed:=seeds[i]; seed<seeds[i]+seeds[i+1]; seed++ {
		//	fmt.Println(seed)
			soil := mapping(seed, s2smaps)
			fertilizer := mapping(soil, s2fmaps)
			water := mapping(fertilizer, f2wmaps)
			light := mapping(water, w2lmaps)
			temperature := mapping(light, l2tmaps)
			humidity := mapping(temperature, t2hmaps)
			location := mapping(humidity, h2lmaps)
			if min == -1 {
				min = location
			}
			if location < min {
				min = location
			}
		}
	}
	fmt.Println(min)

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