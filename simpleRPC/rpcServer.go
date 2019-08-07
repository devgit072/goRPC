package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	A, B int
}

type AdditionResult struct {
	Sum int
}

type SubtractionResult struct {
	Subtracted int
}

type Arith int

func (a *Arith) AddImpl(args *Args, result *AdditionResult) error {
	log.Println("Performing addition with args: ", *args)
	sum := args.A + args.B
	result.Sum = sum
	return nil
}

func (a *Arith) SubImpl(args *Args, result *SubtractionResult) error {
	log.Println("Performing subtraction with args: ", *args)
	if args.A > args.B {
		return fmt.Errorf("Wanted to return error: A is greater than B")
	}
	sub := args.A - args.B
	result.Subtracted = sub

	return nil
}

func main() {
	arith := new(Arith)

	rpc.Register(arith)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
    log.Println("Listening on port 1234")
	rpc.Accept(listener)
}
