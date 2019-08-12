GO = go
PACKAGE = mower
BINDIR = $(CURDIR)/bin
OUT = $(BINDIR)/$(PACKAGE)

GOPATH = $(CURDIR)/.gopath
GOBIN = $(GOPATH)/bin
GOSRC = $(GOPATH)/src
BASE = $(GOSRC)/$(PACKAGE)

GODEP = $(GOBIN)/dep

export GOPATH
export PATH := $(PATH):$(GOBIN)

.PHONY: all
all: vendor build 

##
# Project
##

$(BASE):
	mkdir -p $(dir $@)
	ln -sf $(CURDIR)/src $@

.PHONY: build
build: | $(BASE) ## Build project
	cd $(BASE) && $(GO) build -o $(OUT) main.go commands.go

.PHONY: fmt
fmt: ## Format all Go files in src/
	$(GO) fmt ./src/...

##
# Dependencies management
##

$(BASE)/Gopkg.toml: | $(BASE) $(GODEP)
	if [ ! -f $(BASE)/Gopkg.toml ]; then cd $(BASE) && $(GODEP) init; fi

.PHONY: vendor
VFLAGS=
vendor: | $(BASE)/Gopkg.toml ## Pull/update project dependencies (imports) -- can be used with `VFLAGS=-update`
	cd $(BASE) && $(GODEP) ensure -v $(VFLAGS)

##
# Tools
##

$(GOBIN):
	mkdir -p $@

$(GOBIN)/%: | $(BASE) $(GOBIN)
	$(GO) get $(REPOSITORY)

$(GODEP): REPOSITORY=github.com/golang/dep/cmd/dep

##
# Misc
##

.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean:
	$(RM) $(OUT)

.PHONY: distclean
distclean: clean
	$(RM) -r src/vendor $(BINDIR) $(GOPATH)
