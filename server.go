// SERVER SOCKET 
// Create a socket on a specific port 
// Listen for connections to that port 
// If the connection is successful, communication can begin between client and server 
// The communication should be according to agreed protocol
// Continue listening on the port 
// Once the connection is closed, server stops listening and exits 

package main
import (
	"fmt"
	"net" // for socket programming
	"os"
	"strings"
	"io/ioutil"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "80"
	SERVER_TYPE = "tcp"
	DIR = "www"
)

func processClient(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))

	// Parse the request to extract the requested path
	requestLine := strings.Split(string(buffer[:mLen]), "\n")[0]
	requestParts := strings.Fields(requestLine)
	requestedPath := requestParts[1]


	// If requested path is "/" or "/index.html", serve it, else send 403
	if requestedPath == "/" || requestedPath == "/index.html" {
		requestedPath = DIR + "/index.html"
	} else {
		response := "HTTP/1.1 403 Forbidden\r\nContent-Type: text/plain\r\n\r\nAccess Forbidden\r\n"
		connection.Write([]byte(response))
		connection.Close()
		return
	}

	// Read and serve index.html
	fileContent, err := ioutil.ReadFile(requestedPath)
	if err != nil {
		response := "HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\n\r\nFile Not Found\r\n"
		connection.Write([]byte(response))
		connection.Close()
		return
	}
	
	// Send response with file contents
	response := "HTTP/1.1 200 OK\r\nContent-Type: text/html\r\n\r\n" + string(fileContent) + "\r\n"
	_, err = connection.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
	}

}

func main() {
	fmt.Println("Server running...")

	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go processClient(connection)
	}
}