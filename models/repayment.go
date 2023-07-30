package models

type Repayment struct {
	ID     uint `gorm:"primaryKey;auto_increment" json:"id"`
	LoanId uint
	Loan   Loan `gorm:"foreignKey:LoanId" json:"-"`
	Amount uint
	Status repaymentStatus `gorm:"type:enum('PENDING', 'PAID_REPAYMENT')";"column:status" json:"status"`
}

func (Repayment) TableName() string {
	return "repayment"
}

type repaymentStatus string

const (
	PENDING        repaymentStatus = "PENDING"
	PAID_REPAYMENT repaymentStatus = "PAID_REPAYMENT"
)
