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

.PHONY: demo
demo:	## Create a demo comic directory in 'testing/', contain total 5 images with 1MB
	rm -rf 'testing/[author] title'
	mkdir -p 'testing/[author] title'
	@dd if=/dev/zero of='testing/[author] title/image1.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image2.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image3.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image4.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image5.jpg' bs=1024 count=0 seek=1024