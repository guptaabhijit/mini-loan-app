package requests


import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type ApproveLoanRequet struct {
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
	Body struct{}
}

func (r ApproveLoanRequet) Validate() error {
	return nil
}
