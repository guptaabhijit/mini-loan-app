// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	models "gotham/models"

	mock "github.com/stretchr/testify/mock"
)

// ILoanRepository is an autogenerated mock type for the ILoanRepository type
type ILoanRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: loan
func (_m *ILoanRepository) Create(loan *models.Loan) error {
	ret := _m.Called(loan)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Loan) error); ok {
		r0 = rf(loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: loan
func (_m *ILoanRepository) Delete(loan *models.Loan) error {
	ret := _m.Called(loan)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Loan) error); ok {
		r0 = rf(loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchLoanByUserId provides a mock function with given fields: userID
func (_m *ILoanRepository) FetchLoanByUserId(userID uint) ([]models.Loan, error) {
	ret := _m.Called(userID)

	var r0 []models.Loan
	if rf, ok := ret.Get(0).(func(uint) []models.Loan); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Loan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLoanByID provides a mock function with given fields: ID
func (_m *ILoanRepository) GetLoanByID(ID uint) (models.Loan, error) {
	ret := _m.Called(ID)

	var r0 models.Loan
	if rf, ok := ret.Get(0).(func(uint) models.Loan); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(models.Loan)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Migrate provides a mock function with given fields:
func (_m *ILoanRepository) Migrate() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: loan
func (_m *ILoanRepository) Save(loan *models.Loan) error {
	ret := _m.Called(loan)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Loan) error); ok {
		r0 = rf(loan)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Seed provides a mock function with given fields:
func (_m *ILoanRepository) Seed() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Updates provides a mock function with given fields: loan, updates
func (_m *ILoanRepository) Updates(loan *models.Loan, updates map[string]interface{}) error {
	ret := _m.Called(loan, updates)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Loan, map[string]interface{}) error); ok {
		r0 = rf(loan, updates)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
