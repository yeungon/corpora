# Variables
SERVICE_FILE := corpora.service
INSTALL_PATH := /etc/systemd/system/
# Install the service file to /etc/systemd/system/
clone:
	sudo git clone "https://github.com/yeungon/corpora.git"
copy:
	sudo cp $(SERVICE_FILE) $(INSTALL_PATH)

install: clone build copy build start enable status log

run: build
	@ ./bin/corpora
build: 
	@go build -o bin/corpora -buildvcs=false ./cmd/web
dev:
	go run cmd/web/main.go
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
	sudo systemctl status corpora
#Print out the log after running with systemd
log:
	journalctl -u corpora -f