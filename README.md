# Clone the repository
* Run this command if you have ssh key:
```
git clone git@github.com:Allayar07/go-casbin.git
```
* For http:
```
git clone https://github.com/Allayar07/go-casbin.git
```
# Run  Project
* Use this command:
```
docker compose up
```
# DB Migrations 
* Command:
```
migrate -path ./schema -database 'postgres://postgres:password0701@localhost:5432/practice?sslmode=disable' up
```
# Request:
* First you need create 2 user!!!

* Method ```POST```

URL :
```
http://localhost:9999/create_user
```
* Body for superuser:
```
{
    "name": "User1",
    "password": "pass123",
    "phone": 223433,
    "address":"Ashgabat",
    "role": "superuser"
}
```
* Response:

```
 {
	"id": 0,
	"name": "User1",
	"phone": 223433,
	"address": "Ashgabat",
	"role": "superuser"
}
```

* Body for customer:
```
{
    "name": "User2",
    "password": "pass12345",
    "phone": 223433,
    "address":"Ashgabat",
    "role": "customer"
}
```
* Response:

```
 {
	"id": 0,
	"name": "User2",
	"phone": 223433,
	"address": "Ashgabat",
	"role": "customer"
}
```

* Then you need to log in!!!

* Method: ```POST```
URL :
```
http://localhost:9999/login
```
* Body:
```
{
    "name": "User",
    "password": "pass123"
}
```
* Response:

```
 {
    "token": "token_string"
 }
```

Then set tokenString to Authorization Header
1. Then do request to (```METHOD:``` GET) : ```http://localhost:9999/read```
2. Second endpoint (```METHOD:``` POST): ```http://localhost:9999/book```