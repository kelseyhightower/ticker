```
protoc -I=. --go_out=. quote.proto
```

```
protoc --go_out=plugins=grpc:. quote.proto
```
