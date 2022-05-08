package main // the same package name doesn't need to be imported

import "fmt"

// window go build -o server.exe main.go server.go client.go
// linux go build -o server.exe main.go server.go client.go

func main() {
	fmt.Println("server.start")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

/*

server channel *1
user channel *n

when server receive msg, it will be put into server channel,
then the server listener will get it form server channel,
and put it into all user channels.

when user channel receive msg, it will be written in connection.
then the connection send it to client. (conn.Write)

when user connection receive msg from client, read it. (conn.Read)


*/
