//SERVER
//PROCESS B
//go run processb.go

package main

import (
        "fmt"
        "net"
        "encoding/gob"
)
type message struct{
	to 		string
	from 	string	
	date   	string
	title	string
	content	string
}


func main() {
  fmt.Println("start");
  l, err := net.Listen("tcp", ":1234")
  if err != nil {
          fmt.Println(err)
          return
  }
  for {
    conn, err := l.Accept() // this blocks until connection or error
    if err != nil {
        // handle error
        continue
    }
    //msg := new(message)
    //dec.Decode(msg)
    //fmt.Printf("Received : %+v", msg);
    //conn.Close()

    dec := gob.NewDecoder(conn)
    var msg message

    err = dec.Decode(&msg)
    fmt.Println(msg, err)

    conn.Close()
    fmt.Println("ACK")
   // go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
  }
}

//func handleConnection(conn net.Conn) {

//}

    