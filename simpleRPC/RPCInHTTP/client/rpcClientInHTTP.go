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

	client, err := rpc.DialHTTP("tcp", "localhost:1234")

	if err != nil {
		log.Fatal("Error while client creation: ", err.Error())
	}

	callAdd(client)
	callSubstraction(client)
	defer client.Close()
}

func callAdd(client *rpc.Client) {
	args := Args{4,5}

	addResult := new(AdditionResult)
	err := client.Call("Arith.AddImpl", args, addResult)

	if err != nil {
		log.Println("Error while calling rpc method Add: ", err.Error())
		return
	}
	log.Println("Addition result: ", addResult.Sum)
}

func callSubstraction(client *rpc.Client) {
	args := Args{17,5}

	subResult := new(SubtractionResult)
	err := client.Call("Arith.SubImpl", args, subResult)

	if err != nil {
		log.Println("Error while calling rpc method subtraction: ", err.Error())
		return
	}
	log.Println("Substraction Result result: ", subResult.Subtracted)
}
