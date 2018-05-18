package main

import ( 
	"fmt"
	"flag"
	"strings"
	"net"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	food ".."

)

func main() {
	op:= flag.String("op","s", "s for server, c for client")
	flag.Parse()

	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}
}

func RunProto3Client() {
	fmt.Println ("hello world c")
	nutrition := &food.FoodTypes {
		Fruit: "mango",
		Vegetable: "kale", 
		Carbs: "noodles",
		Protein: "salmon",
		Desert: "chocolate",
	}

	data, err := proto.Marshal(nutrition)
	if err != nil {
		log.Fatal(err)
	}

	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()
	c.Write(data)
}

func RunProto3Server() {
	fmt.Println ("listening for input")
	l, err :=  net.Listen("tcp",":8282")
	
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		defer l.Close()
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}

			n := &food.FoodTypes{}
			// server deserializes data from client
			err = proto.Unmarshal(data, n)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(n)
		}(c)
	}
}