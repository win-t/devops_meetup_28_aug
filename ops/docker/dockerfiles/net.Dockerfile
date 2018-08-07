FROM scratch

ADD extras/docker_pid1 /
ENTRYPOINT ["/docker_pid1"]
