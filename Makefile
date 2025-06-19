.PHONY: benchmark

run:
	docker-compose up --build

test:
	go test -v ./handlers_test/

benchmark:
	go test -bench=. -benchmem -v ./benchmark  

migrate:
	./scripts/migrate.sh
