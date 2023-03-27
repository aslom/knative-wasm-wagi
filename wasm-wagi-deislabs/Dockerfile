# Based on https://github.com/denverdino/knative-wagi/blob/master/Dockerfile
FROM rust:1.61 AS builder
RUN git clone https://github.com/deislabs/wagi
RUN cd wagi && cargo build --release

FROM debian:bullseye-slim
#COPY sources.list /etc/apt/sources.list
COPY --from=builder /wagi/target/release/wagi /usr/local/bin/wagi
COPY modules.toml /
#COPY --from=builder /wagi/examples /examples
CMD [ "wagi", "-l", "0.0.0.0:8080", "-c", "/modules.toml" ]
