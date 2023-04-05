## Run Hello world written in Go with Wasm

Use TinyGo compiler that generates main.was

```
./compile_with_docker.sh
```

Test 

```
wasmtime main.wasm
```

Build local docker image

```
./docker_build.sh
``

Test docker image

```
docker run -ti --rm -p 8080:8080 aslom/tinygo-wasm-wagi:latest
```

Publish image to registry

```
./docker_push.sh
```


kubectl apply -f func.yaml

# TODO use kn ?


# TODO replace registry in used image eith sed// ?
