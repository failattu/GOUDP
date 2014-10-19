package main

import (
	"fmt"
	"net"
	"os"
	"time"
)
func server(){
	addr := net.UDPAddr{
        Port: 5055,
        IP: net.ParseIP("127.0.0.1"),
    }
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("error listening on UDP port ")
		fmt.Println(err)
		return
	}
	defer conn.Close()
	var buf []byte = make([]byte, 1500)

	for {

		time.Sleep(100 * time.Millisecond)

		n,address, err := conn.ReadFromUDP(buf)

		if err != nil {
				fmt.Println("error reading data from connection")
				fmt.Println(err)
				return
		}
		if address != nil {

			fmt.Println("got message from ", address, " with n = ", n)

			if n > 0 {
					fmt.Println("from address", address, "got message:", string(buf[0:n]), n)
			}
		}
	}
}

func client(){
	conn, err := net.Dial("udp", "127.0.0.1:5055")
	if err != nil {
		fmt.Print(err)
	}
	  defer conn.Close()
	for {

		time.Sleep(1000*time.Millisecond)
		n, err := conn.Write([]byte("SOS ... \n"))
		if err != nil {
			fmt.Println("error writing data to server")
			fmt.Println(err)
			return
		}

		if n > 0 {
				fmt.Println("Wrote ",n, " bytes to server at ")
		}
        }
}

func main() {
	fmt.Printf("hello, world\n")
	
	if len(os.Args) > 1 {
		if os.Args[1] == "server" {
			server()
		}else if os.Args[1] == "client"{
			client()
		}else{
			fmt.Printf("Goodbye\n")
		}
	}
}