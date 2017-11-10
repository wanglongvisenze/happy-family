OS=$(shell uname)
UID=$(shell id -u)
GID=$(shell id -g)
ifeq ($(OS),Darwin)
	PROJECT_PWD=${PWD}
else
	PROJECT_PWD=$(shell readlink -f ${PWD})
endif
PROJECT_PKG=github.com/wanglongvisenze/happy-family
DOCKER_GOPATH=/go
DOCKER_RUN_OPT=--rm -v ${GOPATH}:${DOCKER_GOPATH} -w ${DOCKER_GOPATH}/src/${PROJECT_PKG}
DOCKER_RUN_OPT_USER=${DOCKER_RUN_OPT} -u ${UID}:${GID}
GO_SWAGGER=docker run ${DOCKER_RUN_OPT_USER} quay.io/goswagger/swagger
GO_IMAGE=docker run ${DOCKER_RUN_OPT_USER} visenze/golang:1.8
GO_IMAGE_ROOT=docker run ${DOCKER_RUN_OPT} visenze/golang:1.8
SERVER_PKG_ROOT=happy_family/v1
SERVER_EXECUTABLE=bin/happy-family-server
#SWAGGER_OUTPUT=$(SERVER_PKG_ROOT)/cmd/happy-family-server/main.go
SWAGGER_OUTPUT=$(SERVER_PKG_ROOT)/restapi/server.go

default: build

swagger: $(SWAGGER_OUTPUT)

swagger-clean:
	rm -rf $(SWAGGER_OUTPUT)

$(SWAGGER_OUTPUT): swagger.yml
	mkdir -p $(SERVER_PKG_ROOT)
	$(GO_SWAGGER) generate server -t $(SERVER_PKG_ROOT)

$(SERVER_EXECUTABLE): dep
	${GO_IMAGE} go build -o $(SERVER_EXECUTABLE) $(PROJECT_PKG)/$(SERVER_PKG_ROOT)/cmd/happy-family-server

build: $(SERVER_EXECUTABLE)

dep: swagger
	${GO_IMAGE} go get ./$(SERVER_PKG_ROOT)/...

test: build
	$(GO_IMAGE) go test $(PROJECT_PKG)/$(SERVER_PKG_ROOT)/restapi/
