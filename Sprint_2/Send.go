package main


func Send(ch1, ch2 chan int) {
	var list []int = []int{0, 1, 2}
	go func() {
		var val int
		for _, i := range list{
			val = i
			ch1 <- val
		}
		
	}()

	go func() {
		var val int
		for _, i := range list{
			val = i
			ch2 <- val
		}
		
	}()
}