package repositories

import (
	"gotham/infrastructures"
	"gotham/models"
)

type ILoanRepository interface {
	Migratable
	Seedable

	// Create & Save & Updates & Delete
	Create(loan *models.Loan) (err error)
	Save(loan *models.Loan) (err error)
	Updates(loan *models.Loan, updates map[string]interface{}) (err error)
	Delete(loan *models.Loan) (err error)

	GetLoanByID(ID uint) (loan models.Loan, err error)
	FetchLoanByUserId(userID uint) (loan []models.Loan, err error)

}

type LoanRepository struct {
	infrastructures.IGormDatabase
}

/**
 * Seed
 *
 * @return error
 */
func (repository *LoanRepository) Seed() (err error) {
	return nil
}

/**
 * Migrate
 *
 * @return error
 */

func (repository *LoanRepository) Migrate() (err error) {
	return repository.DB().AutoMigrate(models.Loan{})
}

/**
 * Create & Update & Delete
 *
 */

func (repository *LoanRepository) Create(loan *models.Loan) (err error) {
	return repository.DB().Create(loan).Error
}

func (repository *LoanRepository) Save(loan *models.Loan) (err error) {
	return repository.DB().Save(loan).Error
}

func (repository *LoanRepository) Updates(loan *models.Loan, updates map[string]interface{}) (err error) {
	return repository.DB().Model(loan).Updates(updates).Error
}

func (repository *LoanRepository) Delete(loan *models.Loan) (err error) {
	return repository.DB().Delete(loan).Error
}


func (repository *LoanRepository) GetLoanByID(ID uint) (loan models.Loan, err error) {
	err = repository.DB().First(&loan, ID).Error
	return
}

func (repository *LoanRepository) FetchLoanByUserId(ID uint) (loans []models.Loan, err error) {

	err = repository.DB().Model(&models.Loan{}).Where("user_id = ?", ID).Find(&loans).Error
	//db.Model(&User{}).Where("uac.user_id = ?", "userId").Count(&count)

	return
}