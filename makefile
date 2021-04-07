run-serve:
	npx nodemon --exec go run main.go --signal SIGTERM

test:
	go test ./tests -v

test-coverage:
	go test ./tests -v -coverprofile=./tests/coverage.out -coverpkg=./... && go tool cover -html=./tests/coverage.out