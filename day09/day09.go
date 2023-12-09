package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main() {
	var histories [][]int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var history []int
		line := scanner.Text()
		snum := strings.Split(line," ")
		for _,s := range(snum) {
			i,_ := strconv.Atoi(s)
			history = append(history, i)
		}
		histories = append(histories, history)
	}

	sum := 0
	sumback := 0
	i := 0
	for _,hist := range(histories) {
		i++
		deltas := makeDeltas(hist)
		extval := extrapolate(deltas)
		extbackval := extrapolate_back(deltas)
		sum += extval
		sumback += extbackval
	}
	fmt.Println(sum)
	fmt.Println(sumback)

}


func extrapolate(seqs [][]int) int {
	
	if allzeros(seqs[0]) {
		return 0
	} else {
		add := extrapolate(seqs[1:])
		ext := add+seqs[0][len(seqs[0])-1]
		return ext
	}
}

func extrapolate_back(seqs [][]int) int {
	
	if allzeros(seqs[0]) {
		return 0
	} else {
		sub := extrapolate_back(seqs[1:])
		ext := seqs[0][0]-sub
		return ext
	}
}


func makeDeltas(h []int) [][]int {
	var delts [][]int
	delts = append(delts, h)
	for allzeros(delts[len(delts)-1]) != true {
		var delt []int
		d := delts[len(delts)-1]
		for i:=0; i<len(d)-1; i++ {
			delt = append(delt, d[i+1]-d[i])
		}
		delts = append(delts, delt)
	}
	return delts
}


func allzeros(h []int) bool {
	for _,v := range(h) {
		if v != 0 {return false}
	}
	return true
}

