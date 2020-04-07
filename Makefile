BINARY=craftchallenge
DOCKER_IMAGE=craftchallenge
test: 
	go test -v -cover -covermode=atomic ./...

build:
	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
	docker build -t ${DOCKER_IMAGE} .

run:
	docker run ${DOCKER_IMAGE} 

stop:
	docker stop ${DOCKER_IMAGE}

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run \
		--exclude-use-default=false \
		--enable=golint \
		--enable=gocyclo \
		--enable=goconst \
		--enable=unconvert \
		./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint