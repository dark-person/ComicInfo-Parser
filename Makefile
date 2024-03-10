default: help

# Help Menu for makefile
.PHONY: help
help:	## List all available commands
	@echo 'Available commands:'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST) 

.PHONY: clean 
clean: 	## Remove distribution of both frontend and backend
	rm build/bin/*.exe

.PHONY: test
test:	## Perform all tests of golang
	go test ./... -count=1

.PHONY: cgo-on
cgo-on:	## Enable CGO_ENABLED 
	set CGO_ENABLED=1