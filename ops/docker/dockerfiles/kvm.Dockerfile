FROM alpine:latest

ADD extras/docker_pid1 /usr/local/bin
ENTRYPOINT ["docker_pid1"]

RUN set -eux \
 && apk -U add iproute2 qemu-system-x86_64 qemu-img

RUN set -eux \
 && apk -U add sudo
