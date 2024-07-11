package service

import (
	"context"
	"database/sql"
	"log/slog"
	pb "payments/genproto/payment"
	"payments/pkg/logger"
	"payments/storage/postgres"
)

type NewPaymentService struct {
	pb.UnimplementedPaymentServer
	Payment *postgres.PaymentRepo
	Logger  *slog.Logger
}

func NewPaymentServiceRepo(db *sql.DB) *NewPaymentService {
	return &NewPaymentService{Payment: postgres.NewPaymentRepo(db), Logger: logger.NewLogger()}
}

func (P *NewPaymentService) CreatePayments(ctx context.Context, req *pb.CreatePayment) (*pb.Status, error) {
	resp, err := P.Payment.CreatePayments(req)
	if err != nil {
		P.Logger.Error("Payment create error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (P *NewPaymentService) GetByIdPayments(ctx context.Context, req *pb.GetById) (*pb.GetByIdResponse, error) {
	resp, err := P.Payment.GetPaymentStatusById(req)
	if err != nil {
		P.Logger.Error("Read payment error: %v", err)
		return nil, err
	}
	return resp, nil
}

func (P *NewPaymentService) UpdatePayments(ctx context.Context, req *pb.UpdatePayment) (*pb.Status, error) {
	resp, err := P.Payment.UpdatePayments(req)
	if err != nil {
		P.Logger.Error("Update payment error: %v", err)
		return nil, err
	}
	return resp, nil
}
