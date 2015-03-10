package main

import (
	"flag"
	"fmt"
	"strings"
)

func makePairs(n int, i int) [][2]int {
	result := make([][2]int, 0, n/2)

	result = append(result, [2]int{0, i})
	for k := 1; 2*k < n-1; k++ {

		a := 1 + (i-1+k)%(n-1)
		b := 1 + (i-1+(n-1)-k)%(n-1)
		result = append(result, [2]int{a, b})
	}

	return result
}

func usage() {
	fmt.Println("Usage: pairmaker Name1,Name2[,Name3...]")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}
	names := flag.Arg(0)
	nameList := strings.Split(names, ",")
	n := len(nameList)
	if n%2 == 1 {
		nameList = append(nameList, "__")
		n++
	}
	for i := 1; i < n; i++ {
		for j, pair := range makePairs(n, i) {
			if j > 0 {
				fmt.Print(",")
			}
			fmt.Printf("%v,%v", nameList[pair[0]], nameList[pair[1]])
		}
		fmt.Println()
	}
}
