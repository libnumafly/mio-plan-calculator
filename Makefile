BINARY  := iij-kumiawase
WASM    := main.wasm

.PHONY: all build wasm wasm_exec clean

all: build wasm

build:
	go build -o $(BINARY) .

wasm:
	GOOS=js GOARCH=wasm go build -o $(WASM) .

wasm_exec:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .

clean:
	rm -f $(BINARY) $(WASM)
