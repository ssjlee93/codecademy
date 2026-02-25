.PHONY: build test clean run postgres

run:
	go run ./cmd/main.go

crawl:
	go run ./cmd/crawler/main.go

postgres:
	docker run --name postgres -e POSTGRES_DB=recommendation_software -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

#build:
#	go build -o myapp ./cmd/myapp
#
#test:
#	go test ./...
#
#clean:
#	rm -f myapp
#
#run: build
#	./myapp
