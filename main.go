package main

import (
	"fmt"
	"math/rand"
)

func main() {
	grid, err := NewGrid(randomGridString())

	if err != nil {
		fmt.Printf("%v\n",err)
		return
	}

	grid.Print()
}

func randomGridString() string {
	s := make([]byte, 81)
	for i:= 0 ; i < 81; i++ {
		p := rand.Intn(100)

		if p < 80 {
			s[i] = '.'
		} else {
			n := rand.Intn(9) + 1
			s[i] = byte(int('0') + n)
		}
	}

	return string(s)
}