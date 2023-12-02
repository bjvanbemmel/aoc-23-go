package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var digits []int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatalf("Something went wrong when opening file: `%s`", err)
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		clone := strings.Clone(line)
		for _, num := range numbers {
			i := 0
			for i > -1 {
				i = strings.Index(line, num)
				if i < 0 {
					continue
				}

				slice := []byte(line)
				newSlice := slice[:i]
				newSlice = append(newSlice, slice[i])
				newSlice = append(newSlice, alphaToNumeric(num))
				newSlice = append(newSlice, slice[i+len(num)-1])
				newSlice = append(newSlice, slice[i+len(num):]...)

				line = string(newSlice)
			}
		}

		foundDigits := regexp.MustCompile("[0-9]").FindAllString(line, -1)
		if len(foundDigits) < 1 {
			continue
		}

		num, _ := strconv.Atoi(fmt.Sprintf("%v%v", foundDigits[0], foundDigits[len(foundDigits)-1]))
		fmt.Println(clone, line, num)
		digits = append(digits, num)
	}

	var sum int
	for _, num := range digits {
		sum += num
	}

	fmt.Println(sum)
}

func alphaToNumeric(val string) byte {
	switch val {
	case "one":
		return '1'
	case "two":
		return '2'
	case "three":
		return '3'
	case "four":
		return '4'
	case "five":
		return '5'
	case "six":
		return '6'
	case "seven":
		return '7'
	case "eight":
		return '8'
	case "nine":
		return '9'
	default:
		return '0'
	}
}
