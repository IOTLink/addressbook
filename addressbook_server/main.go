package main

import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "addressbook/addressbook"
	. "addressbook/models"
)

const (
	serveraddr = "127.0.0.1:50055"
)

var addressMange AddressBookManage

type server struct{}

func (s *server) SavePerson(ctx context.Context, person *pb.Person) (*pb.MsgReply, error) {
	ret ,err := addressMange.SavePersonToBook(person)
	if ret {
		return &pb.MsgReply{Message:"OK"},nil
	}
	return &pb.MsgReply{Message:"person exist"}, err
}

func (s *server) DelePerson(ctx context.Context,id *pb.Id) (*pb.MsgReply, error) {
	ret, err := addressMange.DelePersonFromBook(id.Id)
	if ret  {
		return &pb.MsgReply{Message:"OK"},nil
	}
	return nil, err
}

func (s *server) GetPersion(ctx context.Context, id *pb.Id) (*pb.Person, error) {
	person := addressMange.QueryPersonFromBook(id.Id)
	return person, nil
}

func (s *server) GetAddrBook(ctx context.Context,limit *pb.Limit) (*pb.AddressBook, error) {
	book := addressMange.QueryAddressBook(limit)
	return book, nil
}

func main() {
	addressMange.Init()
	listen, err := net.Listen("tcp", serveraddr)
	if err != nil {
		log.Fatalf("failed to listen %v\n",err)
	}
	
	s := grpc.NewServer()
	pb.RegisterAddressServerServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server %v",err)
	}
}
