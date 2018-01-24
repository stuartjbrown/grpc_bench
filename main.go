package main

import "log"

func main() {
	var block = make(chan struct{})
	log.Print("Starting GRPC server on port 8888")
	go startGRPC(":8888")

	log.Print("Starting REST server on port 8889")
	go startREST(":8889")
	<-block
}
