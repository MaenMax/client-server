GOPATH:=$(shell pwd)
GO:=go

default: bin/server bin/publisher

all: clean default


bin/server:   src/server.go
	@echo "========== Compiling $@ =========="
	sh -c '$(GO) build   ./src/server.go &&  mv server ./bin'

bin/publisher:  utils/publisher.go
	@echo "========== Compiling $@ =========="
	sh -c '$(GO) build  ./utils/publisher.go && mv publisher ./bin'


clean:
	@echo "Deleting generated binary files ..."; sh -c 'if [ -d bin ];  then  find bin/ -type f -exec rm {} \; -print ; fi; rm -Rf bin/*;rm -rf ./*.tar.bz2'


clear_cache:
	@echo "Deleting local caches ..."; 
	@if [ -d ~/go ]; then find ~/go -type d -exec chmod 755 {} \; && rm -Rf ~/go; fi

deploy:
	@echo "Deploying server..."; 
	bash -c 'mkdir -p client-server && mkdir -p client-server/bin && cp -a bin client-server && cp -a conf client-server && mkdir -p client-server/log && tar -jcvf client-server.tar.bz2 client-server'
