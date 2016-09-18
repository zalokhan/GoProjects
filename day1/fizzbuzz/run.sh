go test -coverprofile cov.out
go tool cover -func cov.out
go tool cover -html cov.out
go test --bench Bench --benchmem
go test -v
