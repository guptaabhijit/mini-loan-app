# mini-loan-app
Mini Loan App consisting of APIs written in golang


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

