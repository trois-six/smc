.PHONY: check clean build swagger security generate

default: build

check:
	@golangci-lint run

clean:
	@rm -f smc_flattener

build:
	@go build -o smc_flattener

swagger: build
	@./smc_flattener
	@echo "  securitySchemes:\n    apiKey:\n      type: http\n      scheme: bearer\n\nsecurity:\n  - apiKey: []" >> swagger_flat.yaml

generate:
	@rm -f ../smc.gen.go && cd tools && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --config=../client.yaml --include-tags="default,$(cat ../swagger_flat.yaml | yq e '.. | select(has("tags")) | .tags[]' | sort -u | grep -Ev 'firewalls|objects|proxy|rules|rule-sets' | paste -sd ',')" ../swagger_flat.yaml
