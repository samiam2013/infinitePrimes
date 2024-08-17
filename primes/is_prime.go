package primes

import "math/big"

func IsPrime(candidate *big.Int) bool {
	if candidate.Cmp(big.NewInt(1)) <= 0 {
		return false
	}
	sqrt := new(big.Int).Sqrt(candidate)
	// if !candidate.ProbablyPrime(20) {
	// 	return false
	// }
	for i := big.NewInt(2); i.Cmp(sqrt) <= 0; i.Add(i, big.NewInt(1)) {
		if new(big.Int).Mod(candidate, i).Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}
	return true
}
