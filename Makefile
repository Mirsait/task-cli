APP_NAME = task-cli
BUILD_DIR = build

all: build

build:
	go build -o $(BUILD_DIR)/$(APP_NAME)

install: build
	install -Dm755 $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)

uninstall:
	rm -f /usr/local/bin/$(APP_NAME)
