default: help

# Help Menu for makefile
.PHONY: help
help:		## List all available commands
	@echo '--- Available commands ----'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST) 
	@echo
	@echo '--- Build/Run program ---'
	@echo 'Please use wails command to build/run program.'
	@echo ' - Use "wails build" to build executable'
	@echo ' - Use "wails dev" to Runs the application in development mode'

.PHONY: audit-fix
audit-fix:	## Run npm audit fix for frontend
	cd frontend; npm audit fix

.PHONY: binding
binding:	## Generate wails binding
	wails generate module

.PHONY: clean 
clean: 		## Remove distribution of both frontend and backend
	cd frontend; rm -rf dist
	rm build/bin/*.exe

.PHONY: demo
demo:		## Create a demo comic directory in 'testing/', contain total 5 images with 1MB
	rm -rf 'testing/[author] title'
	mkdir -p 'testing/[author] title'
	@dd if=/dev/zero of='testing/[author] title/image1.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image2.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image3.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image4.jpg' bs=1024 count=0 seek=1024
	@dd if=/dev/zero of='testing/[author] title/image5.jpg' bs=1024 count=0 seek=1024

.PHONY: reinstall
reinstall: 	## delete and then install all node_modules of frontend
	cd frontend; rm -rf node_modules && npm install

.PHONY: test
test:		## Perform all tests of golang
	go test ./... -count=1
