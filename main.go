package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed     = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start    = flag.Int("start", 1, "Additional param start for random generator.")
	end      = flag.Int("end", 6, "Additional param end for random generator.")
	n        = flag.Int("n", 1, "Additional param n for random generator cycle.")
	norepeat = flag.Bool("norepeat", false, "Additional param norepeat for random generator cycle.")
)

func randInterval(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}

	if *start > *end {
		fmt.Println("Error: start > end")
	} else {
		for i := 0; i < *n; i++ {
			if *norepeat == false {
				fmt.Println(randInterval(*start, *end))
			} else {
				fmt.Println("Work!")
			}
		}
	}
}
