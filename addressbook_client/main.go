package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "addressbook/addressbook"
	"encoding/json"
)

const (
	serveraddr = "127.0.0.1:50055"
)

func main(){
	conn, err := grpc.Dial(serveraddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
	c := pb.NewAddressServerClient(conn)
	
	person := new(pb.Person)
	person.Name =  "lhy"
	person.Id = 1
	person.Email = "liuhangyu@163.com"
	phone := pb.Person_PhoneNumber{Number:"110", Type: 0}
	person.Phones =  append(person.Phones, &phone)
	r, err := c.SavePerson(context.Background(), person)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("save person: %s", r.Message)
	
	//return
	//______________________________________________________________-
	
	person1 := new(pb.Person)
	person1.Name =  "lhy2"
	person1.Id = 2
	person1.Email = "liuhangyu@163.com"
	phone1 := pb.Person_PhoneNumber{Number:"1190", Type: 0}
	person1.Phones =  append(person1.Phones, &phone1)
	r1, err := c.SavePerson(context.Background(), person1)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("save person: %s", r1.Message)
	//______________________________________________________________-
	
	var id pb.Id
	id.Id = 1
	myperson, err := c.GetPersion(context.Background(), &id)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	sperson, _:= json.Marshal(myperson)
	log.Printf("person: %s", sperson)
	//______________________________________________________________
	
	var nid pb.Id
	nid.Id = 1
	ret, err := c.DelePerson(context.Background(), &nid)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("dele person: %s", ret.Message)
	//______________________________________________________________
	
	var lim pb.Limit
	lim.Start = 0
	lim.End = 1
	book, err := c.GetAddrBook(context.Background(), &lim)
	if err != nil {
		log.Fatalf("could not GetAddrBook: %v", err)
	}
	sbook, _:= json.Marshal(book)
	log.Printf("person: %s", sbook)
	
}

/*
go build -o server addressbook_server/main.go
go build -o client addressbook_client/main.go
 */