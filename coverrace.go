package coverrace

func add100() int {
	total := 0
	c := make(chan int, 1)
	for i := 0; i < 100; i++ {
		go func(chan int) {
			c <- 1
		}(c)
	}
	for u := 0; u < 100; u++ {
		total += <-c
	}
	return total
}
