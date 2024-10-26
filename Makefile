default: help

# Help Menu for makefile
.PHONY: help
help:		## List all available commands
	@echo '--- Available commands ----'
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST) 
	@echo
	@echo '--- Run program ---'
	@echo 'Please use wails command to run program.'
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
	rm -f build/bin/*.exe

.PHONY: dist
dist:		## Create a wails executeable distribution
	rm -f build/bin/*.*
	wails build
	@CURRENT_TAG=$$(git tag --points-at);\
	CURRENT_TAG=$$(echo "$${CURRENT_TAG////-}");\
	if [ ! -z "$${CURRENT_TAG}" ]; then\
		cd build/bin; mv ComicInfo-Parser.exe "ComicInfo-Parser-$${CURRENT_TAG}.exe";\
		echo "Distribution ready at 'build/bin/ComicInfo-Parser-$${CURRENT_TAG}.exe'.";\
	else\
		cd build/bin; mv ComicInfo-Parser.exe "ComicInfo-Parser-testing.exe";\
		echo "Distribution ready at 'build/bin/ComicInfo-Parser-testing.exe'.";\
	fi

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

.PHONY: tag-dev
tag-dev: 	## Create local git tag for current time, with timezone in UTC+0
	@TIME_STR=$$(TZ=UTC date +"%Y%m%d-%H%M");\
	git tag -a "dev/$${TIME_STR}" -m "development build" -f;\
	echo "Development tag completed: 'dev/$${TIME_STR}'."

.PHONY: test
test:		## Perform all tests of golang
	go test ./... -count=1
