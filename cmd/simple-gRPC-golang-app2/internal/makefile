BUILD_FOLDER=../build
APPFOLDER=$(GOPATH)/src/Vakratunda/simple-gRPC-golang/cmd/$(APP)/internal/
APP=simple-gRPC-golang-app2

GOCMD= go
GOBUILD=$(GOCMD) build
GOCLEAN= $(GOCMD) clean
GOTEST= $(GOCMD) test
BINARY_NAME=app2
CLIENT_FOLDER=client/

all: test build

client1: 
	$(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v  $(APPFOLDER)$(CLIENT_FOLDER)

build: 
	$(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v 

build-linux:
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64  $(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./test	
clean:
	$(GOCMD) clean
run: 	
	build
	$(BUILD_FOLDER)/$(BINARY_NAME)

run-client:	client1
	$(BUILD_FOLDER)/$(BINARY_NAME)

dep:
	rm -f Gopkg.lock
	rm -f Gopkg.toml
	dep init
	dep ensure

	


