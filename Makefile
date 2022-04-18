clean:
	rm -rf build/bin

build: clean
	npm --prefix ./frontend run build && wails build -clean

launch: build
	open ./build/bin/Varly.app/Contents/MacOS/Varly

test:
	go test ./... -v

sign: release
	gon -log-level=info ./build/darwin/gon-sign.json

notarize:
	xcrun altool --notarize-app --primary-bundle-id "com.varlylabs.Varly" -u "selvin@selvin.co" -p "@env:APPLE_ID_PASSWORD" --asc-provider NA229UVJJB --file ./build/darwin/Varly.dmg --output-format xml

release:
	npm --prefix ./frontend run build && wails build -clean -nsis -webview2 download -platform darwin/amd64,darwin/arm64,windows/amd64,windows/arm64
