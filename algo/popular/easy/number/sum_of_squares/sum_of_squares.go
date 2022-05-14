package main

func sumOfSquares(c, quit chan int) {
	y := 1

	for {
		select {
		case c <- (y * y):
			y++
		case <-quit:
			return
		}
	}
}

func calcSumOfSquares(n int) int {
	iteratorCh := make(chan int)
	quitCh := make(chan int)
	resultCh := make(chan int)

	sum := 0
	go func() {
		for i := 0; i < n; i++ {
			sum += <-iteratorCh
		}

		quitCh <- 0
		resultCh <- sum
	}()

	sumOfSquares(iteratorCh, quitCh)

	return <-resultCh
}
