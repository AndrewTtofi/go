# Postgres Exercise by Savvas (CRUD)

## First initialize DB to test the program
Run

```
cd db
docker build -t postgres-users .
```
Then
```
docker run 
    --name postgres-exercise 
    -p 5432:5432 
    -d postgres-users
```

## Build the application and run
```
cd ../app
go build
./app
```
## Test the application that runs successfully 
```
curl -X POST -d "username=andreastt&password=ttofis123" http://localhost:8080/login
```