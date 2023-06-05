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
docker build -t go-casbin . && docker run -p 8090:8888 -it go-casbin
```

# Request:
First you need log in!!!
* Method: ```POST```
URL :
```
http://localhost:8090/login
```
* Response: 
                ```
                 tokenString
                ```
Then set tokenString to Authorization Header with Bearer
1. Then do request to : ```http://localhost:8090/read```
2. Second endpoint: ```http://localhost:8090/book```