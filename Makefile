run:
	@go run cmd/main.go
air:
	air server --port 3000
postgres-run:
	@docker run --name diploma_backend -p 5432:5432 -e POSTGRES_PASSWORD=some_pass -d postgres
postgres-start:
	@docker start static-link
postgres-stop:
	@docker stop static-link