set -v
docker run --rm -v $(pwd):/src tinygo/tinygo:0.27.0 tinygo build  -target=wasi -o /src/main.wasm /src/main.go
