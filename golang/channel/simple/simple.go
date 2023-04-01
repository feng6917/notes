package simple


func SimpleA() {
	ch := make(chan int)
	ch <- 1
	close(ch)
	<-ch
}

func SimpleA1() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	<-ch
}

func SimpleA2() {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	close(ch)
	go func() {
		<-ch
	}()
}
