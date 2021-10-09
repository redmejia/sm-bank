# small-bank-api 
The Small Bank is a REST-API project written in Go. On this project, I wanted to put in practice Go language concepts such as interface, methods, database connection, routing, and structs. The project Small Bank has four basic endpoints.  

# New client account

## account
The account endpoint for creating new client accounts. These accounts include the creation of saving and checking accounts, also card numbers and CVV numbers are created for new users.
```
	http://127.0.0.1:8081/v1/account
```
# Transactions endpoints

## deposit
The deposit for making a  transaction of type deposit. The client has two options to deposit on the savings or checking accounts.
```
	http://127.0.0.1:8081/v1/deposit/transaction
```

## withdraw
the withdrawal for making a new transaction of type withdraw. Same as the deposit client can withdraw from saving and checking accounts.
```
	http://127.0.0.1:8081/v1/withdraw/transaction
```

## purchase
The purchase endpoint. On this endpoint, an amount purchased is withdrawn from the checking account.
```
	http://127.0.0.1:8081/v1/purchase/transaction
```

For all these transactions card information is required.  The card information is verified with the information that is stored on the database. If card information does not match then the transaction is not complete and an error message is sent to the front-end in JSON format. Moreover, the Small Bank purchase endpoint is implemented on the [walk-api](https://github.com/redmejia/walk-api) every time a new purchase is made. Finally, all the client information is stored on the PostgreSQL database.


 
 
 

