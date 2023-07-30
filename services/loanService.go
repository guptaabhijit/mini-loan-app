package services

import (
	"errors"
	"fmt"
	"gotham/models"
	"gotham/repositories"
)

type ILoanService interface {
	FetchLoan(userID uint) (loan []models.Loan, err error)
	FetchLoanById(loanID uint) (loan models.Loan, err error)
	CreateLoanByUser(userID, amount, term uint)  (loan models.Loan, err error)
	ApproveLoanByAdmin(loanID uint) (loan models.Loan, err error)
	RepayLoan(loan models.Loan, amount uint) (err error)
}

type LoanService struct {
	LoanRepository      repositories.ILoanRepository
	RepaymentRepository repositories.IRepaymentRepository
}

func (service *LoanService) CreateLoanByUser(userID, amount, term uint) (loan models.Loan, err error) {
	loan = models.Loan{
		UserId: userID,
		Amount: amount,
		Term:   term,
		Status: models.NEW,
	}

	if err := service.LoanRepository.Create(&loan); err != nil {
		return loan, err
	}

	// create repayment

	repaymentAmount := amount / term

	for t := 0; t < int(term); t++ {
		rA := models.Repayment{
			LoanId: loan.ID,
			Amount: repaymentAmount,
			Status: models.PENDING,
		}

		if err := service.RepaymentRepository.Create(&rA); err != nil {
			fmt.Println(err)
			return loan, err
		}
	}

	return loan, nil
}

func (service *LoanService) ApproveLoanByAdmin(loanID uint) (loan models.Loan, err error) {
	loan, err = service.LoanRepository.GetLoanByID(loanID)
	if err != nil {
		return loan, err
	}

	loan.Status = models.APPROVED

	return loan, service.LoanRepository.Save(&loan)
}

func (service *LoanService) FetchLoan(userID uint) (loan []models.Loan, err error) {

	return service.LoanRepository.FetchLoanByUserId(userID)
}

func (service *LoanService) FetchLoanById( loanID uint) (loan models.Loan, err error) {

	loan, err = service.LoanRepository.GetLoanByID(loanID)
	if err != nil {
		return loan, err
	}

	return loan, nil
}

func (service *LoanService) RepayLoan(loan models.Loan, amount uint) (err error) {

	minAmount := loan.Amount / loan.Term

	if amount < minAmount {
		return errors.New("amount is less for repayment")
	}

	if loan.Status == models.PAID {
		return errors.New("loan is paid already")
	}

	if loan.Status == models.NEW {
		return errors.New("loan needs to be approved")
	}

	if amount%minAmount != 0 {
		return errors.New("amount should be in multiple of repayment")
	}

	// Fetch all repayments of the loan with status = Pending.
	// If no repayment is found, then mark loan as paid

	repayments, err := service.RepaymentRepository.FetchAllRepayments(loan.ID)
	if err != nil {
		return err
	}

	if len(repayments) == 0 {
		loan.Status = models.PAID

		err = service.LoanRepository.Save(&loan)
		if err!= nil {
			return err
		}

		return errors.New("no repayments left")
	}

	currentAmount := amount

	for _, repay := range repayments {

		if currentAmount < 0 {
			break
		}

		if repay.Amount == currentAmount {
			repay.Status = models.PAID_REPAYMENT
		}

		currentAmount -= repay.Amount

		err = service.RepaymentRepository.Save(&repay)
		if err != nil {
			return err
		}
	}

	return
}