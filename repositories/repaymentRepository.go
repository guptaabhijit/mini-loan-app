package repositories

import (
	"gotham/infrastructures"
	"gotham/models"
)

type IRepaymentRepository interface {
	Migratable
	Seedable

	// Create & Save & Updates & Delete
	Create(repayment *models.Repayment) (err error)
	Save(repayment *models.Repayment) (err error)
	FetchAllRepayments(loanID uint) (repayments []models.Repayment, err error)
}

type RepaymentRepository struct {
	infrastructures.IGormDatabase
}


func (repository *RepaymentRepository) Seed() (err error) {
	return nil
}

func (repository *RepaymentRepository) Migrate() (err error) {
	return repository.DB().AutoMigrate(models.Repayment{})
}

func (repository *RepaymentRepository) Create(repayment *models.Repayment) (err error) {
	return repository.DB().Create(repayment).Error
}

func (repository *RepaymentRepository) Save(repayment *models.Repayment) (err error) {
	return repository.DB().Save(repayment).Error
}

func (repository *RepaymentRepository) FetchAllRepayments(loanID uint) (repayments []models.Repayment, err error) {

	err = repository.DB().Model(&models.Repayment{}).Where("loan_id = ? AND status = ?", loanID, models.PENDING).Find(&repayments).Error

	return
}
