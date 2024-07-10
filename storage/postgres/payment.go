package postgres

import (
	"database/sql"
	"log"
	pb "payments/genproto/payment"
	"time"

	"github.com/google/uuid"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo{
	return &PaymentRepo{db: db}
}

func (p *PaymentRepo) CreatePayments(req *pb.CreatePayment) (*pb.Status,error){
	query:=`
		insert into Payments(
			id, reservation_id, amount, payment_method, payment_status, created_at, update_at
		)values(
			$1,$2,$3,$4,$5,$6,$7
		)
	`
	id:=uuid.NewString()
	newtime:=time.Now()
	_,err:=p.db.Exec(query,id,req.ReservationId,req.Amount,req.Paymentmethod,req.Paymentstatus,newtime,newtime)
	if err!=nil{
		log.Fatalf("Error inserting data: %v",err)
		return nil,err
	}
	return &pb.Status{Message: "Data has been added accordingly",Status: true},nil
}

func (p *PaymentRepo) GetPaymentStatusById(req *pb.GetById)(*pb.GetByIdResponse,error){
	resp:=pb.GetByIdResponse{}
	query:=`
		select 
			payment_status
		from 
			Payments 
		where 
			id=$1
	`
	err:=p.db.QueryRow(query,req.Id).Scan(resp.Paymentstatus)
	if err!=nil{
		log.Fatalf("Error getting data: %v",err)
		return nil,err
	}
	return &resp,nil
}

func (p *PaymentRepo) UpdatePayments(req *pb.UpdatePayment)(*pb.Status,error){
	query:=`
		update 
			Payments 
		set
			amount=$1,
			payment_method=$2,
			payment_status=$3
		where 
			id=$4 and
			deleted_at is null
	`

	_,err:=p.db.Exec(query,req.Amount,req.PaymentMethod,req.PaymentStatus,req.Id)
	if err!=nil{
		log.Fatalf("Error updating data: %v",err)
		return nil,err
	}
	return &pb.Status{Message: "Your information has been updated",Status: true},nil
}