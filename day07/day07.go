package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Hand struct {
	hand string
	bid int
}

type ByHand []Hand
func (a ByHand) Len() int           { return len(a) }
func (a ByHand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool { return handLess(a[i], a[j]) }

type ByBid []Hand
func (a ByBid) Len() int           { return len(a) }
func (a ByBid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBid) Less(i, j int) bool { return bidLess(a[i], a[j]) }

var wildcard bool = true

func main() {

	var hands []Hand
	
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		s := strings.Fields(line)
		hand := s[0]
		sbid := s[1]
		bid,_ := strconv.Atoi(sbid)
		hands = append(hands, Hand{hand,bid})
		identifyHand(Hand{hand,bid})
	}
//	fmt.Println(hands)

	sort.Sort(ByHand(hands))
//	fmt.Println(hands)
	fmt.Println(score(hands))
//	fmt.Println(handLess(Hand{"Q5T46",0}, Hand{"87Q42",0}))
//	fmt.Println(handLess(Hand{"JQ85Q",0}, Hand{"Q9928",0}))
//	fmt.Println(handLess(Hand{"T44T4",0}, Hand{"QTQTT",0}))
//	fmt.Println(handLess(Hand{"32T3K",0}, Hand{"KTJJT",0}))

}

func score(hands []Hand) int {
	ret := 0
	for i,h := range(hands) {
		ret += (i+1)*h.bid
	}
	return ret
}

func handLess(h1, h2 Hand) bool {
	v1,_ := identifyHand(h1)
	v2,_ := identifyHand(h2)
	if v2 != v1 {
		return v2 > v1
	} else {
		i := 0
		for h1.hand[i] == h2.hand[i] {
			i++
		}
		return card2int(string(h1.hand[i])) < card2int(string(h2.hand[i]))
	}
}

func bidLess(h1, h2 Hand) bool {
	return h1.bid < h2.bid
}

func card2int(c string) int {
	v :=0
	switch c {
		case "2":
			v = 2
		case "3":
			v = 3
		case "4":
			v = 4
		case "5":
			v = 5
		case "6":
			v = 6
		case "7":
			v = 7
		case "8":
			v = 8
		case "9":
			v = 9
		case "T":
			v = 10
		case "J":
			if wildcard {
				v = 1
			} else {
				v = 11
			}
		case "Q":
			v = 12
		case "K":
			v = 13
		case "A":
			v = 14
	}
	return v
}

func identifyHand(hand Hand) (int, string) {
	hands := make(map[string]int)

	h := hand.hand
	htype := ""
	hnum := 0
	var nones, ntwos, nthrees, nfours, nfives int

	for _,c := range(h) {
		hands[string(c)]++
	}
	for _,v := range hands {
		switch v {
			case 1:
				nones++
			case 2:
				ntwos++
			case 3:
				nthrees++
			case 4:
				nfours++
			case 5:
				nfives++
		}
	}
	if nfives == 1 {
		htype = "five"
		hnum = 7
	}
	if nfours == 1 {
		htype = "four"
		hnum = 6
	}
	if nthrees == 1 {
		if ntwos == 1 {
			htype = "full"
			hnum = 5
		} else {
			htype = "three"
			hnum = 4
		}
	}
	if ntwos == 2 {
		htype = "twopair"
		hnum = 3
	}
	if ntwos == 1 && nthrees == 0 {
		htype = "pair"
		hnum = 2
	}
	if nones == 5 {
		htype = "high"
		hnum = 1
	}

	if wildcard && hands["J"] > 0 {
		switch hnum {
			case 1:  {
				hnum = 2
				htype = "pair"
			}
			case 2:  {
				hnum = 4
				htype = "three"
			}
			case 3: {
				if hands["J"] == 1 {
					hnum = 5
					htype = "full"
				} else if hands["J"] == 2 {
					hnum = 6
					htype = "four"
				}
			}
			case 4:  {
				hnum = 6
				htype = "four"
			}
			case 5:  {
				hnum = 7
				htype = "five"
			}
			case 6:  {
				hnum = 7
				htype = "five"
			}
			case 7: {
				hnum = 7
				htype = "five"
			}
		}
	}

	return hnum, htype
}

