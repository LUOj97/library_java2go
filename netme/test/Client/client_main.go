package main

import (
	"bufio"
	"fmt"
	"library_java2go/netme/socket"
	"os"
	"strings"
)

func main() {
	socket, err := socket.NewWithHostPort("", 9999)
	defer socket.Close() //close connection
	fmt.Println("Connected!")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(socket, text+"\n")

		message, _ := bufio.NewReader(socket).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
		}
	}
}
