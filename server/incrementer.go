package main

func incrementer(id chan<- int) {
	i := 0
	for {
		id <- i
		i++
	}
}
