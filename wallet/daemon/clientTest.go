package main

import (
	"context"
	"fmt"
	pb "gRPC/assembly/device"
	"log"
	"net/http"
)

func main() {

	client := pb.NewDeviceJSONClient("http://localhost:9000", &http.Client{})

	hat, err := client.Registration(context.Background(), &pb.RequestDeviceRegistration{Phone: 123123, Devicename: "Testname"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("I have a nice new hat: %+v", hat)
}
