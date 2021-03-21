package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func jijo(x int) int {
	return x * x
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	N := nextInt()
    
    for j := 0 ; j < N; j++{
        n := nextInt()
		 
        x := 0
        for i := 0; i < n; i++ {
			k := nextInt()
			if k < 0 {continue}
			x += jijo(k)
           
    }
    fmt.Println(x)
}
}