package main

import (
	"context"
	"fmt"
	"os"
	"time"

	proto "github.com/konjoot/proto-errors/proto"
	service "github.com/konjoot/proto-errors/service"
	micro "github.com/micro/go-micro"
	errors "github.com/micro/go-micro/errors"
)

func main() {
	started := make(chan struct{})
	s1 := micro.NewService(
		micro.Name("service"),
		micro.RegisterTTL(time.Second),
		micro.AfterStart(func() error {
			close(started)
			return nil
		}),
	)

	// Init will parse the command line flags.
	s1.Init()

	// Register handler
	proto.RegisterServiceHandler(s1.Server(), new(service.Service))

	go func() {
		// Run the server
		if err := s1.Run(); err != nil {
			fmt.Println(err)
		}
	}()
	<-started

	// Create a new service. Optionally include some options here.
	s2 := micro.NewService(micro.Name("client"))
	s2.Init()

	// Create a new client
	client := proto.NewService("service", s2.Client())

	// Let's rock!
	fmt.Println()
	fmt.Println("Success")
	fmt.Println("#######")
	rsp, err := client.CreateThingOneOf(
		context.TODO(),
		&proto.CreateThingOneOfRequest{Name: "success"},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if e := rsp.GetError(); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	if t := rsp.GetThing(); t != nil {
		fmt.Printf("%#v\n", t)
	}

	fmt.Println()
	fmt.Println("Business Error")
	fmt.Println("##############")
	rsp, err = client.CreateThingOneOf(
		context.TODO(),
		&proto.CreateThingOneOfRequest{Name: "business-error"},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if e := rsp.GetError(); e != nil {
		fmt.Printf("%#v\n", e)
	}
	if t := rsp.GetThing(); t != nil {
		fmt.Println(t)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Transport Error")
	fmt.Println("###############")
	rsp, err = client.CreateThingOneOf(
		context.TODO(),
		&proto.CreateThingOneOfRequest{Name: "transport-error"},
	)
	if err != nil {
		fmt.Printf("raw transport error: %#v\n", err)
		e := errors.Parse(err.Error())
		fmt.Printf("parsed transport error: %#v\n", e)
	}
	if e := rsp.GetError(); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	if t := rsp.GetThing(); t != nil {
		fmt.Println(t)
		os.Exit(1)
	}
}
