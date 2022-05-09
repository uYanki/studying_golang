package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	IP   string // server
	Port int    // server
	Name string // client
	conn net.Conn
	flag int
}

func NewClient(serverIP string, serverPort int) *Client {
	client := &Client{
		IP:   serverIP,
		Port: serverPort,
		flag: 999,
	}

	// connect tcp server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", client.IP, client.Port))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}

	client.conn = conn

	return client

}

func (this *Client) menu() bool {
	var flag int
	fmt.Println("1.Public chat mode")
	fmt.Println("2.Private chat mode")
	fmt.Println("3.Rename")
	fmt.Println("0.Exit")

	fmt.Scan(&flag)

	if 0 <= flag && flag <= 4 {
		this.flag = flag
		return true
	} else {
		fmt.Println("illegal input")
		return false
	}
}

func (this *Client) pubilcChat() {
	var chatMsg string
	fmt.Println("Please enter something you want to send (enter '#menu' open menu)")
	fmt.Scanln(&chatMsg)
	fmt.Scanln(&chatMsg)
	for chatMsg != "#menu" {
		if len(chatMsg) != 0 {
			_, err := this.conn.Write([]byte(chatMsg))
			if err != nil {
				fmt.Println("conn.Write err: ", err)
				break
			}
		}
		chatMsg = ""
		fmt.Println("Please enter something you want to send (enter '#menu' to open menu)")
		fmt.Scanln(&chatMsg) // not support space
	}
}

func (this *Client) privateChat() {

	var userName string
	var chatMsg string

	this.listuser()
	fmt.Println("Please enter target user name (enter '#menu' to menu)")
	fmt.Scanln(&userName)
	fmt.Scanln(&userName)

	for userName != "#menu" {
		fmt.Println("Please enter  something you want to send (enter '#menu' to open menu)")
		fmt.Scanln(&chatMsg)

		for chatMsg != "#menu" {
			if len(chatMsg) != 0 {
				sendMsg := "#to|" + userName + "|" + chatMsg
				_, err := this.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write err: ", err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("Please enter something you want to send (enter '#menu' to open menu)")
			fmt.Scanln(&chatMsg) // not support space
		}

		// select
		this.listuser()
		fmt.Println("Please enter target user name (enter '#menu' to menu)")
		fmt.Scanln(&userName)
	}

}

func (this *Client) rename() bool {
	fmt.Println("Please enter a new user name: ")
	fmt.Scanln(&this.Name) // 不知道为什么, window 这里要写两行
	fmt.Scanln(&this.Name)
	sendMsg := "#rename|" + this.Name
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}

func (this *Client) listuser() {
	sendMsg := "#list"
	_, err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}
}

func (this *Client) run() {
	for this.flag != 0 {
		// mode select
		for this.menu() != true {
		}

		switch this.flag {
		case 1:
			this.pubilcChat()
			break
		case 2:
			this.privateChat()
			break
		case 3:
			this.rename()
			break
		}
	}
}

func (this *Client) recv() {
	io.Copy(os.Stdout, this.conn)
	/* // same as the following
	for {
		buf := make([]byte, 2048)
		this.conn.Read(buf)
		fmt.Println(buf)
	}
	*/
}

var serverIP string
var serverPort int

func init() {
	// window format: .\client.exe -ip 127.0.0.1 -port 8888
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "server ip (default: 127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "server port (default: 8888)")
}

func main() {

	// parse command line
	flag.Parse()

	// connect server
	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println("--- fail to connect server --")
		return
	}
	fmt.Println("--- success to connect server --")

	// start recv
	go client.recv()

	// run
	client.run()
}
