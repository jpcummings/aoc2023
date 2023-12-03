package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)



type Game struct {
	ID int
	draws []map[string]int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var games  []Game

	for ( scanner.Scan() ) {
		line := scanner.Text()		
		games = append(games, Game{ID(line),Draws(line)})
	}

	sum := 0
	for _, g := range games {
		if IsPossible(g, 12, 13, 14) {
			// fmt.Println("game ", g.ID, " is possible")
			sum += g.ID
		}
	}
	fmt.Println(sum)

	sum = 0
	for _, g := range games {
		sum += Power(g)
	}
	fmt.Println(sum)
}

func IsPossible(game Game, red int, green int, blue int) bool {
	possible := true
	for _, d := range game.draws {
		if d["red"] > red || d["green"] > green || d["blue"] > blue { possible = false }
	}
	return possible
}

func Power(game Game) int {
	maxred, maxgreen, maxblue := 0,0,0
	for _, d := range game.draws {
		if d["red"] > maxred { maxred = d["red"] }
		if d["green"] > maxgreen { maxgreen = d["green"] }
		if d["blue"] > maxblue { maxblue = d["blue"] }
	}
	return maxred*maxgreen*maxblue
}

func Draws(game string) []map[string]int {
	var ret []map[string]int
	s := strings.Split(game, ":")
	draws := strings.Split(s[1], ";")
	for _, draw := range draws {
		m := make(map[string]int)
		cols := strings.Split(draw, ",")
		for _, col := range cols {
			c := strings.Split(col, " ")
			m[c[2]], _ = strconv.Atoi(c[1])
		}
		ret = append(ret, m)
	}
	return ret
}


func ID(game string) int {
	s := strings.Split(game, ":")
	gid := strings.Split(s[0]," ")
	ret, _ := strconv.Atoi(gid[1])
	return ret
}
