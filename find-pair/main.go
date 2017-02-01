package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Used for findPair and bonus
var utxosDiff = make(map[int][]Utxos)
var minDiff = -1

type Utxos struct {
	id  string
	val int
}

func parseUtxos(line string) Utxos {
	spltLn := strings.Split(line, " ")
	if len(spltLn) != 2 {
		fmt.Println("File contains bad input.")
		os.Exit(1)
	}
	val, err := strconv.Atoi(spltLn[1])
	key := spltLn[0]
	if err != nil {
		fmt.Println("File contains bad input.")
		os.Exit(1)
	}
	return Utxos{id: key, val: val}

}

//Runtime - O(n)
func findPair(ulist []Utxos, target int) {
	left := 0
	right := len(ulist) - 1

	for left < right {
		sum := ulist[left].val + ulist[right].val
		if sum == target {
			minDiff = 0
			utxosDiff[minDiff] = []Utxos{ulist[left], ulist[right]}
			break
		}
		if sum >= target {
			diff := target - sum
			if minDiff == -1 {
				minDiff = diff
			}
			if minDiff > diff {
				minDiff = diff
			}
			utxosDiff[minDiff] = []Utxos{ulist[left], ulist[right]}
			right--
		}
		if sum < target {
			left++
		}
	}
	if len(utxosDiff) == 0 {
		fmt.Println("Not possible.")
	} else {
		fmt.Println(utxosDiff[minDiff][0].id, " ", utxosDiff[minDiff][0].val, ", ", utxosDiff[minDiff][1].id, " ", utxosDiff[minDiff][1].val)
	}
}

/////////////////////
//Bonus code  here //
/////////////////////
//Exponential solution to variation of NP-complete problem SubsetSum
//Run time is O(2^n) because we must consider all possible subsets
func powerSet(ulist []Utxos) [][]Utxos {
	if ulist == nil {
		return nil
	}
	ps := [][]Utxos{[]Utxos{}}
	for _, el := range ulist {
		var s [][]Utxos
		for _, ep := range ps {
			s = append(s, append(ep, el))
		}
		ps = append(ps, s...)
	}
	return ps
}

func subsetBonus(ulist []Utxos, target int) {

	var ps = powerSet(ulist)
	for _, sub := range ps {
		sum := 0
		for _, el := range sub {
			sum += el.val
		}
		if sum >= target {
			dif := sum - target
			utxosDiff[dif] = sub
			if minDiff == -1 || dif < minDiff {
				minDiff = dif
			}

		}
		if minDiff == 0 {
			break
		}
	}
	minUtxos := utxosDiff[minDiff]
	for _, u := range minUtxos {
		fmt.Print(u.id, " ", u.val, " ")
	}
	fmt.Println()

}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Bad input, run again")
		os.Exit(1)
	}
	fname := args[0]
	target := args[1]
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	var ulist []Utxos
	s := bufio.NewScanner(f)
	count := 0
	for s.Scan() {
		u := parseUtxos(s.Text())
		ulist = append(ulist, u)
		count++
	}
	t, _ := strconv.Atoi(target)
	if len(args) == 3 && args[2] == "-bonus" {
		subsetBonus(ulist, t)
	} else {
		findPair(ulist, t)

	}

}
