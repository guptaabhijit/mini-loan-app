# mini-loan-app
Mini Loan App consisting of APIs written in golang
This application uses go-gotham framework: https://github.com/tolgaOzen/go-gotham which gives boilerplate for creating RESTful APIs adhering to SOLID principles.

## Env
create a new environment file `.env` from `.env.example`
Update necessary Database configurations.

## Flags

- prevents re-creating container methods from definitions
```
go run gotham -production
```

- run migrate methods of repositories before start
```
go run gotham -migrate
```

- run seed methods of repositories before start
```
go run gotham -seed
```

## Database Schema

3 Tables/Relations

Users, Loan, Repayment

![Screenshot 2023-07-30 at 9 05 20 PM](https://github.com/guptaabhijit/mini-loan-app/assets/7620035/b0cc49ed-fa9c-408c-b12f-eef0ccfcfb30)


## APIs Exposed

- Login and get User's Access JWT token :  `POST /v1/login`
- Fetch all Users (Only By Admin) :        `GET /v1/restricted/users`
- Get User by Authenticated User:          `GET /v1/restricted/users/:userID`

- Create a new loan application by authenticated User:  `POST /v1/restricted/loan`
- Approve a Loan By Admin: `POST /v1/restricted/loan/approve/:loanID`
- Fetch and View a loan application by authenticated User:  `GET /v1/restricted/loan`
- Repay a loan by the authenticated User: `POST /v1/restricted/loan/repay/:loanID`

Postman collection: https://api.postman.com/collections/267026-0d13f189-fbbf-41af-8fcc-0e2640d452b6?access_key=PMAT-01H6KPG19D1QDDC0M23DMCAZQW