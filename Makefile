.DEFAULT_GOAL := help
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: help
help: ## Displays this help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-z0-9A-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: aoc2023
aoc2023: ## Compiles the code for 2023
	@CGO_ENABLED=0 go build -o aoc2023 ./2023/main.go

.PHONY: build
build: ## Compiles the code
	@CGO_ENABLED=0 go build -o adventofcode .

.PHONY: run
run: build ## Runs a specific day or all days (no args)
	@./adventofcode $(ARGS)

## Tiny hack to allow passing of args: https://stackoverflow.com/a/47008498
%:
	@:
