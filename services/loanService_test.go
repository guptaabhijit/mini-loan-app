package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gotham/models"
	"gotham/repositories/mocks"
	"os"
	"testing"
)

func Test_FetchLoan(t *testing.T){
	os.Setenv("CI", "test")

	fmt.Println(os.Getenv("CI"))
	mockLoanRepo := &mocks.ILoanRepository{}

	var loans []models.Loan

	mockLoanRepo.On("FetchLoanByUserId", mock.Anything).Return(loans, nil)

	ls := &LoanService{
		LoanRepository: mockLoanRepo,
	}

	user := models.User{ID: 123}

	loans, err := ls.FetchLoan(user.ID)

	assert.NoError(t, err)
}