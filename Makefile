# Variables
SERVICE_FILE := corpora.service
INSTALL_PATH := /etc/systemd/system/
# Install the service file to /etc/systemd/system/
copy:
	sudo cp $(SERVICE_FILE) $(INSTALL_PATH)
create:
	sudo mkdir -p bin
	sudo chmod 7777 bin
install: copy reload create build start enable status log
run: build
	@ ./bin/corpora
build: 
	@go build -o bin/corpora -buildvcs=false ./cmd/web
dev:
	go run cmd/web/main.go
reload: 
	systemctl daemon-reload
start:
	sudo systemctl start corpora
enable:
	sudo systemctl enable corpora
stop:
	sudo systemctl stop corpora
restart:
	sudo systemctl restart corpora
status:
	sudo systemctl status corpora
pull:
	sudo git pull
update: pull build restart status log

#Print out the log after running with systemd
log:
	journalctl -u corpora -f