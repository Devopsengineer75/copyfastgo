package tuto

import "fmt"

func MakeChannel() {

	messages := make(chan string) //,1) u wait one stuff infinite tunnel
	//Many threads many workers autonomous u can get info from workers from it.

	fmt.Println("1")

	//new independant variable u r receiver
	//<- u get msg and store in msg variables
	//create function that cna be exxecute in asynchronous way
	//go permit to send an function in a asynchronous task
	//go is awaiting an listener and when the listener is online accessible in the api
	go func() {
		fmt.Println("2")
		for {
			messages <- "ping" //emitter
		}

		fmt.Println("3")
	}()
	fmt.Println("4")
	msg := <-messages //listener
	fmt.Println("5")
	println(msg)

}
