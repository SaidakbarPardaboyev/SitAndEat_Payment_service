package postgres

import (
	pb "payments/genproto/payment"
	"testing"

	_ "github.com/lib/pq"
)

func TestPaymentRepo(t *testing.T) {
	db, err := ConnectionDB()

	if err != nil {
		t.Fatal(err)
	}

	paymentservice := NewPaymentRepo(db)

	status,err := paymentservice.CreatePayments(&pb.CreatePayment{
		ReservationId: "2852b719-f2a6-4fa4-80fc-8d46a4b96cff",
		Amount:        100,
		Paymentmethod: "card",
		Paymentstatus: "FAILED",
	})
	if status.Message != "Data has been added accordingly" {
		t.Fatal(err)
	}
}

func TestGetPaymentStatusById(t *testing.T) {
	db, err := ConnectionDB()
	if err != nil {
		t.Fatal(err)
	}
	paymentservice := NewPaymentRepo(db)

	status, err := paymentservice.GetPaymentStatusById(&pb.GetById{
		Id: "98bf6060-06c9-4e55-93d1-784a82af71ce",
	})
	if status.Paymentstatus != "FAILED" {
		t.Fatal(err)
	}
}


func TestUpdatePayment(t *testing.T) {
	db, err := ConnectionDB()
	if err != nil {
		t.Fatal(err)
	}
	paymentservice := NewPaymentRepo(db)

	status, err := paymentservice.UpdatePayments(&pb.UpdatePayment{
		Id: "98bf6060-06c9-4e55-93d1-784a82af71ce",
		Amount: 100,
		PaymentMethod: "Cash",
		PaymentStatus: "COMPLETED",
	})
	if status.Message != "Your information has been updated" {
		t.Fatal(err)
	}
}