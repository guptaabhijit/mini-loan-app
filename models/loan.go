package models

type Loan struct {
	ID     uint `gorm:"primaryKey;auto_increment" json:"id"`
	UserId uint
	User   User `gorm:"foreignKey:UserId" json:"-"`
	Amount uint
	Term   uint
	Status loanStatus `gorm:"type:enum('NEW', 'APPROVED', 'REJECTED','PAID')";"column:status" json:"status"`
}

func (Loan) TableName() string {
	return "loan"
}

type loanStatus string

const (
	NEW      loanStatus = "NEW"
	APPROVED loanStatus = "APPROVED"
	PAID     loanStatus = "PAID"
)
