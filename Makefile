COMMIT_ID=$(shell git rev-parse --short HEAD)
VERSION=$(shell cat VERSION)

NAME=tg-parser

all: clean build

clean:
	@echo ">> cleaning..."
	@rm -f $(NAME)

build: clean
	@echo ">> building..."
	@echo "Commit: $(COMMIT_ID)"
	@echo "Version: $(VERSION)"
	@go build -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID)" -o $(NAME) ./cmd/...
	@chmod +x ./$(NAME)

install:
	@go install -ldflags "-X main.Version=$(VERSION) -X main.CommitID=$(COMMIT_ID) client.Version=$(VERSION)" ./cmd/...

.PHONY: all clean build install
