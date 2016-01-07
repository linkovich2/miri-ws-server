ifeq ($(OS),Windows_NT)
	APPNAME := 'miri-ws-server.exe'
else
  APPNAME := 'miri-ws-server'
endif

all:
	cd app/; go-bindata data/; cd ..; godep go build; ./$(APPNAME)
