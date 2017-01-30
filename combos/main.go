package main

import (
	"fmt"
	"os"
	"math"
	"strings"
)

//Runtime O(n2^k) s.t. k=xcount
func findCombos(pattern string) {
	//find number of Xs
	xcount := 0
	for _, c := range pattern {
		if string(c) == "X" {
			xcount++
		}
	}
	//Generate all possible combos iteratively
	perms := int(math.Pow(2, float64(xcount)))
	for i:=0; i<perms; i++ {
		curr := i
		xs := 0
		for j:=0; j<len(pattern); j++ {
			if string(pattern[j]) != "X" {
				fmt.Print(string(pattern[j]))
			} else {
				xs++
				d:=xcount-xs
				if curr < int(math.Pow(2, float64(d))) {
					fmt.Print("0")
				} else {
					fmt.Print("1")
					curr -= int(math.Pow(2, float64(d)))
				}
			}
		}
		fmt.Println()
	}
}

func isValidInput(pattern string) bool{
	s:="10X"
	for _, c := range pattern {
		if !strings.Contains(s,string(c)) {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Bad input, try again")
		os.Exit(1)
	}
	pattern := args[0]
	if len(pattern) == 0 {
		fmt.Println("Bad input, try again")
		os.Exit(1)
	}

	if ok := isValidInput(pattern); !ok {
		fmt.Println("Bad input, try again")
		os.Exit(1)
	}
	findCombos(pattern)
}
