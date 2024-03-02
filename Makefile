default: help

# Help Menu for makefile
.PHONY: help
help:
	@echo 'Available commands: clean, install, reinstall, test'

# clean all build file
.PHONY: clean
clean:
	rm build/bin/*.exe

## Install Dependencies
.PHONY: install
install:
	cd frontend; npm install

## Reinstall all dependencies
.PHONY: reinstall
reinstall:
	cd frontend; rm -rf node_modules && npm install

# Run all tests for golang
.PHONY: test
test:
	go test ./... -count=1
