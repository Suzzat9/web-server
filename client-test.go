// CLIENT SOCKET FOR TESTING SERVER 
// Establish connection to remote host
// Check if any connection error has occured or not
// Send and receive bytes 
// Close the connection

package main 

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
	DIR = "www"
)

// Test for valid index.html request
func sendRequest(wg *sync.WaitGroup, i int) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", i)

	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
			panic(err)
	}
	defer connection.Close()

	// request the index.html file
	_, err = connection.Write([]byte("GET /index.html HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", i)
	
}

// Test for invalid request
func sendRequestInvalid(wg *sync.WaitGroup, i int) {

	defer wg.Done()

	fmt.Printf("Worker %d invalid request starting\n", i)

	//establish connection
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
			panic(err)
	}
	defer connection.Close()

	// request the invalid file
	_, err = connection.Write([]byte("GET /invalid.html HTTP/1.1\r\nHost: localhost\r\n\r\n"))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	time.Sleep(time.Second)

	fmt.Printf("Worker %d invalid request done\n", i)

}

func main() {

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	// Spin up 10 concurrent clients
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go sendRequest(&wg1, i)
	}

	// Send some invalid requests
	for i := 0; i < 2; i++ {
		wg2.Add(1)
		go sendRequestInvalid(&wg2, i)
	}

	wg1.Wait()
	wg2.Wait()
}