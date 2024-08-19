package primes

import (
	"math/big"
)

func GenCandidates(start *big.Int) func(yield func(v *big.Int) bool) { 
// func GenCandidates() iter.Seq[*big.Int] bool) { // /iter.Seq(2?)/ simplifies the func(yeild func(...
	return func(yield func(v *big.Int) bool) {  // a function that go will generate from the `for v := range ...`
		for i := new(big.Int).Set(start); ; i.Add(i, big.NewInt(2)) { // i = start; ; i++ but with arbitrarily large integers
			if !yield(i) { // output using the function, 
				return // and if the for loop we're driving has exited, return
			}
		}
	}
}
