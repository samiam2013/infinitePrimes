package main

import (
	"math/big"
	"time"

	"github.com/samiam2013/infiniteprimes/primes"
	"github.com/sirupsen/logrus"
)

type computation struct {
	candidate *big.Int
	isPrime   bool
	timeTaken time.Duration
}

func main() {
	c := make(chan *computation)
	go recieve(c)
	start := big.NewInt(1) // int64(math.Pow(2, mersennePower) - 2))
	logrus.Infof("starting at %s", start.String())
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

const avgLen = 100

func recieve(c chan *computation) {
	latestDurs := make([]*time.Duration, 0)
	for {
		select {
		case p := <-c:
			if p.isPrime {
				logrus.Infof("found prime %s in %.2fs ", p.candidate.String(), p.timeTaken.Seconds())
				latestDurs = append(latestDurs, &p.timeTaken)
				if len(latestDurs) > avgLen {
					latestDurs = latestDurs[1:]
				}
				totalDur := time.Duration(0)
				for _, d := range latestDurs {
					totalDur += *d
				}
				avgDur := totalDur / time.Duration(len(latestDurs))
				logrus.Infof("avg for %d primes: %.2fs", len(latestDurs), avgDur.Seconds())
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
