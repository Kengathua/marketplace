TEST_DATABASE_URL=postgres://matawis_user:matawis_pass@localhost:5432/matawis_test

.PHONY: tidy migratetestup migratetestdown test

tidy:
	go mod tidy

migratetestup:
	migrate -database ${TEST_DATABASE_URL} -path db/migrations up

migratetestdown:
	migrate -database ${TEST_DATABASE_URL} -path db/migrations down

test:
	go test -v -timeout 100s -coverprofile=coverage.out -cover ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html
