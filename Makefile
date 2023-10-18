WIN_USER=$(shell powershell.exe -c "[System.Environment]::UserName")
INSTALL_DIR=/mnt/c/Users/$(WIN_USER)/AppData/Local/NeovimWSL
START_MENU_DIR=/mnt/c/Users/$(WIN_USER)/AppData/Roaming/Microsoft/Windows/Start Menu/Programs/NeovimWSL
BIN=NeovimWSL.exe

.PHONY: default
default: build

.PHONY: help
help:
	@echo "Targets:"
	@echo "  help             - print this help"
	@echo "  build (default)  - build executable"
	@echo "  clean            - removes generated files"
	@echo "  icon             - updates the application icon"
	@echo "  install          - installs the application to Windows"
	@echo "  icon             - uninstalls the application from Windows"

.PHONY: build
build:
	GOOS=windows GOARCH=amd64 go build -o $(BIN)

.PHONY: clean
clean:
	rm -f $(BIN)

.PHONY: icon
icon:
	# go install https://github.com/tc-hib/go-winres
	go-winres simply --icon icon.png

.PHONY: install
install:
	mkdir -p "$(INSTALL_DIR)"
	mkdir -p "$(START_MENU_DIR)"
	cp $(BIN) "$(INSTALL_DIR)"
	cp NeovimWSL.lnk "$(START_MENU_DIR)"


.PHONY: uninstall
uninstall:
	rm -rf "$(INSTALL_DIR)"
	rm -rf "$(START_MENU_DIR)"
