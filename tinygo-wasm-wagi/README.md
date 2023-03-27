
./compile_with_docker.sh

wasmtime main.wasm

docker run -ti --rm -p 8080:8080 aslom/tinygo-wasm-wagi:latest

./docker.sh

# TODO replace registry in used image eith sed// ?

kubectl apply -f func.yaml

# TODO use kn ?


