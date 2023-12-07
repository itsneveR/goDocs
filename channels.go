/*
https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go#communicating-safely-between-goroutines-with-channels
*/

/*

func readChannel(ch <-chan int) {
	// ch is read-only
}

*/

/*

func writeChannel(ch chan<- int) {
	// ch is write-only
}

*/

/*

intChan := make(chan int)
intChan <- 10 //write to channel

*/

/*

intChan := make(chan int)
intVar := <- intChan //read from channel

*/

