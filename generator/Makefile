.PHONY: all check clean build swagger security generate
YQ := $(shell command -v yq 2> /dev/null)
GO := $(shell command -v go 2> /dev/null)

default: build

all:
ifndef GO
    $(error "go is not available please install it")
endif

check:
	@golangci-lint run

clean:
	@rm -f smc_flattener

build:
	@go build -o smc_flattener

swagger: build
	@./smc_flattener

generate:
ifndef YQ
    $(error "yq is not available please install it")
endif
	@rm -f ../smc.gen.go && cd tools && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --config=../client.yaml --include-tags="default,$(cat ../swagger_flat.yaml | yq e '.. | select(has("tags")) | .tags[]' | sort -u | grep -Ev 'firewalls|objects|proxy|rules|rule-sets' | paste -sd ',')" ../swagger_flat.yaml
