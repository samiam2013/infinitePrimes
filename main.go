package main

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/samiam2013/infiniteprimes/primes"
)

type computation struct {
	candidate *big.Int
	isPrime   bool
	timeTaken time.Duration
}

const mersennePower = 40

func main() {
	c := make(chan *computation, 1024)
	go recieve(c)
	start := big.NewInt(int64(math.Pow(2, mersennePower) - 1))
	fmt.Printf("starting at %s for mersenne number 2^%d-1\n", start.String(), mersennePower)
	for i := range primes.GenCandidates(start) {
		// fmt.Println("testing candidate", i.String())
		k := new(big.Int)
		k.Set(i)
		go func(m *big.Int) {
			startT := time.Now()
			isPrime := primes.IsPrime(m)
			dur := time.Since(startT)
			c <- &computation{m, isPrime, dur}
		}(k)
	}

}

func recieve(c chan *computation) {
	const avgLen = 10
	latestDurs := make([]*time.Duration, 0)
	for {
		select {
		case p := <-c:
			if p.isPrime {
				fmt.Printf("found prime %s in %.2fs ", p.candidate.String(), p.timeTaken.Seconds())
				latestDurs = append(latestDurs, &p.timeTaken)
				if len(latestDurs) > avgLen {
					latestDurs = latestDurs[1:]
				}
				totalDur := time.Duration(0)
				for _, d := range latestDurs {
					totalDur += *d
				}
				avgDur := totalDur / time.Duration(len(latestDurs))
				fmt.Printf("\tavg for %d primes: %.2fs\n", avgLen, avgDur.Seconds())
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
