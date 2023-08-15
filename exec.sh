$ export PATH="$PATH:$(go env GOROOT)/misc/wasm"
$ GOOS=js GOARCH=wasm go run .
Hello Wasm!
