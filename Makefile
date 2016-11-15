SHELL = bash

# Environment

GIT_BRANCH = $(or $(CI_BUILD_REF_NAME) ,`git rev-parse --abbrev-ref HEAD 2>/dev/null`)
GIT_COMMIT = $(or $(CI_BUILD_REF), `git rev-parse HEAD 2>/dev/null`)
BUILD_DATE = $(or $(CI_BUILD_DATE), `date -u +%Y-%m-%dT%H:%M:%SZ`)

# All

.PHONY: all build-deps deps dev-deps fmt vet ttn ttnctl build link docs clean install dev

all: deps build

# Deps
dev-deps: deps
	@command -v mockgen > /dev/null || go get github.com/golang/mock/mockgen
	@command -v golint > /dev/null || go get github.com/golang/lint/golint
	@command -v forego > /dev/null || go get github.com/ddollar/forego
	@command -v MessageConverter > /dev/null || go get github.com/LoRaWanSoFa/LoRaWanSoFa

fmt:
	[[ -z "`echo "$(GO_PACKAGES)" | xargs go fmt | tee -a /dev/stderr`" ]]

vet:
	echo $(GO_PACKAGES) | xargs go vet

lint:
	for pkg in `echo $(GO_PACKAGES)`; do golint $$pkg | grep -vE 'mock|.pb.go'; done

#Go Test

test: $(GO_FILES)
	go test $(GO_TEST_PACKAGES)

# Go Build

RELEASE_DIR ?= release
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOEXE = $(shell GOOS=$(GOOS) GOARCH=$(GOARCH) go env GOEXE)

splitfilename = $(subst ., ,$(subst -, ,$(subst $(RELEASE_DIR)/,,$1)))
GOOSfromfilename = $(word 2, $(call splitfilename, $1))
GOARCHfromfilename = $(word 3, $(call splitfilename, $1))
LDFLAGS = -ldflags "-w -X main.gitBranch=${GIT_BRANCH} -X main.gitCommit=${GIT_COMMIT} -X main.buildDate=${BUILD_DATE}"
GOBUILD = CGO_ENABLED=0 GOOS=$(call GOOSfromfilename, $@) GOARCH=$(call GOARCHfromfilename, $@) go build -a -installsuffix cgo ${LDFLAGS} -o "$@"


MessageConverter: $(RELEASE_DIR)/MessageConverter-$(GOOS)-$(GOARCH)$(GOEXE)

$(RELEASE_DIR)/MessageConverter-%: $(GO_FILES)
	$(GOBUILD) ./main.go

MQTTClient: $(RELEASE_DIR)/MQTTClient-$(GOOS)-$(GOARCH)$(GOEXE)

$(RELEASE_DIR)/MQTTClient-%: $(GO_FILES)
	$(GOBUILD) ./MQTTClient/main.go

build: MessageConverter MQTTClient

GOBIN ?= $(GOPATH)/bin

link: build
	ln -sf $(PWD)/$(RELEASE_DIR)/MQTTClient-$(GOOS)-$(GOARCH)$(GOEXE) $(GOBIN)/MQTTClient
	ln -sf $(PWD)/$(RELEASE_DIR)/MessageConverter-$(GOOS)-$(GOARCH)$(GOEXE) $(GOBIN)/MessageConverter

# Documentation

install:
	go install -v . ./MQTTClient

dev: install

docs:
	cd cmd/docs && HOME='$$HOME' go run generate.go > README.md
# Clean

clean:
	[ -d $(RELEASE_DIR) ] && rm -rf $(RELEASE_DIR) || [ ! -d $(RELEASE_DIR) ]
