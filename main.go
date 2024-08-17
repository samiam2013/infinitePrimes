package main

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/samiam2013/infiniteprimes/primes"
)

func main() {
	c := make(chan *big.Int, 1024)
	go func() {
		for {
			select {
			case p := <-c:
				fmt.Printf("%s is prime.\n", p.String())
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	for i := range primes.GenCandidates(int64(math.Pow(2, 32) - 1)) {
		// fmt.Println("testing candidate", i.String())
		k := new(big.Int)
		k.Set(i)
		go func(m *big.Int) {
			if primes.IsPrime(m) {
				c <- m
			}
		}(k)
	}

}
