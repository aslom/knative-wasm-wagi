FROM aslom/knative-func-wasm-wagi:latest

COPY main.wasm /
#COPY modules.toml /
CMD [ "wagi", "-l", "0.0.0.0:8080", "-c", "/modules.toml" ]
