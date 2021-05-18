package main

import (
	"bufio"
	"fmt"
	"library_java2go/netme/serverSocket"
	"strings"
)

func main() {

	tcpListener, err := serverSocket.NewWithPort(9999)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tcpListener.Close()

	c := tcpListener.Accept()
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))

		myTime := "Receive message : " + string(netData)
		c.Write([]byte(myTime))
	}
}
