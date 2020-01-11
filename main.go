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
	var a [12]int
	var number int
	var inList, added bool
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
				if *end-*start+1 > *n {
					for added != true {
						number = randInterval(*start, *end)
						for i := 0; i < 12; i++ {
							if number == a[i] {
								inList = true
							}
						}
						if inList == false {
							a[i] = number
							added = true
							fmt.Println(number)
						} else {
							added = false
							inList = false
						}
					}
					added = false
				} else {
					fmt.Println("Error: all numbers cannot be unique with this params")
					break
				}
			}
		}
	}
}
