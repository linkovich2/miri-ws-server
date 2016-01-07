ifeq ($(OS),Windows_NT)
	APPNAME := miri-ws-server.exe
else
  APPNAME := miri-ws-server
endif

all: content test build
content:
	cd app/content/; go-bindata -pkg 'content' json/...; cd ../../;
test:
	go test ./app/test/...;
build:
	godep go build; ./$(APPNAME)
