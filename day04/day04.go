package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"math"
)

type Card struct {
	id int
	copies int
	winners []int
	mine []int
}

func main() {

	var cards []Card

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan()) {
		line := scanner.Text()
		var winners []int
		var mine []int

		tmp := strings.Split(line,": ")
		card := tmp[0]
		nums := tmp[1]
		tmp = strings.Split(card, " ")
		id,_ := strconv.Atoi(tmp[1])
		tmp = strings.Split(nums," | ")
		winnums := strings.Split(tmp[0]," ")
		mynums := strings.Split(tmp[1]," ")
		for _, num := range winnums {
			w, _ := strconv.Atoi(num)
			if w != 0 {
				winners = append(winners, w)
			}
		}
		for _, num := range mynums {
			w, _ := strconv.Atoi(num)
			if w != 0 {
				mine = append(mine, w)
			}
		}
		//fill card
		cards = append(cards, Card{id,1,winners,mine})
	}

	//score it
	total := 0
	for _,card := range cards {
		cardscore := 0
		nmatch := matches(card)
		if nmatch > 0 {
			cardscore = int(math.Pow(2,float64(nmatch-1)))
		}
		total += cardscore
	}
	fmt.Println(total)

	// score part 2
	for icard,_ := range cards {
		nm := matches(cards[icard])
		for jcard := icard+1; jcard <= icard+nm; jcard++ {
			cards[jcard].copies+=cards[icard].copies
		}
	}
	total = 0
	for _,c := range cards {
		total += c.copies
	}
	fmt.Println(total)


}

func matches(c Card) int {
	nmatch := 0
	for _,m := range c.mine {
		for _,w := range c.winners {
			if m == w {
				nmatch++
			}
		}
	}
	return nmatch
}