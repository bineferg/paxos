package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)


//Used for findPair and bonus
var utxos = make(map[string]int)
//Used for findPair
var utxosPairs = make(map[int][]string)
//Used for bonus
var utxosK = make(map[int][]Utxos)
var minDiff = -1

type Utxos struct {
	id string
	val int
}

func parseUtxos(line string) {
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
	utxos[key] = val

}

//Runtime - worst case O(n^2) due to possible duplicate integers with unique ids
func findPair(target int) {
	for k1, v1 := range utxos {
		for k2, v2 :=range utxos {
			if v1 + v2 >= target && k1 != k2 {
				dif := (v1 + v2) - target
				utxosPairs[dif]=[]string{k1, k2}
				if minDiff == -1 || dif < minDiff{
					minDiff = dif
				}

			}
		}
	}
	if len(utxosPairs) != 0 {
		fmt.Println(utxosPairs[minDiff][0], " ", utxos[utxosPairs[minDiff][0]],", ", utxosPairs[minDiff][1], " ", utxos[utxosPairs[minDiff][1]])
	} else {
		fmt.Println("Not possible.")
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

func subsetBonus(target int) {

	ulist := make([]Utxos, len(utxos))
	count := 0
	for k, v :=range utxos {
		u := Utxos{id: k, val: v}
		ulist[count] = u
		count++
	}

	var ps = powerSet(ulist)
	for _, sub := range ps {
		sum :=0
		for _, el := range sub {
			sum+=el.val
		}
		if sum >= target {
			dif := sum-target
			utxosK[dif] = sub
			if minDiff == -1 || dif < minDiff{
				minDiff = dif
			}

		}
	}
	minUtxos := utxosK[minDiff]
	for _, u := range minUtxos {
		fmt.Print(u.id," ",u.val," ")
	}
	fmt.Println()

}


func main(){
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
	s := bufio.NewScanner(f)
	for s.Scan() {
		parseUtxos(s.Text())
	}
	t, _ := strconv.Atoi(target)
	if len(args) == 3 && args[2] == "-bonus" {
		subsetBonus(t)
	} else {
		findPair(t)

	}

}
