package main

import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Node struct {
	ldest, rdest string
}

var directions string
var idir int

func main() {

	var source string
	var dest []string
	var nstep int
	var nsteps []int
	network := make(map[string]Node)

	part2 := flag.Bool("part2", false, "execute part 2")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { continue }
		if directions == "" {
			directions = line
		} else {
			t := strings.Split(line, " = ")
			source = t[0]
			dest = strings.Split(strings.Trim(t[1],"()"), ", ")
			network[source] = Node{dest[0], dest[1]}
		}
	}

	// now navigate network

	if !(*part2) {
		fmt.Println("executing part 1...")
		nstep = navigate("AAA", "ZZZ", network)
		fmt.Println(nstep)
	} else {
		fmt.Println("executing part 2...")
	
		starts := findAs(network)
		for _,start := range(starts) {
			resetDir()
			nstep = multi_navigate(start, "Z", network)
			nsteps = append(nsteps, nstep)
		}
		for len(nsteps) > 1 {
			temp := lcm(nsteps[0],nsteps[1])
			nsteps = nsteps[2:]
			nsteps = append(nsteps, temp)
		}
		fmt.Println(nsteps[0])
	}
	
}

func findAs(network map[string]Node) []string {

	var starts []string
 
	for k,_ := range(network) {
		if k[len(k)-1:] == "A" {
			starts = append(starts, k)
		}
	}
	return starts
}

func navigate(start, end string, network map[string]Node) int {
	nstep := 0
	pos := start
	for pos != end {
		nstep++
		step := getNextDir()
		node := network[pos]
		if step == "L" {
			pos = node.ldest
		} else if step == "R" {
			pos = node.rdest
		} else {
			fmt.Println("unkown direction")
			os.Exit(1)
		}
	}
	return nstep
}

func multi_navigate(start, end string, network map[string]Node) int {
	nstep := 0
	pos := start
	for pos[len(pos)-1:] != end {
		nstep++
		step := getNextDir()
		node := network[pos]
		if step == "L" {
			pos = node.ldest
		} else if step == "R" {
			pos = node.rdest
		} else {
			fmt.Println("unkown direction")
			os.Exit(1)
		}
	}
	return nstep
}



func resetDir() {
	idir = 0
}

func getNextDir() string {

	if idir >= len(directions) {
		idir = 0
	}
	ret := string(directions[idir])
	idir++

	return ret
}


func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if b > a { a, b = b, a }
	return (a/gcd(a, b))*b
}
