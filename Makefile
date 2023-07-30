PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## dependencies: mockgen required for generating mock services under rabbitmq 
dependencies:
	echo "Setting mockgen dependency"
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0

## all: Generates code, install dependencies, compiles the binary, start the client - all in one shot.
all: generate install compile start

## install: Install missing dependencies. Runs `go get` internally. e.g; make install get=github.com/foo/bar
install: dependencies go-get precommit

## start: Start the client.
start: stop
	@echo "  >  $(PROJECTNAME) is available at $(ADDR)"
	@-$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

## stop: Stop the client.
stop:
	@echo "  >  Killing the running instance of the server..."
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)
	@echo "  >  Done."

## compile: Compile the binary.
compile: go-compile

## generate: Generates the client code.
generate: dependencies
	@echo "  >  Starting rabbitmq service mock code generation....."
	mockgen -source=client/service.go -destination=client/mock_service.go -package=client AmqpClientType
	mockgen -source=connection/service.go -destination=connection/mock_service.go -package=connection AmqpConnectionType
	mockgen -source=producer/service.go -destination=producer/mock_service.go -package=producer ProducerType
	mockgen -source=notification/service.go -destination=notification/mock_service.go -package=notification FileNotification
	@echo "  >  Finished generating rabbitmq service mock code generation."
	@echo "  >  Finished generating client code."

## install-amqp: locally deploy the rabbitmq to run the unit tests
install-amqp:
	@echo "  >  locally deploying rabbitmq as docker cointainer....."
	@sudo docker run -d -it --rm --name rabbitmq -p 5673:5672 -p 15673:15672 rabbitmq:3-management &
	@sudo sleep 30

## remove-amqp: remove locally deployed rabbitmq
remove-amqp:
	@echo "  >  removing locally deployed rabbitmq as docker cointainer....."
	@sudo docker ps -q --filter "name=rabbitmq" | grep -q . && sudo docker stop rabbitmq || true
	@sudo sleep 10

## unittest: Runs the Golang unit test files tagged with unit
unit-test: go-test-setup generate go-test go-test-teardown
	
## test: Runs tests.
test: go-test-setup go-test go-test-teardown

## go-test-setup: test: install amqp
go-test-setup: install-amqp

## go-test-teardown: remove amqp
go-test-teardown: remove-amqp

## clean: Clean build files. Runs `go clean` internally.
clean: go-clean

precommit:
	@echo "  >  Installing pre-commit hooks..."
	@git init
	@pre-commit install
	@pre-commit run --all-files

go-test:
	@echo "  >  Running tests..."
	go test ./... -coverprofile=coverage.out

go-compile: go-get go-build

go-build:
	@echo "  >  Building binary..."
	@CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(PROJECTNAME)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	go get $(get)

go-clean:
	@echo "  >  Cleaning build cache"
	go clean

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
