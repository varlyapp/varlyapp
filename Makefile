clean:
	rm -rf build/bin

dev:
	wails dev & npm --prefix ./frontend run dev

build: clean
	npm --prefix ./frontend run release && wails build -clean

launch: build
	open ./build/bin/varly.app/Contents/MacOS/Varly

test:
	go test ./... -v

sign: build
	gon -log-level=info ./build/darwin/gon-sign.json

notarize:
	xcrun altool --notarize-app --primary-bundle-id "app.varly.Varly" -u "selvin@selvin.co" -p "@env:APPLE_ID_PASSWORD" --asc-provider NA229UVJJB --file ./build/Varly.dmg --output-format xml
