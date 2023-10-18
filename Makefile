.PHONY: default
default: build

.PHONY: help
help:
	@echo "Targets:"
	@echo "  help             - print this help"
	@echo "  build (default)  - build executable"
	@echo "  clean            - removes generated files"
	@echo "  icon             - updates the application icon"

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build

.PHONY: clean
clean:
	go clean

.PHONY: icon
icon:
	# go install https://github.com/tc-hib/go-winres
	go-winres simply --icon icon.png


