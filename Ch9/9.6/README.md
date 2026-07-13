# Results

## Initial Run

```
go run main.go
> time to run 6.598083ms
```

## GOMAXPROCS

```
GOMAXPROCS=0 go run ./main.go
>time to run 5.908125ms
GOMAXPROCS=1 go run ./main.go
>time to run 62.987083ms
GOMAXPROCS=2 go run ./main.go
>time to run 35.21675ms
GOMAXPROCS=3 go run ./main.go
>time to run 25.284084ms
GOMAXPROCS=4 go run ./main.go
>time to run 19.711ms
GOMAXPROCS=5 go run ./main.go
>time to run 15.8165ms
```