ifeq ($(OS),Windows_NT)
	APPNAME := 'miri-ws-server.exe'
else
  APPNAME := 'miri-ws-server'
endif

all:
	cd app/content/; go-bindata -pkg 'content' json/...; cd ../../; go test ./app/test/...; godep go build; ./$(APPNAME)
