package main

import (
	"bufio"
	"flag"
	"fmt"
	"goconcurrency/internal/network"
	"log"
	"os"
)

var (
	defaultAddress = "localhost:3223"
	maxBufferSize  = 1024
)

func main() {
	address := flag.String("address", defaultAddress, "server address")
	maxBufferSize := flag.Int("max-buffer-size", maxBufferSize, "max buffer size")
	flag.Parse()

	cc, err := network.NewTcpConnection(*address, uint(*maxBufferSize))
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(os.Stdin)
	for {
		req, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		resp, err := cc.Send([]byte(req))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(fmt.Sprintf("%s", string(resp)))
	}

}
