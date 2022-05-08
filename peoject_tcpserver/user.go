package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	Chan   chan string
	conn   net.Conn
	server *Server
}

// create user
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		Chan:   make(chan string),
		conn:   conn,
		server: server,
	}

	go user.ListenMsg()

	return user
}

// listen user channel
func (this *User) ListenMsg() {
	for {
		// get msg from channel
		msg := <-this.Chan
		// print to console
		this.conn.Write([]byte(msg + "\r\n"))
	}
}

func (this *User) Online() {}

func (this *User) Offline() {}

// server send msg to this user
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg + "\r\n"))
}

func (this *User) DoMsg(msg string) {
	if msg[:1] == "#" {
		if msg == "#list" { // list all online user
			this.server.mapLock.Lock()
			for _, user := range this.server.OnlineMap {
				user_msg := fmt.Sprintf("[%s]%s", user.Addr, user.Name)
				this.SendMsg(user_msg)
			}
			this.server.mapLock.Unlock()
		} else if len(msg) > 8 && msg[:8] == "#rename|" { // private chat
			// msg format: #rename|newName

			newName := msg[8:] // newName := strings.Split(msg, "|")[1]

			// does the user name exist?
			_, ok := this.server.OnlineMap[newName]
			if ok {
				this.SendMsg("user name had been used")
			} else {
				// remove old key and append new key
				this.server.mapLock.Lock()
				delete(this.server.OnlineMap, this.Name)
				this.server.OnlineMap[newName] = this
				this.server.mapLock.Unlock()
				// rename user
				this.Name = newName
				this.SendMsg("username has been updated:" + newName)
			}
		} else if len(msg) > 4 && msg[:4] == "#to|" {
			// msg format: #rename|destName|content

			destName := strings.Split(msg, "|")[1]
			destUser, ok := this.server.OnlineMap[destName]
			if !ok {
				this.SendMsg("user name doesn't exist")
				return
			}

			conntent := strings.Split(msg, "|")[2]
			if conntent == "" {
				this.SendMsg("msg is null")
				return
			}

			destUser.SendMsg(conntent)

		}
		return
	}

	// send msg to other users
	this.server.BoardCast(this, msg)

}
