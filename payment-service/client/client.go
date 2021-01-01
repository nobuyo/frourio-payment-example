package main

import (
	"context"
	"fmt"

	gpay "app/payment-service/proto"
	"log"

	"google.golang.org/grpc"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()

	c := gpay.NewPaymentManagerClient(conn)

	req := &gpay.ChargeRequest{
		Token:  "token",
		Amount: 3000,
		Name:   "sample",
		Desc:   "this is test",
	}
	resp, err := c.Charge(context.Background(), req)
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Println(resp)
}
