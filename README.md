## GO Pairs

### Install packages

```
go mod download        
go mod tidy
go mod vendor   

```

### Execution:
```
go run main.go
```

### API cUrl:
```
curl --location 'http://127.0.0.1:3341/api/find-pairs' \
--header 'Content-Type: application/json' \
--data '{
  "numbers": [1,2,3,4,5,6],
  "target": 6000
}'
```