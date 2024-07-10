package service

import (
	"context"
	"database/sql"
	"log"
	pb "payments/genproto/payment"
)

type NewPaymentService struct {
	pb.UnimplementedPaymentServer
	Payment *postgres.NewPayment
}

func NewPaymentServiceRepo(db *sql.DB) *NewPaymentService {
	return &NewPaymentService{Payment: postgres.NewPaymentRepo(db)}
}

func (P *NewPaymentService) CreatePayments(ctx context.Context, req *pb.CreatePayment) (*pb.Status, error) {
	resp, err := postgres.CreatePayments(req)
	if err != nil {
		log.Fatalf("Payment create error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (P *NewPaymentService) GetByIdPayments(ctx context.Context, req *pb.GetById) (*pb.GetByIdResponse, error) {
	resp, err := postgres.GetByIdPayments(req)
	if err != nil {
		log.Fatalf("Read payment error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (P *NewPaymentService) UpdatePayments(ctx context.Context, req *pb.UpdatePayment) (*pb.Status, error) {
	resp, err := postgres.UpdatePayments(req)
	if err != nil {
		log.Fatalf("Update payment error: %v", err)
		return nil, err
	}
	return resp, nil
}
