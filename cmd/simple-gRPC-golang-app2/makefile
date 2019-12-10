BUILD_FOLDER=./build
APPFOLDER=$(GOPATH)/src/Vakratunda/simple-gRPC-golang/cmd/$(APP)/internal/
APP=simple-gRPC-golang-app2

GOCMD= go
GOBUILD=$(GOCMD) build
GOCLEAN= $(GOCMD) clean
GOTEST= $(GOCMD) test
BINARY_NAME=app2
CLIENT_FOLDER=client/
PROTO_FILES= ./internal/proto
RPC_FOLDER=$(GOPATH)/src/Vakratunda/simple-gRPC-golang/cmd/rpc

.PHONY: test build
all: test build

clientb: 
	$(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v  $(APPFOLDER)$(CLIENT_FOLDER)

proto1:
	mkdir -p $(RPC_FOLDER)
	mkdir -p $(PROTO_FILES)
	protoc -I=proto --go_out=plugins=grpc:$(RPC_FOLDER) $(PROTO_FILES)/*


build: 
	cd ./internal/client/ && $(GOBUILD) -o ../../$(BUILD_FOLDER)/$(BINARY_NAME) -v 

build-linux:
	 GOOS=linux GOARCH=amd64  $(GOBUILD) -o $(BUILD_FOLDER)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./test	
clean:
	$(GOCMD) clean
run: build1
	echo $(PWD)
	$(BUILD_FOLDER)/$(BINARY_NAME)

run-client:	client1
	$(BUILD_FOLDER)/$(BINARY_NAME)

dep:
	rm -f Gopkg.lock
	rm -f Gopkg.toml
	dep init
	dep ensure

	


