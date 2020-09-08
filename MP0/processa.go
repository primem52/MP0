//MP0client.go
//PROCESS A

//go run processa.go 127.0.0.1
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
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
			fmt.Println(err)
			return
	}
	var email 	[]message
	var to 		string
	var from 	string	
	var date   	string
	var title	string
	var content	string

	//takes command line input
	fmt.Print("To: ")
	fmt.Scanln(&to)
	fmt.Print("From: ")
	fmt.Scanln(&from)
	fmt.Print("Date: ")
	fmt.Scanln(&date)
	fmt.Print("Subject: ")
	fmt.Scanln(&title)
	fmt.Print("Content: ")
	fmt.Scanln(&content)

	//creates email 
	email = append(email, message{to,from,date,title,content})
	
	//email = append(email, message{to,from,date,title,content})
	//encoder := gob.NewEncoder(conn)
	//encoder.Encode(email)

    defer conn.Close()
	enc := gob.NewEncoder(conn)
	enc.Encode(email)
	
}

