clean:
	rm -rf build/bin

build: clean
	wails build -clean

launch: build
	open ./build/bin/varly.app/Contents/MacOS/varly

test:
	go test ./... -v
