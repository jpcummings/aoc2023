package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


type Race struct {
	time int
	record int
}

func main() {

	var time []int
	var record []int
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		snum := strings.Fields(line)
		ltype := "unkown"
		for _,w := range(snum) {
			if w == "Time:" {
				ltype = "time"
				continue
			} else if w == "Distance:" {
				ltype = "record"
				continue
			} else {  // assume it is a number
				switch ltype {
					case "time":
						t,_ := strconv.Atoi(w)
						time = append(time, t)
					case "record":
						d,_ := strconv.Atoi(w)
						record = append(record, d)
				}
			}
			
		}
	}

	mult := get_margin(time, record)
	fmt.Println(mult)

	// rekern input data - time and record
	time = rekern(time)
	record = rekern(record)

	mult = get_margin(time, record)
	fmt.Println(mult)

}

func rekern(list []int) []int {
	// rekern input data - time or record
	var ret []int

	st := ""
	for _,t := range list {
		st += strconv.Itoa(t)
	}
	t,_ := strconv.Atoi(st)
	ret = append(ret, t)
	return ret
}

func dist(btime, ttime int) int {
	return btime*(ttime-btime)
}

func get_margin(time, record []int) int {
	mult := 1
	for i,_ := range(time) {
		nwins := 0
		for button := 0; button < time[i]; button++ {
			d := dist(button, time[i])
			if d > record[i] {
				nwins++
			}
		}
		mult *= nwins
	}
	return mult
}
