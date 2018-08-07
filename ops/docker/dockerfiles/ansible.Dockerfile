FROM alpine:latest

ADD extras/docker_pid1 /usr/local/bin
ENTRYPOINT ["docker_pid1"]

RUN set -eux \
 && apk -U add gcc make python2-dev musl-dev libffi-dev openssl-dev \
               python2 py2-pip \
               openssh-client rsync \
 && pip install --upgrade pip \
 && pip install ansible~=2.6
