all: test build lint

# golang executions
run:
	cd golang && \
	go run cmd/run.go

build:
	cd golang && \
	env GOOS=linux go build -o bin/run cmd/run.go && \
	cd bin && zip upload.zip run && cd .. && \
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
	