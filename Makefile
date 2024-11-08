# Variables
SERVICE_FILE := corpora.service
INSTALL_PATH := /etc/systemd/system/
# Install the service file to /etc/systemd/system/
install:
	sudo cp $(SERVICE_FILE) $(INSTALL_PATH)
run: build
	@ ./bin/corpora
build: 
	@go build -o bin/corpora -buildvcs=false
dev:
	go run main.go
start:
	sudo systemctl start corpora
stop:
	sudo systemctl stop corpora
restart:
	sudo systemctl restart corpora
status:
	sudo systemctl status corpora
#Print out the log after running with systemd
log:
	journalctl -u corpora -f