package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

import (
	"github.com/tkxkd0159/buf-proto/grpc/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	person := &pb.Person{
		Id:    1234,
		Name:  "JS Lee",
		Email: "no-reply@gihub.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "123-4567", Type: pb.Person_HOME},
		},
		LastUpdated: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   int32(time.Now().Nanosecond()),
		},
	}
	out, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("Failed to encode person:", err)
	}

	fname := "proto-writter"
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error occured while open the file", err)
	}
	defer SafeClose(f)

	if _, err := f.Write(out); err != nil {
		log.Fatalln("Error occured while write a marshaled data", err)
	}

	f, err = os.Open(fname)
	ShowErr(err)
	defer SafeClose(f)

	b, err := io.ReadAll(f)
	person2 := &pb.Person{}
	if err := proto.Unmarshal(b, person2); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(person2)

}

func SafeClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalln("Error: while file is closed", err)
	}
}

func ShowErr(err error) {
	if err != nil {
		log.Printf("%#v", err)
	}
}
