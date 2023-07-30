package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoanApplyRequest struct {
	validation.Validatable `json:"-" form:"-" query:"-"`

	/**
	 * PathParams
	 */
	PathParams struct {}

	/**
	 * QueryParams
	 */
	QueryParams struct{}

	/**
	 * Body
	 */
	Body struct{
		Amount    string `json:"amount" form:"amount" xml:"amount"`
		Term string `json:"term" form:"term" xml:"term"`
	}
}

func (r LoanApplyRequest) Validate() error {
	return nil
}
