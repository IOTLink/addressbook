package models

import (
	"io/ioutil"
	"os"
	"fmt"
	"log"
	"protobuf/proto"
	pb "addressbook/addressbook"
)

const (
	DATAFILE = "./addressbook.dat"
)
type AddressBookManage struct {
	AddrBook *pb.AddressBook
	DBfile   string
}

func (manage *AddressBookManage) Init() {
	manage.AddrBook = new(pb.AddressBook)
	manage.DBfile = DATAFILE
	
	in, err := ioutil.ReadFile(manage.DBfile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Sprintf("%s file not found. create new file.\n",manage.DBfile)
		}else {
			log.Fatalln("error reading file:",err)
		}
	}
	
	if err := proto.Unmarshal(in, manage.AddrBook); err != nil {
		 log.Fatalln("failed to parse address book:", err)
	}
}

func (manage *AddressBookManage) SavePersonToBook(person *pb.Person) (bool, error) {
	for _, v := range manage.AddrBook.People {
		if v.Name == person.Name && v.Id == person.Id {
			return false,nil
		}
	}
	manage.AddrBook.People = append(manage.AddrBook.People, person)
	
	out, err := proto.Marshal(manage.AddrBook )
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(manage.DBfile, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	
	return true,err
}

func (manage *AddressBookManage) DelePersonFromBook(id int32) (bool, error) {
	var myerr error
	for i, v := range manage.AddrBook.People {
		if v.Id == id {
			nextId := i + 1
			manage.AddrBook.People = append(manage.AddrBook.People[:i], manage.AddrBook.People[nextId:]...)
			
			out, err := proto.Marshal(manage.AddrBook )
			if err != nil {
				log.Fatalln("Failed to encode address book:", err)
				myerr = err
			}
			if err := ioutil.WriteFile(manage.DBfile, out, 0644); err != nil {
				log.Fatalln("Failed to write address book:", err)
				myerr = err
			}
			return true,nil
		}
	}
	return false, myerr
}

func (manage *AddressBookManage) QueryPersonFromBook(id int32) *pb.Person {
	for _, v := range manage.AddrBook.People {
		if v.Id == id {
			return v
		}
	}
	return nil
}

func (manage *AddressBookManage) QueryAddressBook(limit *pb.Limit) *pb.AddressBook {
	if int(limit.End - limit.Start) > len(manage.AddrBook.People) {
		return nil
	}
	
	book := new(pb.AddressBook)
	for i, v := range manage.AddrBook.People {
		if int32(i) == limit.Start {
			book.People = append(book.People, v)
		}
		if int32(i) == limit.End {
			break
		}
	}
	return book
}