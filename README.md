To run the tests

```
go test -run ^YourTestName
```

or just
```
go test -shuffle=on -v
```

to run the benchmark tests only

```
go test -bench . -run notest
```

To add module requirements
```
go mod tidy
```

To format the code
```
gofmt -w *.go
```


Download dependancies
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```