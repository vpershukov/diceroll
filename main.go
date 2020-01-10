package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "Additional param start for random generator.")
	end = flag.Int("end", 6, "Additional param end for random generator.")
	n = flag.Int("n", 1, "Additional param n for random generator cycle.")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(min, max int) int {
	return rand.Intn(max - min + 1) + min
}


func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	// Dice roll 1..6
	fmt.Println(randInterval(1,6))

	if *start > *end{
		fmt.Println("Error: program stopped")
	} else {
		fmt.Println(randInterval(*start, *end))
	}

	for i := 0; i < *n; i++ {
		fmt.Println(randInterval(*start, *end))
	}
}
