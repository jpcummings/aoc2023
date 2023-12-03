package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
)


func main() {

	code := 0
	nums := [10]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	scanner := bufio.NewScanner(os.Stdin)

	for ( scanner.Scan() ) {
		line := scanner.Text()

		firstdigit := ""
		for i := 0; i <=len(line) ; i++ {
			foundfirst := false
			// is it a spelled out number?
			for n, num := range nums {
				if (i+len(num))<len(line) && line[i:i+len(num)] == num {
					firstdigit = strconv.Itoa(n)
					foundfirst = true
					break
				}
			}
			if foundfirst { break }

			// is it a digit?
			c := line[i]
			if unicode.IsDigit(rune(c)) {
				firstdigit = string(c)
				break
			} 			
		}

		lastdigit := ""
		for i := 0; i <len(line) ; i++ {
			// is it a spelled out number?
			for n, num := range nums {
				if (i+len(num))<=len(line) && line[i:i+len(num)] == num {
					lastdigit = strconv.Itoa(n)
					break
				}
			}
			// is it a digit?
			c := line[i]
			if unicode.IsDigit(rune(c)) {
				lastdigit = string(c)
			} 			
		}

		i,_ := strconv.Atoi(firstdigit+lastdigit)
		code += i
	}
	fmt.Println(code)
}