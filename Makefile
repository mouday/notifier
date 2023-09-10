# 开发环境
.PHONY: dev
dev:
	air

# 编译到 Linux
.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/notifier-linux ./src/main.go 

# 编译到 macOS
.PHONY: build-darwin
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/notifier-darwin ./src/main.go

# 编译到 windows
.PHONY: build-windows
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/notifier-windows.exe ./src/main.go 

# 编译到 全部平台
.PHONY: build
build:
	make clean
	mkdir -p ./build
	make build-linux
	make build-darwin
	make build-windows

.PHONY: clean
clean:
	rm -rf ./build
