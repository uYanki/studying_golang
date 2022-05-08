package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// conn - connection
// chan - channel
// cli  - client
// recv - receive

type Server struct {
	IP   string
	Port int

	// online user's list
	OnlineMap map[string]*User
	mapLock   sync.RWMutex // 读写锁

	// broadcast channel
	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// listen msg, when it get msg, send to all user
func (this *Server) ListenMsg() {
	for {
		// get msg from channel
		msg := <-this.Message

		this.mapLock.Lock()
		// boardcast msg to all user
		for _, cli := range this.OnlineMap {
			cli.Chan <- msg
		}
		this.mapLock.Unlock()
	}
}

// boardcast msg
func (this *Server) BoardCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	this.Message <- sendMsg
}

// do user msg
func (this *Server) Handle(conn net.Conn) {
	// add user to onlineMap
	user := NewUser(conn, this)
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// boardcast user login
	this.BoardCast(user, "login")
	fmt.Println("[ user login ]", user.Name)

	// user do something when login
	user.Online()

	// whether user send a msg in 60 second
	// if not, kick out the user
	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096) //4KB
		for {
			// receive msg from user
			n, err := conn.Read(buf)
			if n == 0 {
				// user do something when logout
				user.Offline()
				// remove user from onlineMap
				this.mapLock.Lock()
				delete(this.OnlineMap, user.Name)
				this.mapLock.Unlock()
				// boardcast user logout
				this.BoardCast(user, "logout")
				fmt.Println("[ user logout ]", user.Name)
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn.Read err:", err)
				return
			}

			isLive <- true

			// boardcast msg which from user
			msg := string(buf[:n])

			// user do something when recv msg
			user.DoMsg(msg)

		}
	}()

	for {
		select {

		case <-isLive:
			// nothing

		case <-time.After(time.Second * 60):
			// no msg sent in 60s, kick user
			user.SendMsg("you are forced to close the connection")
			close(user.Chan)
			// disconnect
			conn.Close()
			return // runtime.Goexit()

		}
	}

}

func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.IP, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// close socket listen
	defer listener.Close()

	// listen msg
	go this.ListenMsg()

	for {
		// accept connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			continue
		}
		// handle user login
		go this.Handle(conn)
	}
}
