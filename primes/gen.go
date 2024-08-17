package primes

import (
	"math/big"
)

/*
	func Reversed[V any](s []V) iter.Seq[V] {
		return func(yield func(V) bool) {
			for i := len(s) - 1; i >= 0; i-- {
				if !yield(s[i]) {
					return
				}
			}
		}
	}
*/
// func GenCandidates() iter.Seq[*big.Int] bool) {
func GenCandidates(start *big.Int) func(yield func(v *big.Int) bool) {
	return func(yield func(v *big.Int) bool) {
		for i := new(big.Int).Set(start); ; i.Add(i, big.NewInt(2)) {
			if !yield(i) {
				return
			}
		}
	}
}
