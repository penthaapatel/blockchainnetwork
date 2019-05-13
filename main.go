package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"

	b "github.com/penthaapatel/blockchainnetwork/blocks"
)

func main() {
	portFlag := flag.String("port", "8080", "Port to run tcp server on")
	flag.Parse()
	port := ":" + *portFlag
	b.Blockchain = append(b.Blockchain, b.GenesisBlock)
	//Create a tcp server
	fmt.Println("Starting tcp server on localhost" + port)
	server, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		fmt.Println("Closing server. . .")
		server.Close()
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	fmt.Println(c.RemoteAddr().String() + " joined.")
	defer func() {
		fmt.Println(c.RemoteAddr().String() + " says bye!")
		c.Close()
	}()

	scanner := bufio.NewScanner(c)
	c.Write([]byte("Enter input data for new Block generation: "))

	for scanner.Scan() {
		//Prints on server terminal
		newData := scanner.Text()
		newBlock := b.GenerateBlock(b.Blockchain[len(b.Blockchain)-1], newData)
		if b.CheckBlock(b.Blockchain[len(b.Blockchain)-1], newBlock) {
			fmt.Println("NEW BLOCK CREATED")
			output, err := b.OutputJSON(newBlock)
			if err != nil {
				fmt.Println("Failed to render json output")
			}
			fmt.Println(output)
			newBlockchain := append(b.Blockchain, newBlock)
			b.GenerateChain(newBlockchain)
		}
		c.Write([]byte("Enter input data for new Block generation: "))

	}
}
