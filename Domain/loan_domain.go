package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	LoanTakerID     string             `json:"loan_taker" bson:"loan_taker"`
	LoanGiverID     primitive.ObjectID `json:"loan_giver_id" bson:"loan_giver_id"`
	LoanAmount      int64              `json:"loan_amount" bson:"loan_amount"`
	LoanCreatedDate time.Time          `json:"loan_date" bson:"loan_date"`
	LoanUpdatedDate time.Time          `json:"loan_updated_date" bson:"loan_updated_date"`
	LoanStatus      string             `json:"loan_status" bson:"loan_status"`
	LoanType        string             `json:"loan_type" bson:"loan_type"`
	LoanTerm        string             `json:"loan_term" bson:"loan_term"`
	InterestRate    float64            `json:"interest_rate" bson:"interest_rate"`
}

type LoanRepository interface {
	CreateLoan(loan Loan) error
	GetLoans() ([]Loan, error)
	GetLoan(id string) (Loan, error)
	UpdateLoanStatus(id string, loan Loan) (Loan, error)
	DeleteLoan(id string) error
}
