/*
	Write a general-purpose unit-conversion program analogous to `cf` that reads
	numbers from its command-line arguments or from the standard input if there
	are no arguments, and converts each number into units like temperature in `Celsius`
	and `Fahrenheit`, length in `feet` and `meters`, weight in `pounds` and `kilograms`, and the like.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pounds float64
type kilograms float64

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Which unit you want to convert?\nq) Temperature\nw) Length\ne) Weight")

	for scanner.Scan() {
		choice := scanner.Text()

		if choice == "q" {
			fmt.Println("Do temperature things")
			fmt.Println("Which unit you want to convert?\nq) Temperature\nw) Length\ne) Weight")
		} else if choice == "w" {
			fmt.Println("Do length things")
			fmt.Println("Which unit you want to convert?\nq) Temperature\nw) Length\ne) Weight")
		} else if choice == "e" {
			fmt.Println("a [value]) Pounds to Killograms\nb [value]) Killograms to Pounds")
		}
		
		if choice[0] == 'a' {
			inputValue, err := strconv.Atoi(strings.Split(choice, " ")[1])
			if err != nil {
				fmt.Println("Wrong value")
				continue
			}

			fmt.Println(pounds(inputValue).toKilograms())
			fmt.Println("Which unit you want to convert?\nq) Temperature\nw) Length\ne) Weight")
		} else if choice[0] == 'b' {
			inputValue, err := strconv.Atoi(strings.Split(choice, " ")[1])
			if err != nil {
				fmt.Println("Wrong value")
				continue
			}

			fmt.Println(kilograms(inputValue).toPounds())
			fmt.Println("Which unit you want to convert?\nq) Temperature\nw) Length\ne) Weight")
		}
	}
}

func (p pounds) toKilograms() kilograms {
	return kilograms(p / pounds(2.205))
}

func (k kilograms) toPounds() pounds {
	return pounds(k * kilograms(2.205))
}