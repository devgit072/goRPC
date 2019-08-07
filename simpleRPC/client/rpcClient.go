package main

import (
	"log"
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

func main() {

	client, err := rpc.Dial("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Error while client creation: ", err.Error())
	}

	args := Args{4,5}

	addResult := new(AdditionResult)
	err = client.Call("Arith.AddImpl", args, addResult)

	if err != nil {
		log.Println("Error while calling rpc method...")
		return
	}
	log.Println("Addition result: ", addResult.Sum)
}
