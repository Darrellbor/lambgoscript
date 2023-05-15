all: test build lint

# golang executions
run-sort:
	cd golang && \
	go run cmd/sort/run.go

run-fibonacci:
	cd golang && \
	go run cmd/fibonacci/run.go

build-sort:
	cd golang && \
	env GOOS=linux go build -o bin/sort/run cmd/sort/run.go && \
	cd bin/sort && zip upload-sort.zip run && cd .. && cd .. && \
	go test . -run xxx

build-fibonacci:
	cd golang && \
	env GOOS=linux go build -o bin/fibonacci/run cmd/fibonacci/run.go && \
	cd bin/fibonacci && zip upload-fibonacci.zip run && cd .. && cd .. && \
	go test . -run xxx

test:
	cd golang && \
	go test .
	go test -race .

lint:
	cd golang && \
	golangci-lint run .

tidy:
	cd golang && \
	go mod tidy

fmt:
	cd golang && \
	gofmt -s -w .

#typescript executions
packageInstall: 
	cd typescript && \
	npm install

buildTypescript: 
	cd typescript && \
	npm run build

startDev: 
	cd typescript && \
	npm run start.dev

start:  buildTypescript
	cd typescript && \
	npm start

testTypescript: 
	cd typescript && \
	npm run test

watchTestTypescript: 
	cd typescript && \
	npm run watch-test
	