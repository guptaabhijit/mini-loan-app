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

![Screenshot 2023-07-30 at 9 05 20 PM](https://github.com/guptaabhijit/mini-loan-app/assets/7620035/b0cc49ed-fa9c-408c-b12f-eef0ccfcfb30)
