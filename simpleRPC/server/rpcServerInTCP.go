package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// Args for the RPC method.
// Also all field should be exported type i.e global and name starts with Caps letter.
type Args struct {
	A, B int
}

type AdditionResult struct {
	Sum int
}

type SubtractionResult struct {
	Subtracted int
}

// the Type can be any type: int, float etc
type Arith int

/*
RPC method must satisfy few criteria:
1) Method has exactly two arguments: First parameter will be request args, second parameter will be result args and MUST be pointer.
2) Return type will be error
3) Method is exported i.e name starts with Cap
 */
func (a *Arith) AddImpl(args *Args, result *AdditionResult) error {
	log.Println("Performing addition with args: ", *args)
	sum := args.A + args.B
	result.Sum = sum
	return nil
}

func (a *Arith) SubImpl(args *Args, result *SubtractionResult) error {
	log.Println("Performing subtraction with args: ", *args)
	if args.A < args.B {
		return fmt.Errorf("Wanted to return error: A is lesser than B")
	}
	sub := args.A - args.B
	result.Subtracted = sub

	return nil
}

func main() {
	arith := new(Arith)

	err := rpc.Register(arith)
	if err != nil {
		log.Fatal(err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
		return
	}
    log.Println("Listening on port 1234")
	rpc.Accept(listener)
	/*
	Or can be also used as
	for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        rpc.ServeConn(conn)
    }
	 */
}
