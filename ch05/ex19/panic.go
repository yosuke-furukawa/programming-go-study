package main

import "log"

func chaos() (n int) {
	defer func() {
		p := recover()
		n = p.(int)
	}()
	panic(1)
}

func main() {
	c := chaos()
	log.Print(c)
}
