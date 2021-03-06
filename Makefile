VERSION   := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GO_CMD      = go
GO_GET      = $(GO_CMD) get
GO_BUILD    = $(GO_CMD) build
GO_CLEAN    = $(GO_CMD) clean
GO_LDFLAGS  = -s -w
GO_LDFLAGS += -X main.buildVersion=$(VERSION)
GO_LDFLAGS += -X main.buildTime=$(BUILDTIME)
GO_FLAGS    = -ldflags "$(GO_LDFLAGS)"

.PHONY: all
all: blocket

blocket: *.go
	CGO_ENABLED=1 $(GO_BUILD) -o blocket $(GO_FLAGS) *.go

.PHONY: clean
clean:
	$(GO_CLEAN)

.PHONY: deps
deps:
	$(GO_GET) .
