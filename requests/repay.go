package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RepayLoanRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct {
		Loan uint `param:"loan"`
	}

	/**
	 * QueryParams
	 */
	QueryParams struct{}

	/**
	 * Body
	 */
	Body struct {
		Amount uint `json:"amount" form:"amount" xml:"amount"`
	}
}

func (r RepayLoanRequest) Validate() error {
	return nil
}
