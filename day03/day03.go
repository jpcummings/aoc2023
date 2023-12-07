package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
	"slices"
)

type Pair struct {
	n int
	gname string
}


func main() {

	var nrows int
	var irow int
	var maprow []string
	var treemap [][]string

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line := scanner.Text()
		nrows++
		maprow = nil
		for _, c := range line {
			maprow = append(maprow, string(c))
		}
		treemap = append(treemap,maprow)
		irow++
	}
	p, gearpairs := findparts(treemap)
	sum := 0
	for _,n := range p {
		sum += n
	}
	fmt.Println(sum)
	fmt.Println(gearratios(gearpairs))
}

func gearratios(pairs []Pair) int {

	var geardone []string
	sum := 0
	ratio := 1

	for _, p := range pairs {
		if !slices.Contains(geardone, p.gname) {
			nterms := 0
			ratio = 1
			gearon := p.gname
			geardone = append(geardone, gearon)
			// multiply em up
			for _, pp := range pairs {
				if pp.gname == gearon {
					ratio *= pp.n
					nterms++
				}
			}
			if nterms == 2 {
				sum += ratio
			}
			nterms = 0
		}
	}
	return sum
}

func findparts(schematic [][]string) ([]int, []Pair) {

	var parts []int
	innum := false
	symbolNeighbor := false
	foundgear := false
	gearname := ""
	var gearpairs []Pair

	num := ""
	for r,_ := range schematic {
		for c,char := range schematic[r] {
			if unicode.IsDigit(rune(char[0])) {
				// find end
				num = num+string(char)
				if checkNeighbors(schematic,r,c) {
					symbolNeighbor = true
					foundgear, gearname = gearNeighbor(schematic,r,c)
				}
				innum = true
				continue
			} else if innum  {
				n, _ := strconv.Atoi(num)
				if symbolNeighbor == true {
					parts = append(parts,n)
					symbolNeighbor = false
					if foundgear {
						p := Pair{n,gearname}
						gearpairs = append(gearpairs, p)
					}
				} else {
					//fmt.Println(n, " has NO neighbor")
				}
				num = ""
				innum = false
			}
		}
	}
	return parts, gearpairs
}

func gearNeighbor(schematic [][]string, r, c int) (bool, string) {

	foundgear := false
	gearname := ""

	maxi := len(schematic)
	maxj := len(schematic[1])

	for i := r-1; i <= r+1; i++ {
		for j := c-1; j <= c+1; j++ {
			if i >= 0 && i < maxi && j >= 0 && j < maxj {
				char := schematic[i][j]
				if char == "*" {
					foundgear = true
					gearname = "gear"+strconv.Itoa(i)+"-"+strconv.Itoa(j)
				}
			}
		}
	}
	return foundgear, gearname
}


func checkNeighbors(schematic [][]string, r, c int) bool {

	foundSym := false

	maxi := len(schematic)
	maxj := len(schematic[1])

	for i := r-1; i <= r+1; i++ {
		for j := c-1; j <= c+1; j++ {
			if i >= 0 && i < maxi && j >= 0 && j < maxj {
				char := schematic[i][j]
				if char != "." && !unicode.IsDigit(rune(char[0])) {
					foundSym = true
				}
			}
		}
	}
	return foundSym
}
