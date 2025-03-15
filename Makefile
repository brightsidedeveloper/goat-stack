.PHONY: build

build:
	GOOS=js GOARCH=wasm go build -o client/public/main.wasm .
	touch client/src/main.js

dev:
	(find . -name "*.templ" | entr -r templ generate & \
	find . -name "*.go" | entr -r make build & \
	wait) || (kill 0; exit 1)