package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"gotham/models"
	"gotham/policies"
	"gotham/requests"
	"gotham/services"
	"gotham/viewModels"
)

type LoanController struct {
	UserService services.IUserService
	LoanService services.ILoanService

	UserPolicy policies.IUserPolicy
}

func (u LoanController) CreateLoan(c echo.Context) (err error) {
	auth := models.ConvertUser(c.Get("auth"))
	// Request Bind And Validation

	request := new(requests.LoanApplyRequest)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, &request.PathParams); err != nil {
		return err
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &request.Body); err != nil {
		return err
	}

	v := request.Validate()
	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	// Policy Control
	if !u.UserPolicy.LoanApply(auth) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	var user models.User
	user, err = u.UserService.GetUserByID(auth.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected - no user"))
	}

	amount, term := getDetails(request)

	loan, err := u.LoanService.CreateLoanByUser(user.ID, amount, term)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, viewModels.MResponse("unexpected error while creating loan"))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(loan))
}


func getDetails(request *requests.LoanApplyRequest)(amount, term uint){

	a, err := strconv.ParseUint(request.Body.Amount, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	t, err := strconv.ParseUint(request.Body.Term, 10, 32)
	if err != nil {
		fmt.Println(err)
	}

	return uint(a), uint(t)
}


func (u LoanController) ApproveLoan(c echo.Context) (err error) {
	auth := models.ConvertUser(c.Get("auth"))
	// Request Bind And Validation

	request := new(requests.ApproveLoanRequet)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, &request.PathParams); err != nil {
		return err
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &request.Body); err != nil {
		return err
	}

	v := request.Validate()
	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	// Policy Control
	if !u.UserPolicy.Index(auth) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	loan, err := u.LoanService.ApproveLoanByAdmin(request.PathParams.Loan)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, viewModels.MResponse("unexpected error while approving loan"))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(loan))
}


func (u LoanController) FetchLoan(c echo.Context) (err error) {
	auth := models.ConvertUser(c.Get("auth"))
	// Request Bind And Validation

	request := new(requests.LoanApplyRequest)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, &request.PathParams); err != nil {
		return err
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &request.Body); err != nil {
		return err
	}

	v := request.Validate()
	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	// Policy Control
	if !u.UserPolicy.LoanApply(auth) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	var user models.User
	user, err = u.UserService.GetUserByID(auth.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected - no user"))
	}

	loan, err := u.LoanService.FetchLoan(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, viewModels.MResponse("unexpected error while creating loan"))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(loan))
}

func (u LoanController) RepayLoan(c echo.Context) (err error){
	auth := models.ConvertUser(c.Get("auth"))
	// Request Bind And Validation

	request := new(requests.RepayLoanRequest)

	if err := (&echo.DefaultBinder{}).BindPathParams(c, &request.PathParams); err != nil {
		return err
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &request.Body); err != nil {
		return err
	}

	v := request.Validate()
	if v != nil {
		return c.JSON(http.StatusUnprocessableEntity, viewModels.ValidationResponse(v))
	}

	// Policy Control
	if !u.UserPolicy.LoanApply(auth) {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected "))
	}

	// Fetch User
	var user models.User
	user, err = u.UserService.GetUserByID(auth.ID)
	if err != nil {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected - no user"))
	}

	// Fetch Loan
	loan, err := u.LoanService.FetchLoanById(request.PathParams.Loan)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, viewModels.MResponse("unexpected error while creating loan"))
	}

	// check user-loan mapping
	if loan.UserId != user.ID {
		return c.JSON(http.StatusForbidden, viewModels.MResponse("unauthorized transaction detected - no user-loan"))
	}

	//RepayLoan

	err = u.LoanService.RepayLoan(loan, request.Body.Amount)
	if err != nil {
		return c.JSON(http.StatusBadRequest, viewModels.MResponse(err.Error()))
	}

	// Response
	return c.JSON(http.StatusOK, viewModels.SuccessResponse(loan))
}