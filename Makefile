# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=mybinary
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v ./...
clean: 
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
		rm -f $(BINARY_UNIX)
run:
		make clean
		$(GOBUILD) -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
deps:
		$(GOGET) github.com/markbates/goth
		$(GOGET) github.com/markbates/pop
dbrestart:
		docker stop Gmysql Gredis
		docker rm Gmysql Gredis && True
		docker run --name Gmysql -v $(PWD)/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql
		docker run --name Gredis -p 6379:6379 -v $(PWD)/data:/data  -d redis redis-server --appendonly yes
dbstart:
		docker run --name Gmysql -v $(PWD)/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql
		docker run --name Gredis -p 6379:6379 -v $(PWD)/data:/data  -d redis redis-server --appendonly yes
dockerrun:
		docker run --entrypoint=/bin/bash --restart=always --name=autotool -v /root/autotool/hook:/root/hook -v /root/autotool/post: /root/website/content/post -ti -p 8080:8080  autotool:v8.3