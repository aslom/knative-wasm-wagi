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

Directly deploy using YAML:

```
kubectl apply -f hello-wasm-yaml.yaml
```

Use kn CLI:

```
kn service create hello-wasm-kn --image docker.io/aslom/tinygo-wasm-wagi:latest --port 8080
```

# TODO replace registry in used image eith sed// ?
