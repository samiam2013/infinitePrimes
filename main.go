package main

func main() {
	for i, k := range primes.Gen() {
		if i > 100 {
			break
		}
		println(i)
	}
}
