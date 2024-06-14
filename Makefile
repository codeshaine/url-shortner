
.PHONY :build run test test-v test-force clean 
build:
	go build -o bin/urlShorner ./cmd/main.go

run:
	go run ./cmd/main.go 
 
test:
	go test ./test 

test-v:
	go test ./test -v
	 
test-force :
	go test -count=1 ./test -v

clean:
	rm -rf ./bin

# for removing the unwanted packages and update the packages
sync:
	go mod tidy