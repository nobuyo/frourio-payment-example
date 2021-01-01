package main

import (
	"context"
	"log"
	"net"
	"os"

	gpay "app/payment-service/proto"

	payjp "github.com/payjp/payjp-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Charge(ctx context.Context, req *gpay.ChargeRequest) (*gpay.ChargeResponse, error) {
	pay := payjp.New(os.Getenv("PAYJP_SECRET_KEY"), nil)

	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token,
		Capture:     false,
		Description: req.Name + " " + req.Desc,
	})
	if err != nil {
		return nil, err
	}

	res := &gpay.ChargeResponse{
		Id:       charge.ID,
		Paid:     charge.Paid,
		Refunded: charge.Refunded,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}

	return res, nil
}

func (s *server) Capture(ctx context.Context, req *gpay.ChargeRequest) (*gpay.ChargeResponse, error) {
	pay := payjp.New(os.Getenv("PAYJP_SECRET_KEY"), nil)

	charge, err := pay.Charge.Capture(req.ChargeId)
	if err != nil {
		return nil, err
	}

	res := &gpay.ChargeResponse{
		Id:       charge.ID,
		Paid:     charge.Paid,
		Refunded: charge.Refunded,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}

	return res, nil
}

func (s *server) Refund(ctx context.Context, req *gpay.ChargeRequest) (*gpay.ChargeResponse, error) {
	pay := payjp.New(os.Getenv("PAYJP_SECRET_KEY"), nil)

	charge, err := pay.Charge.Refund(req.ChargeId, "")
	if err != nil {
		return nil, err
	}

	res := &gpay.ChargeResponse{
		Id:       charge.ID,
		Paid:     charge.Paid,
		Refunded: charge.Refunded,
		Captured: charge.Captured,
		Amount:   int64(charge.Amount),
	}

	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	gpay.RegisterPaymentManagerServer(s, &server{})

	reflection.Register(s)
	log.Printf("gRPC server started: localhost%s\n", port)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %w", err)
	}
}
