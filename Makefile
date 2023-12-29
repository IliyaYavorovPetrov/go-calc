build:
	go build -o bin/main cmd/main.go

test:
	go test -v ./app/...

clean:
	go clean
	rm -f ./bin/main

run:
	./bin/main
